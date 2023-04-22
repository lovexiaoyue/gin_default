package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lovexiaoyue/gin-default/app/controller"
	"github.com/lovexiaoyue/gin-default/docs"
	"github.com/lovexiaoyue/gin-default/middleware"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {


	gin.SetMode(viper.GetString("run_mode"))
	// swagger info
	docs.SwaggerInfo.Title = viper.GetString("swagger.title")
	docs.SwaggerInfo.Description = viper.GetString("swagger.desc")
	docs.SwaggerInfo.Version = viper.GetString("swagger.version")
	docs.SwaggerInfo.Host = viper.GetString("swagger.host")
	docs.SwaggerInfo.BasePath = viper.GetString("swagger.base_path")
	docs.SwaggerInfo.Schemes = []string{"http","https"}

	router := gin.Default()
	router.Use(middlewares...)
	router.Use(middleware.LoggerMiddleware(),middleware.IPAuthMiddleware())
	router.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/demo")
	v1.Use(middleware.TranslationMiddleware())
	{
		controller.DemoRegister(v1)
	}
	return router
}
