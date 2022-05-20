// Package db provides singleton database client
package db

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/url"
	"sync"
)

func init() {
	viper.SetConfigFile(`./configs/db-configs.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

// singleton database connection
var client *sql.DB
var once sync.Once

//func getOption() option.ClientOption {
//	rawString := os.Getenv("SERVICE_ACCOUNT_KEY")
//	return option.WithCredentialsJSON([]byte(rawString))
//}

func getDBDSN() string {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	dbLocation := viper.GetString(`database.location`)

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", dbLocation)
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	return dsn
}

// GetClient return singleton database client instance
func GetClient() (*sql.DB, error) {
	once.Do(func() {
		dsn := getDBDSN()
		dbConn, err := sql.Open(`mysql`, dsn)
		if err != nil {
			panic(err)
		}
		//defer client.Close()
		client = dbConn
	})
	return client, nil
}
