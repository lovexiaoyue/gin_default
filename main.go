package main

import (
	"github.com/lovexiaoyue/gin-default/config"
	"github.com/lovexiaoyue/gin-default/router"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main()  {
	// 配置初始化
	err := config.InitViper("dev")
	if err != nil{
		log.Fatal("read config failed :%v", err)
	}
	router.HttpServerRun()
	quit := make(chan os.Signal)
	signal.Notify(quit,syscall.SIGKILL,syscall.SIGQUIT,syscall.SIGINT,syscall.SIGTERM)
	<- quit
	router.HttpServerStop()
}