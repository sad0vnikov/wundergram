package daily_notifications

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/boltdb/bolt"
	"github.com/sad0vnikov/wundergram/db"
	"github.com/sad0vnikov/wundergram/util"
)

//DailyNotificationConfig is a number of preferences for user's daily tasks notificatons
type DailyNotificationConfig struct {
	UserID                int
	ChatID                int64
	NotificationTimestamp int //Notification time is stored as as number of seconds from the beginning of the day
	LastTimeActivated     int64
}

var dailyNotificationsBucketName = []byte("daily_notifications")

//CheckIsTimeToSend returns true if a time for sending notification already came
func (notification DailyNotificationConfig) CheckIsTimeToSend(currentTime time.Time) bool {
	currentTimestamp := currentTime.Unix()
	dayStart := util.GetDayStart(currentTime).Unix()
	currentDayOffset := currentTimestamp - dayStart
	return currentDayOffset >= int64(notification.NotificationTimestamp)
}

//CheckWasSentToday returns true if notification was already sent today
func (notification DailyNotificationConfig) CheckWasSentToday(currentTime time.Time) bool {
	dayStart := util.GetDayStart(currentTime).Unix()

	return notification.LastTimeActivated >= dayStart
}

//EnableNotificationsForUser saves user daily notification time
//notificationTime param should be a time formatted as HH:MM
func EnableNotificationsForUser(userID int, chatID int64, notificationTime string) error {
	dayOffset, err := stringTimeToDayOffset(notificationTime)
	if err != nil {
		return err
	}

	return Save(DailyNotificationConfig{UserID: userID, ChatID: chatID, NotificationTimestamp: dayOffset})
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
	result := []DailyNotificationConfig{}
	err := db.GetDB().View(func(tx *bolt.Tx) error {
		b := tx.Bucket(dailyNotificationsBucketName)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var nextConfig DailyNotificationConfig
			err := json.Unmarshal(v, &nextConfig)
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
	result := DailyNotificationConfig{}
	err := db.GetDB().View(func(tx *bolt.Tx) error {
		resJSON := tx.Bucket(dailyNotificationsBucketName).
			Get([]byte(strconv.Itoa(userID)))

		return json.Unmarshal(resJSON, &result)

	})
	return result, err
}
