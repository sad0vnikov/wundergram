package tasks

import (
	"time"
)

//RunDailyNotifications runs daily notifications loop
func RunDailyNotifications() {
	ticker := time.NewTicker(time.Minute)

	for range ticker.C {
		checkDailyNotifications()
	}

}

func checkDailyNotifications() {
	// notifications, err := daily_notifications.GetAll()
	// if err != nil {
	// 	log.Panic("error getting notifiations list from db")
	// 	return
	// }
	// for _, n := range notifications {

	// }
}
