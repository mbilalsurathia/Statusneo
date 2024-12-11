package rest

import (
	"io"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	addr           string
	userController UserController
}

// NewHttpServer create server instance
func NewHttpServer(
	addr string,
	userController UserController,

) *HttpServer {
	return &HttpServer{
		addr:           addr,
		userController: userController,
	}
}

func (server *HttpServer) Start() {
	gin.SetMode(gin.DebugMode)
	gin.DefaultWriter = io.Discard

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

	v1 := r.Group("/api/v1")
	{
		api := v1.Group("")

		user := api.Group("/users")
		{
			//paths
			user.POST("/message-request", server.userController.CreateMessage)
			user.PATCH("/message-request", server.userController.UpdateMessage)
			user.GET("/message-request", server.userController.GetMessages)
		}

	}

	err := r.Run(server.addr)
	if err != nil {
		panic(err)
	}

}
