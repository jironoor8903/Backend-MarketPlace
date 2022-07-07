package main

import (
	"database/sql"
	"entryleveltask/api"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/go-sql-driver/mysql"
)

func main() {

	cfg := mysql.Config{
		User:   "root",
		Passwd: "insert password here",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "//insert name",
	}
	// Get a database handle.
	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	router := gin.Default()
	api.InitApi(router, db)
	router.Run()

	//next task: figure out how to connect database in service

}
