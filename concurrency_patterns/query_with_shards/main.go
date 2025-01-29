package main

import (
	"context"
	"sync"

	"golang.org/x/sync/errgroup"
)

type Database interface {
	Query(query string) (string, error) // simple interface for example
}

func DistributedQuery(shards []Database, query string) ([]string, error) {
	var mutex sync.Mutex
	responses := make([]string, 0, len(shards))
	group, ctx := errgroup.WithContext(context.TODO()) // nice to have parent context

	for _, shard := range shards {
		group.Go(func() error {
			type result struct {
				response string
				err      error
			}

			resultCh := make(chan result, 1)
			go func() {
				response, err := shard.Query(query)
				resultCh <- result{response: response, err: err}
			}()

			select {
			case <-ctx.Done():
				return ctx.Err()
			case result := <-resultCh:
				if result.err != nil {
					return result.err
				}

				mutex.Lock()
				responses = append(responses, result.response)
				mutex.Unlock()

				return nil
			}
		})
	}

	if err := group.Wait(); err != nil {
		return nil, err
	} else {
		return responses, nil
	}
}

func main() {
	shards := []*ClickHouseDatabase{
		NewClickHouseDatabase("127.0.0.1:5432"),
		NewClickHouseDatabase("127.0.0.2:5432"),
		NewClickHouseDatabase("127.0.0.3:5432"),
	}

	response, err := DistributedQuery(shards, "query to clickhouse...")
	_ = response
	_ = err
}
