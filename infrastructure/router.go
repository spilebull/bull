package infrastructure

import (
	"server/ent"
	"server/middleware/errorhandler"
	"server/middleware/logger"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupServer(store *ent.Client) (*Server, error) {
	server := &Server{Store: store}
	r := gin.New()
	r.Use(logger.GinLogger(logger.Logger), logger.GinRecovery(logger.Logger, true))
	r.Use(errorhandler.HandleErrors())
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, //利用するURIに書き換えてください
		AllowMethods: []string{
			"POST",
			"GET",
			"PATCH",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	v1 := r.Group("sample/api/v1")
	{
		v1.GET("/hello", func(c *gin.Context) { c.JSON(200, gin.H{"message": "ハローワールド"}) })
		v1.GET("/date", func(c *gin.Context) { c.JSON(200, gin.H{"date": time.Now()}) })

		users := v1.Group("/users")
		{
			users.POST("", server.createUser)
			users.PUT("/:id", server.updateUser)
			users.GET("", server.getUserList)
			users.GET("/:id", server.getUser)
			users.DELETE("/:id", server.deleteUser)
		}
	}
	server.Router = r

	return server, nil
}
