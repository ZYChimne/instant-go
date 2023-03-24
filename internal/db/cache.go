package database

import (
	"zychimne/instant/internal/util/lru"
)

const maxUserCacheSize = 128

var UserCache = lru.New(maxUserCacheSize)
