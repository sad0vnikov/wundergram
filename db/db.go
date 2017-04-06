package db

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/sad0vnikov/wundergram/logger"
)

var db *bolt.DB

//Init initializes db connection
func Init() (*bolt.DB, error) {
	const dbName = "wundergram.db"
	conn, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		panic(fmt.Sprintf("cannot connect to db %v: %v", dbName, err))
	}
	logger.Get("main").Infof("connected to db %v", dbName)
	db = conn
	return conn, err
}

//GetDB returs BoltDB connection
func GetDB() *bolt.DB {
	if db == nil {
		Init()
	}
	return db
}
