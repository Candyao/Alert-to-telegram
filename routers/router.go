package routers

import (
	"alert/middlewares"
	"alert/routers/api"
	"github.com/gin-gonic/gin"
)

func  InitRouter() *gin.Engine{
	r:=gin.New()
	v1:=r.Group("/api/alert")
	v1.Use(middlewares.CheckToken())
	{
		v1.POST("/gcp",api.GCPAlert)
		v1.POST("/prom",api.PromAlert)
	}
	return r
}