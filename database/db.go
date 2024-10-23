package database

import (
	"context"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type DB struct {
	conn *pgxpool.Pool
}

var (
	instance *DB
	once     sync.Once
)

func NewDB() *DB {
	once.Do(func() {
		DSN := os.Getenv("DSN")
		dbPool, err := pgxpool.New(context.Background(), DSN)
		if err != nil {
			log.Fatal().Msgf("Unable to connect to database: %v", err)
		}

		// Ping the database
		err = dbPool.Ping(context.Background())
		if err != nil {
			log.Fatal().Msgf("Unable to ping database: %v", err)
		}

		instance = &DB{
			conn: dbPool,
		}

		log.Info().Msg("👋 Postgres connected successfully!")
	})

	return instance
}
