package main

import "sync"

func lockAnyTimes() {
	mutex := sync.Mutex{}
	mutex.Lock()
	mutex.Lock()
}

func unlockWithoutLock() {
	mutex := sync.Mutex{}
	mutex.Unlock()
}

func unlockFromAnotherGoroutine() {
	mutex := sync.Mutex{}
	mutex.Lock()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		mutex.Unlock()
	}()

	wg.Wait()

	mutex.Lock()
	mutex.Unlock()
}

func RUnlockLockedMutex() {
	m := sync.RWMutex{}
	m.Lock()
	m.RUnlock()
}

func UnlockRLockedMutex() {
	m := sync.RWMutex{}
	m.RLock()
	m.Unlock()
}

func LockRLockedMutex() {
	m := sync.RWMutex{}
	m.Lock()
	m.RLock()
}

func main() {
}
