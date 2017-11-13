package core

import (
	. "github.com/hazelcast/go-client/serialization"
	"time"
)

type IMap interface {
	IDistributedObject
	Put(key interface{}, value interface{}) (oldValue interface{}, err error)
	Get(key interface{}) (value interface{}, err error)
	Remove(key interface{}) (value interface{}, err error)
	RemoveIfSame(key interface{}, value interface{}) (ok bool, err error)
	Size() (size int32, err error)
	ContainsKey(key interface{}) (found bool, err error)
	ContainsValue(value interface{}) (found bool, err error)
	Clear() (err error)
	Delete(key interface{}) (err error)
	IsEmpty() (empty bool, err error)
	AddIndex(attributes *string, ordered bool) (err error)
	Evict(key interface{}) (evicted bool, err error)
	EvictAll() (err error)
	Flush() (err error)
	ForceUnlock(key interface{}) (err error)
	Lock(key interface{}) (err error)
	LockWithLeaseTime(key interface{}, lease int64, leaseTimeUnit time.Duration) (err error)
	Unlock(key interface{}) (err error)
	IsLocked(key interface{}) (locked bool, err error)
	Replace(key interface{}, value interface{}) (oldValue interface{}, err error)
	ReplaceIfSame(key interface{}, oldValue interface{}, newValue interface{}) (replaced bool, err error)
	Set(key interface{}, value interface{}) (err error)
	SetWithTtl(key interface{}, value interface{}, ttl int64, ttlTimeUnit time.Duration) (err error)
	PutIfAbsent(key interface{}, value interface{}) (oldValue interface{}, err error)
	PutAll(mp *map[interface{}]interface{}) (err error)
	EntrySet() (resultPairs []IPair, err error)
	EntrySetWithPredicate(predicate IPredicate) (resultPairs []IPair, err error)
	TryLock(key interface{}) (locked bool, err error)
	TryLockWithTimeout(key interface{}, timeout int64, timeoutTimeUnit time.Duration) (locked bool, err error)
	TryLockWithTimeoutAndLease(key interface{}, timeout int64, timeoutTimeUnit time.Duration, lease int64, leaseTimeUnit time.Duration) (locked bool, err error)
	TryPut(key interface{}, value interface{}) (ok bool, err error)
	TryRemove(key interface{}, timeout int64, timeoutTimeUnit time.Duration) (ok bool, err error)
	GetAll(keys []interface{}) (entryPairs []IPair, err error)
	GetEntryView(key interface{}) (entryView IEntryView, err error)
	PutTransient(key interface{}, value interface{}, ttl int64, ttlTimeUnit time.Duration) (err error)
	AddEntryListener(listener interface{}, includeValue bool) (registrationID *string, err error)
	AddEntryListenerToKey(listener interface{}, key interface{}, includeValue bool) (registrationID *string, err error)
	RemoveEntryListener(registrationId *string) (removed bool, err error)
	ExecuteOnKey(key interface{}, entryProcessor interface{}) (result interface{}, err error)
	ExecuteOnKeys(keys []interface{}, entryProcessor interface{}) (keyToResultPairs []IPair, err error)
	ExecuteOnEntries(entryProcessor interface{}) (keyToResultPairs []IPair, err error)
}
