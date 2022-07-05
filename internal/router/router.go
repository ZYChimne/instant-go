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
	// Chat
	chatRouterGroup := r.Group("chat")
	chatRouterGroup.GET("echo", api.Echo)
	// Comment
	commentRouterGroup := r.Group("comment").Use(auth())
	commentRouterGroup.GET("get", api.GetComments)
	commentRouterGroup.POST("post", api.PostComment)
	// commentRouterGroup.POST("like", api.LikeComment)
	// commentRouterGroup.POST("share", api.ShareComment)
	// Friend
	friendRouterGroup:=r.Group("friend").Use(auth())
	friendRouterGroup.GET("get", api.GetFollowings)
	friendRouterGroup.POST("add", api.AddFollowing)
	friendRouterGroup.DELETE("remove", api.RemoveFollowing)
	friendRouterGroup.GET("potential", api.GetPotentialFollowing)
	// Profile
	profileRouterGroup:=r.Group("profile").Use(auth())
	profileRouterGroup.GET("get", api.GetUserInfo)
	r.Run(":8081")
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("auth is running")
		token := c.GetHeader("Authentication")
		userID, err := util.VerifyJwt(token)
		if err != nil {
			log.Println(err.Error())
			c.AbortWithStatus(http.StatusUnauthorized);
			return
		}
		c.Set("UserID", userID)
	}
}
