package model

import (
	"fmt"
	"alert/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type Alert struct {
	ID        uint `gorm:"primary_key"`
	ChannelID int
	CreatedAt time.Time
	UpdatedAt time.Time
	StartAt time.Time
	DeletedAt *time.Time
	AlertName string
	Resource string
	Severity string
	Summary string
	Type string
	Description string
}

type Channel struct {
	ID int `gorm:"primary_key"`
	Service string
	ChartID string
}

type DataBase struct {
	database *gorm.DB
}

var Alertdb=&DataBase{}

func(db *DataBase) ConnectionDB() {
	var err error
	dbName:=config.Configmanager.Secret.Mysql.Connections.Database
	user:=config.Configmanager.Secret.Mysql.Connections.User
	password:=config.Configmanager.Secret.Mysql.Connections.Passwd
	host:=config.Configmanager.Secret.Mysql.Connections.Host
	connstr:=fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		user,
		password,
		host,
		dbName)

	gConfig := &gorm.Config{}
	if config.Configmanager.Secret.Mysql.Logmod {
		gConfig.Logger = logger.Default.LogMode(logger.Info)
	} else {
		gConfig.Logger = logger.Default.LogMode(logger.Silent)
	}

	db.database,err= gorm.Open(mysql.Open(connstr),gConfig)
	if err !=nil{
		panic(err)
	}

	sqlDB,_:=db.database.DB()
	sqlDB.SetMaxOpenConns(config.Configmanager.Secret.Mysql.MaxOpenConns)
	sqlDB.SetMaxIdleConns(config.Configmanager.Secret.Mysql.MaxIdleConns)
}
