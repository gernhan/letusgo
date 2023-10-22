package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var pool *pgxpool.Pool

func GetPool() *pgxpool.Pool {
	return pool
}

// InitDB initializes the database connection pool.
func InitDB(connString string) error {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return err
	}

	// Customize connection pool configuration.
	config.MaxConns = 20                       // Set maximum number of connections in the pool.
	config.MaxConnIdleTime = 5 * time.Minute   // Set the maximum idle time for a connection.
	config.HealthCheckPeriod = 1 * time.Minute // Set the interval between health checks.
	// Add more specialized configurations as needed.

	pool, err = pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return err
	}

	return nil
}

// CloseDB closes the database connection pool.
func CloseDB() {
	pool.Close()
}

