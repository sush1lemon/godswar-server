package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/upper/db/v4/adapter/mysql"
	"godswar/config"
	"godswar/pkg/logger"
	"godswar/pkg/server"
	"time"
)

func main()  {

	var appConfig config.AppConfig
	appConfig.GetConfig()
	dbConfig := appConfig.Database

	dbSettings := mysql.ConnectionURL{
		User:     dbConfig.User,
		Password: dbConfig.Password,
		Database: dbConfig.Database,
		Host:     fmt.Sprintf("%s:%s", dbConfig.Host, dbConfig.Port),
	}

	db, err := mysql.Open(dbSettings)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	logger.BasicLog(fmt.Sprintf("MYSQL Connected to Host: [%s] at Port: [%s] using DB: [%s]", dbConfig.Host, dbConfig.Port, dbConfig.Database))
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	sv := server.Config{
		Host: appConfig.GameServer.Host,
		Port: appConfig.GameServer.Port,
		Type: appConfig.GameServer.Type,
	}
	gameServer := server.Server{Db: db, Config: sv}
	gameServer.NewServer()
}