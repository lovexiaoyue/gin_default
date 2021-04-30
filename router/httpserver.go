package router

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

var (
	HttpSrvHandler *http.Server
)

func HttpServerRun()  {

	r := InitRouter()
	httpPort := viper.GetString("http.port")
	readTimeout := time.Duration(viper.GetInt("http.read_timeout")) * time.Second
	writeTimeout := time.Duration(viper.GetInt("http.write_timeout")) * time.Second
	maxHeader := uint(viper.GetInt("http.max_header"))

	HttpSrvHandler = &http.Server{
		Addr:              fmt.Sprintf(":%s",httpPort),
		Handler:           r,
		ReadTimeout:       readTimeout,
		WriteTimeout:      writeTimeout,
		MaxHeaderBytes:    1 << maxHeader,
	}

	go func() {
		log.Printf(" [INFO] HttpServerRun:%s\n",httpPort)
		if err := HttpSrvHandler.ListenAndServe(); err != nil{
			log.Fatalf(" [ERROR] HttpServerRun:%s err:%v\n",httpPort,err)
		}
	}()

}

func HttpServerStop()  {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpSrvHandler.Shutdown(ctx); err !=nil{
		log.Fatalf(" [ERROR] HttpServerStop err:%v\n",err)
	}
	log.Printf(" [INFO] HttpServerStop stopped]\n")
}