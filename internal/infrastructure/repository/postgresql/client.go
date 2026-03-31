package postgresql

import (
	"context"
	"dudanseai_bot/configs"
	"dudanseai_bot/pkg/utils"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Client struct {
	pool *pgxpool.Pool
}

func (c *Client) Ping(ctx context.Context) error {
	return c.pool.Ping(ctx)
}

func (c *Client) Exec(ctx context.Context, query string, args ...any) (commandTag pgconn.CommandTag, err error) {
	commandTag, err = c.pool.Exec(ctx, query, args...)
	if err != nil {
		if pgErr := c.handlePgError(err); pgErr != nil {
			return commandTag, pgErr
		}
	}

	return commandTag, nil
}

func (c *Client) Query(ctx context.Context, query string, args ...any) (rows pgx.Rows, err error) {
	rows, err = c.pool.Query(ctx, query, args...)
	if err != nil {
		if pgErr := c.handlePgError(err); pgErr != nil {
			return nil, pgErr
		}
	}

	return rows, err
}

func (c *Client) QueryRow(ctx context.Context, query string, args ...any) pgx.Row {
	return c.pool.QueryRow(ctx, query, args...)
}

func (c *Client) Close() {
	if c.pool != nil {
		c.pool.Close()
	}
}

func NewClient(ctx context.Context, maxAttempts int, cfg configs.DbConfig) (client Client, err error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	err = repeatable.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		client.pool, err = pgxpool.New(ctx, dsn)
		if err != nil {
			return err
		}

		return nil
	}, maxAttempts, 5*time.Second)

	if err != nil {
		return Client{}, err
	}

	return client, err
}

func (c *Client) handlePgError(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return fmt.Errorf("SQL error: message: %s, detail: %s, where: %s, code: %s",
			pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code)
	}
	return nil
}
