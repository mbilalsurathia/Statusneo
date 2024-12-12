package rest

import (
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"maker-checker/service"
	"net/http"
	"time"
)

func StartServer(container *service.Container) *http.Server {
	/*
	 */
	gin.SetMode(gin.DebugMode)
	userController := NewUserController(container.UserService)
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.ForwardedByClientIP = true

	api := r.Group("/api/v1")
	{
		user := api.Group("/message-request")
		{
			//paths
			user.POST("/", userController.CreateMessage)
			user.PATCH("/", userController.UpdateMessage)
			user.GET("/", userController.GetMessages)
		}

	}

	// Setup routes and middleware...
	srv := &http.Server{
		Addr:    container.GbeConfigService.GetConfig().Rest.Addr,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	return srv
}
