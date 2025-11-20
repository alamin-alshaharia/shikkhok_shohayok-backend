package main

import (
	"database/sql"
	"log"
	"shikkhok_shohayok/routes"
	"shikkhok_shohayok/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	if err = conn.Ping(); err != nil {
		log.Fatal("cannot ping db:", err)
	}
	log.Println("successfully connected to db")

	router := routes.SetupRouter(conn, config)
	router.Run(config.ServerAddress)
}
