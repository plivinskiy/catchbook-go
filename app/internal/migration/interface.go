package migration

import (
	"catchbook/internal/config"
	"catchbook/pkg/db"
	"context"
)

type Migration interface {
	Migrate()
}

func Migrations() []Migration {
	ctx := context.Background()
	cfg := config.CreateConfig()
	conn := db.NewMysqlClient(ctx, cfg.DatabaseDsn)
	var migrations []Migration
	migrations = append(migrations, V01{Db: conn})
	return migrations
}
