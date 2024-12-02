package db

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
	"github.com/rs/zerolog/log"

	"github.com/effective_mobile_task/config"
)

//go:embed migrations/*.sql
var embedSQLMigrations embed.FS

func InitPostgresDB(cfg *config.Config) (*pg.DB, error) {
	var (
		opts = &pg.Options{
			Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			User:     cfg.User,
			Password: cfg.Password,
			Database: cfg.DBName,
		}
		err error
	)
	db := pg.Connect(opts)

	collection := migrations.NewCollection()
	err = collection.DiscoverSQLMigrationsFromFilesystem(http.FS(embedSQLMigrations), "migrations")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to scan migrations files")
		return nil, err
	}

	_, _, err = collection.Run(db, "init")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed migrations init")
		return nil, err
	}

	oldVersion, newVersion, err := collection.Run(db, "up")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed migrations up")
		return nil, err
	}
	if newVersion != oldVersion {
		log.Info().Msg(fmt.Sprintf("migrated from version %d to %d\n", oldVersion, newVersion))
	} else {
		log.Info().Msg(fmt.Sprintf("version is %d\n", oldVersion))
	}
	return db, nil
}
