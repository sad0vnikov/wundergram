package tokens

import (
	"github.com/boltdb/bolt"
	"github.com/sad0vnikov/wundergram/db"
)

var bucketName = []byte("tokens")

//Put saves Wunderlist token for given Telegram user ID
func Put(userID int, token string) error {
	return db.GetDB().Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			return err
		}

		err = b.Put([]byte(string(userID)), []byte(token))
		return err
	})
}

//Get returns Wunderlist token for given Telegram user ID
func Get(userID int) (string, error) {
	var token string
	error := db.GetDB().View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		v := b.Get([]byte(string(userID)))
		token = string(v)
		return nil
	})

	return token, error
}
