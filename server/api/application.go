package api

import (
	"database/sql"
	"os"
	"sync"

	"github.com/lucabrx/wuhu/internal/aws"
	"github.com/sashabaranov/go-openai"

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
	config config.AppConfig
	AWS    aws.AWS
	AI     *openai.Client
}

//	@title			Linear-Clone API
//	@version		1.0
//	@description	API for Linear-Clone
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		petstore.swagger.io
//	@BasePath	/v2

func NewApplication(db *sql.DB, cfg config.AppConfig, ai *openai.Client) *app {
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	return &app{
		logger: &logger,
		DB:     data.NewModals(db),
		config: cfg,
		AWS:    aws.NewAws(cfg.AwsAccessKeyId, cfg.AwsSecretAccessKey),
		AI:     ai,
	}
}
