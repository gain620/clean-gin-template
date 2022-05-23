// Package db provides singleton database client
package db

import (
	"clean-gin-template/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"net/url"
	"sync"
)

// singleton database connection
var client *sql.DB
var once sync.Once

func GetDSN(cfg *config.Config) string {
	dbHost := cfg.Database.Host
	dbPort := cfg.Database.Port
	dbUser := cfg.Database.User
	dbPass := cfg.Database.Pass
	dbName := cfg.Database.Name
	dbLocation := cfg.Database.Location

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", dbLocation)
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	fmt.Println(dsn)
	return dsn
}

// NewClient return singleton database client instance
func NewClient(cfg *config.Config) (*sql.DB, error) {
	once.Do(func() {
		dsn := GetDSN(cfg)
		dbConn, err := sql.Open(cfg.Database.Type, dsn)
		if err != nil {
			log.Fatal(err)
		}
		//defer client.Close()
		client = dbConn
	})
	return client, nil
}
