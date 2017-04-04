package daily_notifications

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/sad0vnikov/wundergram/db"
)

//DailyNotificationConfig is a number of preferences for user's daily tasks notificatons
type DailyNotificationConfig struct {
	UserID                int
	NotificationTimestamp int //Notification time is stored as as number of seconds from the beginning of the day
}

var dailyNotificationsBucketName = []byte("daily_notifications")

//EnableNotificationsForUser saves user daily notification time
//notificationTime param should be a time formatted as HH:MM
func EnableNotificationsForUser(userID int, notificationTime string) error {
	dayOffset, err := stringTimeToDayOffset(notificationTime)
	if err != nil {
		return err
	}

	return Save(DailyNotificationConfig{UserID: userID, NotificationTimestamp: dayOffset})
}

func stringTimeToDayOffset(notificationTime string) (int, error) {
	marshalls := strings.Split(notificationTime, ":")
	if len(marshalls) != 2 {
		return 0, errors.New("Invalid time format, expected HH:MM format")
	}

	hours, hoursErr := strconv.ParseInt(marshalls[0], 10, 0)
	minutes, minErr := strconv.ParseInt(marshalls[1], 10, 0)

	if hoursErr != nil || minErr != nil {
		return 0, errors.New("Invalid time value")
	}

	return int(hours*60*60 + minutes*60), nil
}

//Save saves daily notification config
func Save(config DailyNotificationConfig) error {
	configJSON, err := json.Marshal(config)

	if err != nil {
		return err
	}

	return db.GetDB().Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(dailyNotificationsBucketName)
		if err != nil {
			return err
		}
		return b.Put([]byte(strconv.Itoa(config.UserID)), configJSON)
	})
}

//GetAll gets all the stored daily notification configs
func GetAll() ([]DailyNotificationConfig, error) {
	var result []DailyNotificationConfig
	err := db.GetDB().View(func(tx *bolt.Tx) error {
		b := tx.Bucket(dailyNotificationsBucketName)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var nextConfig DailyNotificationConfig
			err := json.Unmarshal(v, nextConfig)
			if err != nil {
				return err
			}
			result = append(result, nextConfig)

		}

		return nil
	})

	return result, err
}

//GetByUserID returns a daily notifications config for user
func GetByUserID(userID int) (DailyNotificationConfig, error) {
	var result DailyNotificationConfig
	err := db.GetDB().View(func(tx *bolt.Tx) error {
		resJSON := tx.Bucket(dailyNotificationsBucketName).
			Get([]byte(strconv.Itoa(userID)))

		return json.Unmarshal(resJSON, result)

	})
	return result, err
}