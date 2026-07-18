package database

import "github.com/jackc/pgx/v5/pgxpool"

// DB is the global database connection pool.
var DB *pgxpool.Pool
