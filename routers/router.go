package routers

import (
	_ "github.com/EDDYCJY/go-gin-example/docs"
	v1 "github.com/EDDYCJY/go-gin-example/routers/api/v1"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/EDDYCJY/go-gin-example/routers/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	//gin.SetMode(setting.RunMode)

	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("api/v1")
	{
		apiv1.GET("/tags", v1.GetTags)

		apiv1.POST("/tags", v1.AddTag)

		apiv1.PUT("/tags", v1.EditTag)

		apiv1.DELETE("/tags", v1.DeleteTag)

		apiv1.POST("/ansible", v1.AnsiblePlayExecutor)
	}

	//r.GET("/test", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "test",
	//	})
	//})

	return r
}
