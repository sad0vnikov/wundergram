package commands

import (
	"github.com/sad0vnikov/wundergram/bot/dialog"
)

//BuildConversationTree returns a bot conversation tree
func BuildConversationTree() dialog.Tree {
	dialogRoot := dialog.NewConversationTreeNode(start)

	showDailyNotificationsTimeSelector := dialog.NewConversationTreeNode(showDailyNotificationsTimeSelector).
		WithKeywords([]string{"send"}).
		WithGoBackKeywords([]string{"forget", "cancel"})

	enableDailyNotifications := dialog.NewConversationTreeNode(enableDailyNotifications).
		WithRegexp(`\d{2}:\d{2}`)

	showDailyNotificationsTimeSelector.AddChild(&enableDailyNotifications)

	dialogRoot.AddChild(&showDailyNotificationsTimeSelector)

	tree := dialog.NewConversationTree(&dialogRoot)

	return tree
}
