package pg

import (
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"
)

//Config def
type Config struct {
	Host         string
	Port         int
	User         string
	Password     string
	DB           string
	MaxOpenConns int
	MaxIdleConns int
}

//NewPostgres create a sql conn
func NewPostgres(cfg *Config) *sql.DB {
	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
			cfg.Host,
			cfg.Port,
			cfg.User,
			cfg.DB,
			cfg.Password))
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	return db
}
