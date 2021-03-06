package tokens

import (
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/sad0vnikov/wundergram/db"
	"github.com/sad0vnikov/wundergram/logger"
)

var bucketName = []byte("tokens")

//Put saves Wunderlist token for given Telegram user ID
func Put(userID int, token string) error {
	return db.GetDB().Update(func(tx *bolt.Tx) error {

		b, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			return err
		}

		err = b.Put([]byte(strconv.Itoa(userID)), []byte(token))
		if err != nil {
			logger.Get("main").Errorf("erro saving token: %v", err)
		}

		return nil
	})
}

//Get returns Wunderlist token for given Telegram user ID
func Get(userID int) (string, error) {
	var token string
	error := db.GetDB().View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		if b == nil {
			return nil
		}
		v := b.Get([]byte(strconv.Itoa(userID)))
		token = string(v)
		return nil
	})

	return token, error
}
