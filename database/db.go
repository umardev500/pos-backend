package database

import (
	"context"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/pos-backend/constants"
	"github.com/umardev500/pos-backend/contracts"
)

type DB struct {
	conn *pgxpool.Pool
}

var (
	instance *DB
	once     sync.Once
)

func NewDB(ctx context.Context) *DB {
	once.Do(func() {
		DSN := os.Getenv("DSN")
		dbPool, err := pgxpool.New(ctx, DSN)
		if err != nil {
			log.Fatal().Msgf("Unable to connect to database: %v", err)
		}

		// Ping the database
		err = dbPool.Ping(ctx)
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

func (d *DB) WithTransaction(ctx context.Context, f func(context.Context) error) error {
	tx, err := d.conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	ctx = context.WithValue(ctx, constants.QueryKey, tx)
	err = f(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) GetConn(ctx context.Context) contracts.Query {
	return d.conn
}
