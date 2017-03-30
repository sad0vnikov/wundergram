package dialog

import (
	"regexp"

	"strings"

	"gopkg.in/telegram-bot-api.v4"
)

//TreeProcessor calculates users traverses by the dialog tree
type TreeProcessor interface {
	GetNodeToMoveIn(msg *tgbotapi.Message, bot *tgbotapi.BotAPI) *TreeNode
}

//TreePositions stores users posions on the dialog tree
type TreePositions struct {
	tree          *Tree
	userPositions map[int]*TreeNode
}

//NewTreeProcessor initializes new tree processor
func NewTreeProcessor(tree *Tree) TreeProcessor {
	return &TreePositions{tree: tree, userPositions: map[int]*TreeNode{}}
}

//GetNodeToMoveIn takes user message and traverses user by the dialog tree
func (treeProcessor *TreePositions) GetNodeToMoveIn(msg *tgbotapi.Message, bot *tgbotapi.BotAPI) *TreeNode {
	msgText := msg.Text
	messageWords := strings.Split(msgText, " ")

	curNode := treeProcessor.findCurrentNodeForUser(msg.From.ID)
	var nodeToMoveIn *TreeNode
	for _, w := range messageWords {

		w = clearMessageWord(w)
		if curNode.parent != nil && curNode.parent.keywords[w] == true {
			return curNode.parent
		}

		for _, child := range curNode.children {
			if child.keywords[w] == true {
				nodeToMoveIn = child
				break
			}
		}
	}

	if nodeToMoveIn == nil {
		return curNode
	}

	return nodeToMoveIn
}

func clearMessageWord(word string) string {
	r := regexp.MustCompile(`\P{L}`)

	word = string(r.ReplaceAllString(word, ""))
	return strings.ToLower(word)
}

func (treeProcessor *TreePositions) findCurrentNodeForUser(userID int) *TreeNode {
	curNode := treeProcessor.userPositions[userID]
	if curNode == nil {
		curNode = treeProcessor.tree.root
		treeProcessor.updateCurrentNodeForUser(userID, curNode)
	}

	return curNode
}

func (treeProcessor *TreePositions) updateCurrentNodeForUser(userID int, node *TreeNode) {
	treeProcessor.userPositions[userID] = node
}
