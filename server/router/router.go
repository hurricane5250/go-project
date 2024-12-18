package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hurricane5250/MyGoProject/server/api/v1"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", v1.BaseApi.Index)

	api := server.Group("/v1")
	userApi := v1.UserApi
	{
		api.GET("/users/:id", userApi.GetById) //文章详情
	}
}
