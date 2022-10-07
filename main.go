package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	api "github.com/mr-od/Asusu-Igbo/api"
	db "github.com/mr-od/Asusu-Igbo/db/sqlc"
	"github.com/mr-od/Asusu-Igbo/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Loading config file failed:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
