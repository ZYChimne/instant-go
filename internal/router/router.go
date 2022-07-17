package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"zychimne/instant/internal/api"
	"zychimne/instant/internal/util"
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
	// Auth
	authRouterGroup := r.Group("auth")
	authRouterGroup.POST("register", api.Register)
	authRouterGroup.POST("getToken", api.GetToken)
	// Instant
	instantRouterGroup := r.Group("instant").Use(auth())
	instantRouterGroup.GET("", api.GetInstants)
	instantRouterGroup.PUT("", api.UpdateInstant)
	instantRouterGroup.POST("", api.PostInstant)
	instantRouterGroup.POST("like", api.LikeInstant)
	instantRouterGroup.POST("share", api.ShareInstant)
	instantRouterGroup.GET("instants", api.GetInstantsByUserID)
	instantRouterGroup.GET("getLikesUsername", api.GetLikesUsername)
	// Chat
	chatRouterGroup := r.Group("chat")
	chatRouterGroup.GET("echo", api.Echo)
	// Comment
	commentRouterGroup := r.Group("comment").Use(auth())
	commentRouterGroup.GET("", api.GetComments)
	commentRouterGroup.POST("", api.PostComment)
	// commentRouterGroup.POST("like", api.LikeComment)
	// commentRouterGroup.POST("share", api.ShareComment)
	// Relation
	relationRouterGroup := r.Group("relation").Use(auth())
	relationRouterGroup.GET("followings", api.GetFollowings)
	relationRouterGroup.GET("followers", api.GetFollowings)
	relationRouterGroup.GET("potential", api.GetPotentialFollowings)
	relationRouterGroup.GET("all", api.GetAllUsers)
	relationRouterGroup.POST("", api.AddFollowing)
	relationRouterGroup.DELETE("", api.RemoveFollowing)
	// Profile
	profileRouterGroup := r.Group("profile").Use(auth())
	profileRouterGroup.GET("", api.GetUserInfo)
	r.Run(":8081")
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// log.Println("auth is running")
		token := c.GetHeader("Authentication")
		userID, err := util.VerifyJwt(token)
		if err != nil {
			log.Println(err.Error())
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("UserID", userID)
	}
}
