package postgres

import (
	"database/sql"
	"fmt"
	config "hotel-service/internal/pkg/load"

	_ "github.com/lib/pq"
)

func ConnectDB(cfg config.Config) (*sql.DB, error) {

	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", 
	cfg.Postgres.User, 
	cfg.Postgres.Password, 
	cfg.Postgres.Host, 
	cfg.Postgres.Port, 
	cfg.Postgres.Database)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}