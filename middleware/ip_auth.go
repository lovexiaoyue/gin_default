package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)



func IPAuthMiddleware() gin.HandlerFunc  {
	return func(c *gin.Context) {
		isMatched := false
		for host,_ := range viper.GetStringMap("whitelist"){
			if c.ClientIP() == host {
				isMatched = true
			}
		}

		if !isMatched{
			ResponseError(c,InternalErrorCode,errors.New(fmt.Sprintf("%v, not in iplist",c.ClientIP())))
			c.Abort()
			return
		}
		c.Next()
	}

}
