package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"godswar/config"
	"godswar/pkg/logger"
	"godswar/pkg/server"
	"time"
)

func main()  {

	var appConfig config.AppConfig
	appConfig.GetConfig()

	dbConfig := appConfig.Database
	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database))
	if err != nil {
		panic(err)
	}
	logger.BasicLog(fmt.Sprintf("MYSQL Connected to Host: [%s] at Port: [%s] using DB: [%s]", dbConfig.Host, dbConfig.Port, dbConfig.Database))
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	sv := server.ServerConfig{
		Host: appConfig.GameServer.Host,
		Port: appConfig.GameServer.Port,
		Type: appConfig.GameServer.Type,
	}
	gameServer := server.Server{Db: *db, Config: sv}
	gameServer.NewServer()
}