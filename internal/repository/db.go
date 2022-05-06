package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/artacone/go-url-short/internal/models"
)

type pg struct {
	pool *pgxpool.Pool
}

func newDB(dsn string) (*pg, error) {
	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}
	return &pg{pool: pool}, nil
}

// Create updates the url table with newly received values.
func (db *pg) Create(ctx context.Context, url models.URL, code models.ShortURL) (models.ShortURL, error) {
	const query = `insert into urls (url, short) VALUES ($1, $2) returning short`
	err := db.pool.QueryRow(ctx, query,
		url.URL,
		code.Code,
	).Scan(&code.Code)
	return code, err
}

// GetURL tries to get an url corresponding to the code stored in database.
// It returns an error if such code is not present.
func (db *pg) GetURL(ctx context.Context, code models.ShortURL) (models.URL, error) {
	var url models.URL
	const query = `select url from urls where short = $1`

	err := db.pool.QueryRow(ctx, query, code.Code).Scan(&url.URL)
	if errors.Is(err, pgx.ErrNoRows) {
		err = ErrNotFound
	}
	return url, err
}

// GetCode tries to get a code corresponding to the url stored in database.
// It returns an error if such url is not present.
func (db *pg) GetCode(ctx context.Context, url models.URL) (models.ShortURL, error) {
	var code models.ShortURL
	const query = `select short from urls where url = $1`

	err := db.pool.QueryRow(ctx, query, url.URL).Scan(&code.Code)
	if errors.Is(err, pgx.ErrNoRows) {
		err = ErrNotFound
	}
	return code, err
}
