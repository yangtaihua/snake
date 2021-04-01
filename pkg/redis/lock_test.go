package redis

import (
	"testing"
	"time"
)

func TestLockWithDefaultTimeout(t *testing.T) {
	InitTestRedis()

	lock := NewLock(RedisClient, "test:lock")
	ok, err := lock.Lock()
	if err != nil {
		t.Error(err)
	}

	err = lock.Unlock()
	if err != nil {
		t.Error(err)
	}

	t.Log(ok)
}

func TestLockWithTimeout(t *testing.T) {
	InitTestRedis()
	lock := NewLock(RedisClient, "test:lock", Timeout(3*time.Second))
	ok, err := lock.Lock()
	if err != nil {
		t.Error(err)
	}

	err = lock.Unlock()
	if err != nil {
		t.Error(err)
	}

	t.Log(ok)
}
