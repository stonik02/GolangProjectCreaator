package cnst

const (
	App_go = `package main

	import "fmt"
	
func main() {
	fmt.Println("Hello World")
}
	`
	Docker_header = `version: '3.8'
services:

`
	Docker_pg = `  postgres:
    image: postgres:latest
    restart: always
    environment:
      POSTGRES_DB: EQueue
      POSTGRES_PASSWORD: qwerty123
      POSTGRES_USER: denis_postgresql
    ports:
      - 5432:5432

`
	Docker_redis = `  redis:
    image: redis:latest
    command: redis-server
    volumes:
      - redis:/var/lib/redis
      - redis-config:/usr/local/etc/redis/redis.conf
    ports:
      - 6379:6379
    networks:
      - redis-network
volumes:
  redis:
  redis-config:

networks:
  redis-network:
    driver: bridge
`
	Pg_client_file = `
package psql

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

)
	
type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}
	
func NewClient(ctx context.Context) (pool *pgxpool.Pool, err error) {
	maxAttempts := 5
	dsn := fmt.Sprintf("postgresql://user:password@host:port/db_name")
	err = DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
	
		pool, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			return err
		}

		return nil
	}, maxAttempts, 5*time.Second)
	
	if err != nil {
		log.Fatal("error do with tries postgresql")
	}

	return pool, nil
}
	
`

	Repeatable_file = `
package psql

import "time"

func DoWithTries(fn func() error, attemtps int, delay time.Duration) (err error) {
	for attemtps > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attemtps--

			continue
		}

		return nil
	}

	return
}

`
	Redis_client_file = `
package redis

import (
	"context"

	"github.com/go-redis/redis"

)
	
func NewClient(ctx context.Context) (rdb *redis.Client) {
	rdb = redis.NewClient(&redis.Options{
		Addr:       "localhost:6379",
		Password:   "", // no password set
		DB:         0,  // use default DB)
		MaxRetries: 5,
	})
	return rdb
}
	
`

	Config_Go = `
package config

import (
	"log/slog"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

//

type Config struct {
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		slog.Info("read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig("../../config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			slog.Info(help)
			slog.Error(err.Error())
			panic(err)
		}
	})
	return instance
}

`

	Cmd            = "cmd"
	Cmd_main       = "main"
	Internal       = "internal"
	Pkg            = "pkg"
	Docker_compose = "docker-compose.yml"
	App_golang     = "app.go"
	Redis          = "redisclient.go"
	Pg             = "pgclient.go"
	Config         = "config.go"
	Client         = "client"
	Go_redis       = "github.com/go-redis/redis"
	Pgx            = "github.com/jackc/pgx/v4"
	Pgx_pool       = "github.com/jackc/pgx/v4/pgxpool"
	CleanEnv       = "github.com/ilyakaznacheev/cleanenv"
)
