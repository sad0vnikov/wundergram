package commands

import "github.com/sad0vnikov/wundergram/bot/dialog"

//BuildConversationTree returns a bot conversation tree
func BuildConversationTree() dialog.Tree {
	dialogRoot := dialog.NewConversationTreeNode(start)

	enablingNotificationsNode := dialog.NewConversationTreeNode(enableDailyNotifications).
		WithKeywords([]string{"send"}).
		WithGoBackKeywords([]string{"forget", "cancel"})

	dialogRoot.AddChild(&enablingNotificationsNode)

	tree := dialog.NewConversationTree(&dialogRoot)

	return tree
}
