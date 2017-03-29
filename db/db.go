package db

import (
	"log"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

//Init initializes db connection
func Init() (*bolt.DB, error) {
	const dbName = "wundergram.db"
	conn, err := bolt.Open(dbName, 0600, nil)
	log.Printf("connected to db %v", dbName)
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
