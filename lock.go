package utils

import "sync"

func LockRun(lock sync.Locker, handler func()) {
	lock.Lock()
	handler()
	lock.Unlock()
}
