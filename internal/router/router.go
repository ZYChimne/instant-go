package router

import (
	"time"
	"zychimne/instant/internal/api"
	database "zychimne/instant/internal/db"
	"zychimne/instant/middleware"

	"github.com/gin-gonic/gin"
	"github.com/zychimne/gin-cache"
	"github.com/zychimne/gin-cache/persist"
)

func Create(useCors bool) *gin.Engine {
	r := gin.Default()
	if useCors {
		r.Use(middleware.Cors())
	}
	r.Use(middleware.Error())
	localStore := persist.NewMemoryStore(1 * time.Minute)
	redisStore := persist.NewRedisStore(database.RedisClient)
	// V1 API
	v1RouterGroup := r.Group("v1")
	// Ping
	v1RouterGroup.GET("ping", api.Ping)
	// Account
	accountRouterGroup := v1RouterGroup.Group("account")
	accountRouterGroup.POST("", api.CreateAccount)
	accountRouterGroup.GET("check", cache.CacheByRequestURI(redisStore, 1*time.Minute), api.CheckIfAccountExists)
	accountRouterGroupWithAuth := accountRouterGroup.Group("").Use(middleware.Auth())
	accountRouterGroupWithAuth.DELETE("", api.DeleteAccount)
	accountRouterGroupWithAuth.GET("", api.GetAccount)
	accountRouterGroupWithAuth.GET("search", api.SearchAccounts)
	// Auth
	authRouterGroup := v1RouterGroup.Group("auth")
	authRouterGroup.POST("token", api.GetToken)
	// Feed
	feedRouterGroup := v1RouterGroup.Group("feed").Use(middleware.Auth())
	feedRouterGroup.GET("", api.GetFeed)
	// Instant
	instantRouterGroup := v1RouterGroup.Group("instant").Use(middleware.Auth())
	instantRouterGroup.GET("", api.GetInstants)
	instantRouterGroup.PUT("", api.UpdateInstant)
	instantRouterGroup.POST("", api.AddInstant)
	instantRouterGroup.DELETE("", api.DeleteInstant)
	// Like
	likeRouterGroup := v1RouterGroup.Group("like").Use(middleware.Auth())
	likeRouterGroup.POST("", api.Like)
	likeRouterGroup.DELETE("", api.Unlike)
	likeRouterGroup.GET("", api.GetLikes)
	// Comment
	commentRouterGroup := v1RouterGroup.Group("comment").Use(middleware.Auth())
	commentRouterGroup.GET("", api.GetComments)
	commentRouterGroup.POST("", api.AddComment)
	// commentRouterGroup.POST("like", api.LikeComment)
	// commentRouterGroup.POST("share", api.ShareComment)
	// Share
	shareRouterGroup := v1RouterGroup.Group("share").Use(middleware.Auth())
	shareRouterGroup.GET("", api.ShareInstant)
	// Chat
	chatRouterGroup := v1RouterGroup.Group("chat").Use(middleware.Auth())
	chatRouterGroup.Use(middleware.ServerSentEvent()).GET("", api.Receive)
	chatRouterGroup.POST("", api.Send)
	// Geo
	geoRouterGroup := v1RouterGroup.Group("geo")
	geoRouterGroup.GET("country", cache.CacheByRequestURI(localStore, 1*time.Minute), api.GetCountries)
	geoRouterGroup.GET("state", cache.CacheByRequestURI(redisStore, 1*time.Minute), api.GetStates)
	geoRouterGroup.GET("city", cache.CacheByRequestURI(redisStore, 1*time.Minute), api.GetCities)
	// Relation
	relationRouterGroup := v1RouterGroup.Group("relation").Use(middleware.Auth())
	relationRouterGroup.POST("", api.Follow)
	relationRouterGroup.DELETE("", api.Unfollow)
	relationRouterGroup.GET("following", api.GetFollowings)
	relationRouterGroup.GET("follower", api.GetFollowings)
	relationRouterGroup.GET("potential", api.GetPotentialFollowings)
	relationRouterGroup.GET("friend", api.GetFriends)
	return r
}
