package wunderlist

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/sad0vnikov/wundergram/logger"
	"github.com/sad0vnikov/wundergram/storage/tokens"
	"github.com/sad0vnikov/wundergram/util"
)

const userAuthLink = "https://www.wunderlist.com/oauth/authorize"

var clientID = os.Getenv("WUNDERLIST_CLIENT_ID")
var clientSecret = os.Getenv("WUNDERLIST_CLIENT_SECRET")

var wunderlistAuthRedirectURL = os.Getenv("WUNDERLIST_AUTH_REDIRECT_URL")
var stateString = util.GetRandomString(12)

//GetUserAuthLink returns a link for user for authorizing in Wunderlist
func GetUserAuthLink() string {

	return fmt.Sprintf(userAuthLink+"?client_id=%v&redirect_uri=%v&state=%v", clientID, url.QueryEscape(wunderlistAuthRedirectURL), stateString)
}

//OnWundelistAuthCallback processes Wunderlist auth callbacks
func OnWundelistAuthCallback(w http.ResponseWriter, r *http.Request, telegramBotLink string) {
	state := r.URL.Query().Get("state")
	if state != stateString {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	code := r.URL.Query().Get("code")
	http.Redirect(w, r, telegramBotLink+"?start="+code, http.StatusFound)
}

//AuthorizeUser saves Wunderlist user access token
func AuthorizeUser(userID int, authCode string) error {
	token, err := GetUserAccessToken(authCode)
	if err != nil {
		return err
	}
	if len(token) == 0 {
		return errors.New("didn't manage to get access token from Wunderlist")
	}

	logger.Get("main").Infof("saved token for user %v", userID)
	err = tokens.Put(userID, token)
	if err != nil {
		logger.Get("main").Errorf("error saving user token: %v", err)
	}

	return err
}

//IsUserAuthorized returns is user with given ID is authorized in Wunderlist
func IsUserAuthorized(userID int) (bool, error) {
	token, err := tokens.Get(userID)

	if err != nil {
		return false, err
	}

	return len(token) > 0, nil
}
