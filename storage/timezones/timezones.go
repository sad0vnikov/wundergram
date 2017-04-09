package timezones

import (
	"strconv"
	"time"

	"github.com/boltdb/bolt"
	"github.com/sad0vnikov/wundergram/db"
	"github.com/sad0vnikov/wundergram/logger"
)

var bucketName = []byte("user_timezones")

//Get returns a stored user location
func Get(userID int) (*time.Location, error) {
	var location *time.Location
	locationIsStored := false
	err := db.GetDB().View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		if b == nil {
			return nil
		}

		tz := b.Get([]byte(strconv.Itoa(userID)))
		loc, _ := time.LoadLocation(string(tz))
		location = loc
		locationIsStored = true
		return nil
	})

	if err != nil {
		logger.Get("main").Errorf("error getting user location: $%v", err)
		return location, err
	}

	return location, nil
}

//UserHasLocation returns true if user has stored location
func UserHasLocation(userID int) bool {
	userHasLocation := false
	err := db.GetDB().View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		if b == nil {
			return nil
		}
		tz := b.Get([]byte(strconv.Itoa(userID)))
		if len(tz) > 0 {
			userHasLocation = true
		}
		return nil
	})

	if err != nil {
		logger.Get("main").Errorf("error checking if user location is stored: $%v", err)
	}

	return userHasLocation
}

//Put saves user timezone
func Put(userID int, timezone *time.Location) error {
	err := db.GetDB().Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			return err
		}

		b.Put([]byte(strconv.Itoa(userID)), []byte(timezone.String()))

		return nil
	})

	if err != nil {
		logger.Get("main").Errorf("error saving user location: $%v", err)
	}

	return err
}
