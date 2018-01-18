package main

import (
	"sync"
	"time"

	"github.com/dop251/goja"
)

type CacheMgr struct {
	mu            sync.RWMutex
	gcIntervalSec int64
	items         map[string]*CacheItem
}

type CacheItem struct {
	value         interface{}
	expireTimeSec int64
}

func NewCacheMgr(gcIntervalSec int64) *CacheMgr {
	mgr := &CacheMgr{
		gcIntervalSec: gcIntervalSec,
		items:         make(map[string]*CacheItem),
	}
	go mgr.gc()
	return mgr
}

func (mgr *CacheMgr) gc() {
	for {
		<-time.After(time.Duration(mgr.gcIntervalSec) * time.Second)
		mgr.mu.Lock()
		for key := range mgr.items {
			mgr.isExpired(key)
		}
		mgr.mu.Unlock()
	}
}

func (mgr *CacheMgr) isExpired(key string) bool {
	if item, ok := mgr.items[key]; ok {
		if item.expireTimeSec <= 0 {
			return false
		}
		if item.expireTimeSec < time.Now().Unix() {
			delete(mgr.items, key)
			return true
		}
		return false
	}
	return true
}

func (mgr *CacheMgr) Add(key string, v int64, expireSec int64) int64 {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()

	expireTimeSec := int64(-1)
	if expireSec > 0 {
		expireTimeSec = time.Now().Unix() + expireSec
	}

	if item, ok := mgr.items[key]; ok {
		if oldVal, ok := item.value.(int64); ok {
			item.expireTimeSec = expireTimeSec
			item.value = oldVal + v
			return oldVal
		}
	}
	mgr.items[key] = &CacheItem{
		value:         v,
		expireTimeSec: expireTimeSec,
	}
	return 0
}

func (mgr *CacheMgr) Set(key string, value interface{}, expireSec int64) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()

	expireTimeSec := int64(-1)
	if expireSec > 0 {
		expireTimeSec = time.Now().Unix() + expireSec
	}
	mgr.items[key] = &CacheItem{
		value:         value,
		expireTimeSec: expireTimeSec,
	}
}

func (mgr *CacheMgr) Get(key string) (interface{}, bool) {
	mgr.mu.RLock()
	defer mgr.mu.RUnlock()
	if !mgr.isExpired(key) {
		if v, ok := mgr.items[key]; ok {
			return v.value, true
		}
	}
	return nil, false
}

func (mgr *CacheMgr) Del(key string) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	delete(mgr.items, key)
}

func (mgr *CacheMgr) Flush() {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	mgr.items = make(map[string]*CacheItem)
}

///////////////////////////////////////////////////////////////////////////

func NewCache(runtime *goja.Runtime, cacheMgr *CacheMgr) *goja.Object {
	o := runtime.NewObject()
	o.Set("set", func(call goja.FunctionCall) goja.Value {
		key := call.Argument(0).String()
		value := call.Argument(1).Export()
		expireSec := call.Argument(2).ToInteger()
		if IsValidType(value) {
			cacheMgr.Set(key, value, expireSec)
			return nil
		}
		panic(runtime.NewTypeError("value type %T is not permitted", value))
	})

	o.Set("get", func(call goja.FunctionCall) goja.Value {
		key := call.Argument(0).String()
		if value, ok := cacheMgr.Get(key); ok {
			return runtime.ToValue(value)
		}
		return goja.Null()
	})

	o.Set("del", func(call goja.FunctionCall) goja.Value {
		key := call.Argument(0).String()
		cacheMgr.Del(key)
		return nil
	})

	o.Set("flush", func(call goja.FunctionCall) goja.Value {
		cacheMgr.Flush()
		return nil
	})

	o.Set("add", func(call goja.FunctionCall) goja.Value {
		key := call.Argument(0).String()
		value := call.Argument(1).ToInteger()
		expireSec := call.Argument(2).ToInteger()
		oldValue := cacheMgr.Add(key, value, expireSec)
		return runtime.ToValue(oldValue)
	})

	o.Set("sub", func(call goja.FunctionCall) goja.Value {
		key := call.Argument(0).String()
		value := call.Argument(1).ToInteger()
		expireSec := call.Argument(2).ToInteger()
		oldValue := cacheMgr.Add(key, -value, expireSec)
		return runtime.ToValue(oldValue)
	})

	return o
}
