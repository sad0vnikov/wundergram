package dialog

import (
	"gopkg.in/telegram-bot-api.v4"
)

//Tree represents conversation tree
type Tree struct {
	Root *TreeNode
}

//TreeNode represents bot conversation tree node
type TreeNode struct {
	keywords       map[string]bool //Keywords for moving to this node
	goBackKeywords map[string]bool
	regexp         string
	Handler        func(*tgbotapi.Message, *tgbotapi.BotAPI)
	Parent         *TreeNode
	Children       []*TreeNode
}

//NewConversationTree initializes a new conversation tree
func NewConversationTree(rootNode *TreeNode) Tree {
	return Tree{Root: rootNode}
}

//NewConversationTreeNode creates a new tree node
func NewConversationTreeNode(handlerFunc func(*tgbotapi.Message, *tgbotapi.BotAPI)) TreeNode {
	return TreeNode{
		Handler:        handlerFunc,
		Children:       make([]*TreeNode, 0),
		goBackKeywords: map[string]bool{},
		keywords:       map[string]bool{},
	}
}

//WithKeywords sets keywords required for moving into the node
func (node TreeNode) WithKeywords(keywords []string) TreeNode {
	node.keywords = makeKeywordsMap(keywords)

	return node
}

//WithGoBackKeywords sets keywords required to move to parent node
func (node TreeNode) WithGoBackKeywords(keywords []string) TreeNode {
	node.goBackKeywords = makeKeywordsMap(keywords)

	return node
}

//WithRegexp sets a regexp the message needs to match for traversing into the node
func (node TreeNode) WithRegexp(regexp string) TreeNode {
	node.regexp = regexp

	return node
}

func makeKeywordsMap(keywords []string) map[string]bool {
	keywordsMap := map[string]bool{}
	for _, s := range keywords {
		keywordsMap[s] = true
	}
	return keywordsMap
}

//AddChild adds a child to the node
func (node *TreeNode) AddChild(childNode *TreeNode) {
	node.Children = append(node.Children, childNode)
}
