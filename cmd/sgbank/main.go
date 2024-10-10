package main

import (
	"database/sql"
	"log"

	"github.com/NhutHuyDev/sgbank/internal/infra/db"
	"github.com/NhutHuyDev/sgbank/internal/rest"
	"github.com/NhutHuyDev/sgbank/pkg/utils"
	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".", "app")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server, err := rest.NewServer(config, store)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = server.StartServer(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
