package callbacklistener

import (
	"net/http"
	"os"

	"github.com/sad0vnikov/wundergram/bot"
	"github.com/sad0vnikov/wundergram/logger"
	"github.com/sad0vnikov/wundergram/wunderlist"
)

//Init starts httpserver
func Init() {
	port := os.Getenv("httpport")
	if len(port) == 0 {
		port = "8080"
	}

	http.HandleFunc("/c", func(w http.ResponseWriter, r *http.Request) {
		wunderlist.OnWundelistAuthCallback(w, r, bot.GetTelegramBotLink())
	})
	logger.Get("main").Infof("listening for Wundelist callback on :%v", port)

	http.ListenAndServe(":"+port, nil)
}
