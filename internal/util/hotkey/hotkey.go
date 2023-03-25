package hotkey

import (
	"sync"
	"time"
	"zychimne/instant/internal/util/lru"
	"zychimne/instant/internal/util/topk"
)

type HotKeyCache struct {
	blackList       map[string]int64
	whiteList       map[string]int64
	mutex           sync.Mutex
	heavyKeeper     *topK.HeavyKeeper
	lruCache        *lru.Cache
	cacheExpireTime int64
}

type LRUItem struct {
	item      any
	timestamp int64
}

func NewHotKeyCache(k int, rows uint64, cols uint64, decay float64, minFreq int, cacheExpireTime int64) *HotKeyCache {
	return &HotKeyCache{
		blackList:       map[string]int64{},
		whiteList:       map[string]int64{},
		heavyKeeper:     topK.NewHeavyKeeper(k, rows, cols, decay, minFreq),
		lruCache:        lru.New(minFreq),
		cacheExpireTime: cacheExpireTime,
	}
}

func (h *HotKeyCache) GetWhiteList() map[string]int64 {
	return h.whiteList
}

func (h *HotKeyCache) GetBlackList() map[string]int64 {
	return h.blackList
}

func (h *HotKeyCache) InWhiteList(key string) bool {
	_, ok := h.whiteList[key]
	return ok
}

func (h *HotKeyCache) InBlackList(key string) bool {
	_, ok := h.blackList[key]
	return ok
}

func (h *HotKeyCache) IsHotKey(key string) bool {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	if h.InWhiteList(key) {
		return true
	}
	if h.InBlackList(key) {
		return false
	}
	return h.heavyKeeper.Find(key)
}

func (h *HotKeyCache) AddToWhiteList(key string, ttl int64) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.whiteList[key] = ttl
}

func (h *HotKeyCache) AddToBlackList(key string, ttl int64) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.blackList[key] = ttl
}

func (h *HotKeyCache) RemoveFromWhiteList(key string) bool {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	if h.InWhiteList(key) {
		delete(h.whiteList, key)
		return true
	}
	return false
}

func (h *HotKeyCache) RemoveFromBlackList(key string) bool {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	if h.InBlackList(key) {
		delete(h.blackList, key)
		return true
	}
	return false
}

func (h *HotKeyCache) Add(key string, item any) bool {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	if evict, ok := h.heavyKeeper.Add(key); ok {
		h.lruCache.Add(key, LRUItem{item, time.Now().Unix()})
		if evict != "" {
			h.lruCache.Remove(evict)
		}
		return true
	}
	return false
}

func (h *HotKeyCache) Get(key string) (any, bool) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	if item, ok := h.lruCache.Get(key); ok {
		if item.(LRUItem).timestamp+h.cacheExpireTime > time.Now().Unix() {
			return item.(LRUItem).item, true
		} else {
			h.lruCache.Remove(key)
		}
	}
	return nil, false
}
