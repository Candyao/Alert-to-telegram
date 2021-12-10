package api

import (
	"alert/model"
	"alert/request"
	"alert/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GCPAlert(ctx *gin.Context) {
	var notification *request.Notification
	var data *model.GcpData
	var repG=&util.GinSelf{Ctx:ctx}

	if err:=ctx.ShouldBindJSON(&notification);err!=nil{
		repG.Response(http.StatusInternalServerError, util.ERROR_BIND_JSON,nil)
		log.Println(err)
		return
	}

	if notification.State=="open"{
		data=&model.GcpData{
			Notify: notification,
		}
		err:=data.AddGcpAlert()
		if err !=nil{
			repG.Response(http.StatusInternalServerError,util.ERROR_SEND_MSG,nil)
			log.Println(err)
			return
		}
	}

	msg:=util.AlertFormatTemplate(notification,util.LoadTemplate("gcp.tmpl"))
	if err:=util.Sendmsg(msg,"");err !=nil{
		repG.Response(http.StatusInternalServerError,util.ERROR_SEND_MSG,nil)
		log.Println(err)
		return
	}
	repG.Response(http.StatusOK,util.SUCCESS,nil)
}