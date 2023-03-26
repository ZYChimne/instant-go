package api

import (
	"zychimne/instant/internal/util/hotkey"
)

const k int = 1000
const rows uint64 = 2
const cols uint64 = 0x10000
const decay float64 = 1.08
const minFreq int = 100

var CacheExpireTime int64 = 60 // unit: seconds

var LikeInstantCache = hotkey.NewHotKeyCache(k, rows, cols, decay, minFreq, CacheExpireTime, onLikeInstantEvicted)
var SimpleUserCache = hotkey.NewHotKeyCache(k, rows, cols, decay, minFreq, CacheExpireTime, nil)
var UserCache = hotkey.NewHotKeyCache(k, rows, cols, decay, minFreq, CacheExpireTime, nil)
var FriendsCache = hotkey.NewHotKeyCache(k, rows, cols, decay, minFreq, CacheExpireTime, nil)
