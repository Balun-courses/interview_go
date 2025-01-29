package main

import (
	"context"
	"errors"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

const retriesNumber = 3
const retriesInitialPauseDuration = time.Millisecond * 100
const requestTimeout = time.Second * 5
const syncInterval = time.Minute

type RedisDatabase interface {
	Get(context.Context, string) (string, error)
	MGet(context.Context, []string) ([]*string, error)
	Keys(context.Context) ([]string, error)
}

type RedisDatabaseWithCache struct {
	database RedisDatabase
	group    singleflight.Group

	mutex sync.RWMutex
	cache map[string]string
}

func NewRedisDatabaseWithCache(ctx context.Context, database RedisDatabase) (*RedisDatabaseWithCache, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	if database == nil {
		return nil, errors.New("incorrect database")
	}

	c := &RedisDatabaseWithCache{
		database: database,
		cache:    make(map[string]string),
	}

	go c.synchronize(ctx)
	return c, nil
}

func (c *RedisDatabaseWithCache) synchronize(ctx context.Context) {
	ticker := time.NewTicker(syncInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			c.synchronizeImpl(ctx)
		}
	}
}

func (c *RedisDatabaseWithCache) synchronizeImpl(ctx context.Context) {
	var keys []string
	err := withRetries(ctx, func(ctx context.Context) error {
		var err error
		keys, err = c.database.Keys(ctx)
		return err
	})

	if err != nil {
		return // logs, metrics
	}

	var values []*string
	err = withRetries(ctx, func(ctx context.Context) error {
		var err error
		values, err = c.database.MGet(ctx, keys)
		return err
	})

	if err != nil {
		return // logs, metrics
	}

	cache := make(map[string]string, len(keys))
	for idx, key := range keys {
		value := values[idx]
		if value != nil {
			cache[key] = *value
		}
	}

	c.mutex.Lock()
	c.cache = cache
	c.mutex.Unlock()
}

// Get need to proxy only one method from interface
func (c *RedisDatabaseWithCache) Get(ctx context.Context, key string) (string, error) {
	var found bool
	var value string
	withLock(c.mutex.RLocker(), func() {
		value, found = c.cache[key]
	})

	if found {
		return value, nil
	}

	valueFromDB, err, _ := c.group.Do(key, func() (any, error) {
		var value string
		err := withRetries(ctx, func(ctx context.Context) error {
			var err error
			value, err = c.database.Get(ctx, key)
			return err
		})

		if err != nil {
			return "", err
		}

		withLock(&c.mutex, func() {
			c.cache[key] = value
		})

		return value, err
	})

	return valueFromDB.(string), err
}

func withLock(mutex sync.Locker, action func()) {
	if action == nil {
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	action()
}

func withRetries(ctx context.Context, action func(ctx context.Context) error) error {
	if action == nil {
		return errors.New("incorrect action")
	}

	var err error
	for idx := 1; idx <= retriesNumber; idx++ {
		ctx, cancel := context.WithTimeout(ctx, requestTimeout)
		defer cancel()

		err := action(ctx)
		if err == nil {
			return nil
		}

		time.Sleep(time.Duration(idx) * retriesInitialPauseDuration)
	}

	return err
}
