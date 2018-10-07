package db

import (
	"database/sql"
	"fmt"

	// drivers
	_ "github.com/lib/pq"
)

// Config stores info required to connect the database
type Config struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
	Name string `json:"name"`
}

func (c *Config) String() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Pass, c.Name)
}

var db *sql.DB

// Init inits new connection pool
func Init(conf *Config) (err error) {
	db, err = sql.Open("postgres", conf.String())
	if err != nil {
		return
	}

	db.SetMaxOpenConns(2)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(0)

	return db.Ping()
}
