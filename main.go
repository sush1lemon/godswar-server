package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"godswar/logger"
	"godswar/networking"
	"net"
	"os"
	"time"
)

func main() {
	var config config
	config.getConfig()
	// Listen for incoming connections.
	dbConfig := config.Database
	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database))
	if err != nil {
		panic(err)
	}
	logger.BasicLog(fmt.Sprintf("MYSQL Connected to Host: [%s] at Port: [%s] using DB: [%s]", dbConfig.Host, dbConfig.Port, dbConfig.Database))
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	l, err := net.Listen(config.Server.Type, fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	// Close the listener when the application closes.
	defer l.Close()
	logger.BasicLog("Listening on " + config.Server.Host + ":" + config.Server.Port)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		network := networking.NewNetwork(conn)
		go network.HandleConnection()
	}

	fmt.Println("asdasdasd")
}
