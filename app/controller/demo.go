package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lovexiaoyue/gin-default/middleware"
)

type DemoController struct {

}

func DemoRegister(router *gin.RouterGroup)  {
	demo := DemoController{}
	router.GET("/index",demo.Index)
	router.POST("/index",demo.AddInfo)
}

func (demo *DemoController) Index(c *gin.Context)  {
	middleware.ResponseSuccess(c,"dada")
	err := "dsadsadsa"
	middleware.Log.Errorf("read config failed :%v", err)
	return
}

func (demo *DemoController) AddInfo(c *gin.Context)  {
	middleware.ResponseSuccess(c,"success")
	return
}