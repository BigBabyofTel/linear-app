package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/lucabrx/wuhu/api"
	"github.com/lucabrx/wuhu/config"
	_ "github.com/lucabrx/wuhu/docs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//	@title			Linear Clone API
//	@version		1.0
//	@description	API for managing linear clones
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		petstore.swagger.io
//	@BasePath	/v2

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		return
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger = log.With().Caller().Logger()

	dbInstance, err := accessDB(cfg.DbUrl)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect with database")
	}

	defer dbInstance.Close()
	log.Info().Msg("Connected to the database!")

	a := api.NewApplication(dbInstance, cfg)

	if err = a.Serve(); err != nil {
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
