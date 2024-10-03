package main

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"log"
	"servCliGrantex/grpcServ/config"
)

const (
	dialect     = "pgx"
	fmtDbString = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"
)

func main() {
	c := config.NewDB()
	dbString := fmt.Sprintf(fmtDbString, c.Host, c.Username, c.Password, c.DBName, c.Port)

	fmt.Println("starting migration")
	db, err := goose.OpenDBWithDriver(dialect, dbString)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = db.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	dir := "app/migrationssql"

	fmt.Println("migrating")
	err = goose.RunContext(context.Background(), "up", db, dir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("finished")

}
