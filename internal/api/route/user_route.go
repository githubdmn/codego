package route

import (
	"github.com/gin-gonic/gin"
	"github.com/githubdmn/codego/ent"
	"github.com/githubdmn/codego/internal/api/handler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup(r *gin.Engine, client *ent.Client) {
	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API v1
	v1 := r.Group("/api/v1")

	// User routes
	userHandler := handler.NewUserHandler(client)
	users := v1.Group("/users")
	{
		users.GET("", userHandler.GetUsers)
	}
}
