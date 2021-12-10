package api

import (
	"alert/model"
	"alert/request"
	"alert/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func PromAlert(ctx *gin.Context)  {
	var alerts *request.Alerts
	var repG=&util.GinSelf{Ctx:ctx}

	if err:= ctx.ShouldBindJSON(&alerts);err!=nil{
		repG.Response(http.StatusOK,util.ERROR_BIND_JSON,nil)
		log.Println(err)
		return
	}

	if alerts.Status == "firing" {
		err:=model.AddPromAlert(alerts)
		if err!=nil{
			repG.Response(http.StatusInternalServerError,util.ERROR_ADD_AlERT_FAIL,nil)
			log.Println(err)
			return
		}
	}

	svc:=alerts.Alerts[0].Labels["service"]
	msg:=util.AlertFormatTemplate(alerts,util.LoadTemplate("prom.tmpl"))
	var chatID string

	if svc!=nil{
		chatID=model.GetChatIDByService(svc.(string))
	}
	if err:=util.Sendmsg(msg,chatID);err !=nil{
		repG.Response(http.StatusInternalServerError, util.ERROR_SEND_MSG,nil)
		log.Println(err)
		return
	}
	repG.Response(http.StatusOK,util.SUCCESS,nil)
}
