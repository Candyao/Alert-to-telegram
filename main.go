package main

import (
	"alert/config"
	"alert/model"
	"alert/routers"
)

func main() {
	config.Configmanager.LoadConfig(nil,nil)
	model.Alertdb.ConnectionDB()
	r:=routers.InitRouter()
	_=r.Run(config.Configmanager.Config.Service.Port)
}