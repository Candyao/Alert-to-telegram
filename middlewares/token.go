package middlewares

import (
	"alert/config"
	"alert/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		repG:=&util.GinSelf{Ctx:c}
		token:=c.Query("token")
		if token != config.Configmanager.Config.Service.Token {
			repG.Response(http.StatusUnauthorized,util.INVALID_PARAMS,nil)
			c.Abort()
			return
		}
		log.Println("get message success")
		c.Next()
	}
}