package router

import (
	"time"
	"zychimne/instant/internal/api/auth"
	"zychimne/instant/internal/api/chat"
	"zychimne/instant/internal/api/comment"
	"zychimne/instant/internal/api/instant"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	authRouterGroup.POST("register", apiauth.Register)
	authRouterGroup.POST("getToken", apiauth.GetToken)
	// Instant
	instantRouterGroup := r.Group("instant")
	instantRouterGroup.POST("get", apiinstant.GetInstants)
	instantRouterGroup.POST("update", apiinstant.UpdateInstant)
	instantRouterGroup.POST("post", apiinstant.PostInstant)
	instantRouterGroup.POST("like", apiinstant.LikeInstant)
	instantRouterGroup.POST("share", apiinstant.ShareInstant)
	// Chat
	chatRouterGroup := r.Group("chat")
	chatRouterGroup.GET("echo", apichat.Echo)
	// Comment
	commentRouterGroup := r.Group("comment")
	commentRouterGroup.POST("get", apicomment.GetComments)
	commentRouterGroup.POST("post", apicomment.PostComment)
	commentRouterGroup.POST("like", apicomment.LikeComment)
	commentRouterGroup.POST("share", apicomment.ShareComment)
	r.Run(":8081")
}
