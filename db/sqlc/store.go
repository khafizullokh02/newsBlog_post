package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	Querier
}

type SQLStore struct {
	DB *pgxpool.Pool
	*Queries
}

func NewStore(ctx context.Context, psqlURI string) Store {
	db, err := pgxpool.New(ctx, psqlURI)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

	log.Println("Connected to DB!")

	return &SQLStore{
		Queries: New(db),
		DB:      db,
	}
}
