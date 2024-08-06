package main

import (
    "database/sql"
    "log"
    "fmt"

    "github.com/techschool/simplebank/util"

    _ "github.com/lib/pq"
    "github.com/techschool/simplebank/api"
    db "github.com/techschool/simplebank/db/sqlc"
)

func main() {
    config, err := util.LoadConfig(".")
    if err != nil {
        log.Fatal("cannot load config:", err)
    }

    fmt.Printf("CONFIG", config.DBSource);

    conn, err := sql.Open(config.DBDriver, config.DBSource)
    if err != nil {
        log.Fatal("cannot connect to db:", err)
    }

    store := db.NewStore(conn)
    server, err := api.NewServer(config, store)
    if err != nil {
        log.Fatal("cannot create server", err)
    }

    err = server.Start(config.ServerAddress)
    if err != nil {
        log.Fatal("cannot start server:", err)
    }
}