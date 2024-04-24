package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"sync"

	"github.com/lucabrx/wuhu/internal/aws"
	"github.com/lucabrx/wuhu/internal/cache"
	"github.com/redis/go-redis/v9"

	"github.com/lucabrx/wuhu/config"
	"github.com/lucabrx/wuhu/internal/data"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type app struct {
	logger *zerolog.Logger
	wg     sync.WaitGroup
	DB     data.Models
	Cache  cache.RedisCache
	config config.AppConfig
	AWS    aws.AWS
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		return
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger = log.With().Caller().Logger()

	dbInstance, err := accessDB("postgresql://postgres:postgres@localhost:5431/postgres?sslmode=disable")

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect with database")
	}

	defer dbInstance.Close()
	log.Info().Msg("Connected to the database!")

	redisClient, err := accessRedis("localhost:6379", "", 0)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to Redis")
	}

	defer redisClient.Close()
	log.Info().Msg("Connected to Redis!")

	a := app{
		logger: &log.Logger,
		DB:     data.NewModals(dbInstance),
		config: cfg,
		AWS:    aws.NewAws(cfg.AwsAccessKeyId, cfg.AwsSecretAccessKey),
		Cache:  *cache.NewRedisCache(redisClient),
	}

	if err = a.serve(); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}

func accessDB(dbUrl string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dbUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

func accessRedis(addr, password string, db int) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // default DB
	})

	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	return rdb, nil
}
