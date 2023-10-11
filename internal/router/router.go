package router

import (
	"log"
	"net/http"
	"strings"
	"time"
	"zychimne/instant/internal/api"
	database "zychimne/instant/internal/db"
	"zychimne/instant/internal/util"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/zychimne/gin-cache"
	"github.com/zychimne/gin-cache/persist"
)

func Create() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authentication"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))
	redisStore := persist.NewRedisStore(database.RedisClient)
	// V1 API
	v1RouterGroup := r.Group("v1")
	// Auth
	authRouterGroup := v1RouterGroup.Group("auth")
	authRouterGroup.GET("ping", api.Ping)
	authRouterGroup.POST("register", api.Register)
	authRouterGroup.POST("token", api.GetToken)
	// Instant
	instantRouterGroup := v1RouterGroup.Group("instant").Use(auth())
	instantRouterGroup.GET("", api.GetInstants)
	instantRouterGroup.PUT("", api.UpdateInstant)
	instantRouterGroup.POST("", api.PostInstant)
	instantRouterGroup.POST("like", api.LikeInstant)
	instantRouterGroup.POST("share", api.ShareInstant)
	instantRouterGroup.GET("instants", api.GetInstantsByUserID)
	instantRouterGroup.GET("like/users", api.GetLikesUserInfo)
	// Chat
	chatRouterGroup := v1RouterGroup.Group("chat")
	chatRouterGroup.GET("echo", api.Echo)
	// chatRouterGroup.GET("history", api.GetChatHistory)
	chatRouterGroup.POST("", api.Chat)
	// Comment
	commentRouterGroup := v1RouterGroup.Group("comment").Use(auth())
	commentRouterGroup.GET("", api.GetComments)
	commentRouterGroup.POST("", api.PostComment)
	// commentRouterGroup.POST("like", api.LikeComment)
	// commentRouterGroup.POST("share", api.ShareComment)
	// Geo
	geoRouterGroup := v1RouterGroup.Group("geo")
	geoRouterGroup.GET("countries", cache.CacheByRequestURI(redisStore, 100*time.Second), api.GetCountries)
	geoRouterGroup.GET("states", cache.CacheByRequestURI(redisStore, 100*time.Second), api.GetStates)
	geoRouterGroup.GET("cities", cache.CacheByRequestURI(redisStore, 100*time.Second), api.GetCities)
	// Relation
	relationRouterGroup := v1RouterGroup.Group("relation").Use(auth())
	relationRouterGroup.GET("followings", api.GetFollowings)
	relationRouterGroup.GET("followers", api.GetFollowings)
	relationRouterGroup.GET("potential", api.GetPotentialFollowings)
	relationRouterGroup.GET("friends", api.GetFriends)
	relationRouterGroup.POST("", api.AddFollowing)
	relationRouterGroup.DELETE("", api.RemoveFollowing)
	// Profile
	profileRouterGroup := v1RouterGroup.Group("profile").Use(auth())
	profileRouterGroup.GET("", api.GetUserProfile)
	profileRouterGroup.GET("detail", api.GetUserProfileDetail)
	profileRouterGroup.GET("query", api.QueryUsers)
	r.Run(":8081")
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if !strings.HasPrefix(token, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid Token"})
			return
		}
		userID, err := util.VerifyJwt(token[7:])
		if err != nil {
			log.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid Token"})
			return
		}
		c.Set("UserID", userID)
	}
}
