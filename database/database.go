package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

type Db struct{
	Conn *pgx.Conn
}

func (db *Db) CreateConnection(ctx context.Context) error{
	conn, err := pgx.Connect(ctx,os.Getenv("DB_URL"))

	db.Conn = conn
	return err
}

