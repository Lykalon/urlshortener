package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgStorage struct {
	pool *pgxpool.Pool
}

func (p *PgStorage) Init() {
	pool, err := pgxpool.New(context.Background(), os.Getenv("PG_CONNECTION_STRING"))
	if err != nil {
		log.Fatal("Unable to create connection pool")
		log.Println("Try switch to local storage")
		os.Exit(1)
	}
	p.pool = pool
	p.pool.Exec(context.Background(), `create table if not exists "data" (
									id serial primary key,
									short_url bigint not null,
									full_url text not null);`)
	p.pool.Exec(context.Background(), "create index if not exists idx_short_url on data (short_url);")
}

func (p *PgStorage) Save(shortLink int64, fullLink string) {
	tx, err := p.pool.Begin(context.Background())
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), "insert into data (short_url, full_url) values ($1, $2)", shortLink, fullLink)
	if err != nil {
		log.Fatalln(err)
		return
	}
	err = tx.Commit(context.Background())
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func (p *PgStorage) FindFull(shortLink int64) (string, bool) {
	var fullLink string
	err := p.pool.QueryRow(context.Background(), "select full_url from data where short_url=$1", shortLink).Scan(&fullLink)
	if err == pgx.ErrNoRows {
		return "", false
	}
	return fullLink, true
}

func (p *PgStorage) FindShort(fullLink string) (int64, bool) {
	var shortLink int64
	err := p.pool.QueryRow(context.Background(), "select short_url from data where full_url=$1", fullLink).Scan(&shortLink)
	if err == pgx.ErrNoRows {
		return 0, false
	}
	return shortLink, true
}

func (p *PgStorage) Close() {
	p.pool.Close()
}
