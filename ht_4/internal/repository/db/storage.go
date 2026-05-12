package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	conn *pgxpool.Pool
}

func New() *Storage { //New(dbDSN string)
	conn, err := pgxpool.New(context.TODO(), "postgres://postgres:postgresql_pass@localhost:5432/ToDoList?sslmode=disable") //New(context.TODO(), dbDSN)
	if err != nil {
		panic(err)
	}
	return &Storage{
		conn: conn,
	}
}

func (s *Storage) Close() {
	s.conn.Close()
}
