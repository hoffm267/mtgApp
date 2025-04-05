package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func InitDB(dbUrl string) error {
	var err error

	Conn, err = pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		return err
	}
	return nil
}

// var Test string = "postgres://postgres:1243@localhost:5432/mtgApp"
// Conn, err error = pgx.Connect(context.Background(), Test)
// if err != nil {
// 	fmt.Print(err)
// }
