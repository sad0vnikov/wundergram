package dialog

import (
	"regexp"

	"strings"
	"sync"

	"gopkg.in/telegram-bot-api.v4"
)

//Processor calculates users traverses by the dialog tree
type Processor interface {
	GetNodeToMoveIn(msg *tgbotapi.Message, bot *tgbotapi.BotAPI) *TreeNode
	RunNodeHandler(node *TreeNode, msg *tgbotapi.Message, bot *tgbotapi.BotAPI)
}

//TreeProcessor stores users posions on the dialog tree
type TreeProcessor struct {
	tree          *Tree
	userPositions map[int]*TreeNode
}

//NewProcessor initializes new tree processor
func NewProcessor(tree *Tree) Processor {
	return &TreeProcessor{tree: tree, userPositions: map[int]*TreeNode{}}
}

//GetNodeToMoveIn takes user message and traverses user by the dialog tree
func (processor *TreeProcessor) GetNodeToMoveIn(msg *tgbotapi.Message, bot *tgbotapi.BotAPI) *TreeNode {
	msgText := msg.Text
	messageWords := strings.Split(msgText, " ")

	curNode := processor.findCurrentNodeForUser(msg.From.ID)
	if len(curNode.Children) == 0 {
		curNode = processor.tree.Root
	}

	var nodeToMoveIn *TreeNode
	nodeToMoveIn = findNodeByKeywords(curNode, processor.tree.Root, messageWords)
	if nodeToMoveIn == nil {
		nodeToMoveIn = findNodeByRegex(curNode, processor.tree.Root, msgText)
	}

	if nodeToMoveIn == nil {
		return curNode
	}

	return nodeToMoveIn
}

//RunNodeHandler moves user to node and runs a handler
func (processor *TreeProcessor) RunNodeHandler(node *TreeNode, msg *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	node.Handler(msg, bot)
	processor.updateCurrentNodeForUser(msg.From.ID, node)
}

func findNodeByKeywords(curNode *TreeNode, treeRooteNode *TreeNode, messageWords []string) *TreeNode {
	for _, w := range messageWords {

		w = clearMessageWord(w)

		if checkNodeHasKeyword(treeRooteNode, w) {
			return treeRooteNode
		}

		if curNode.Parent != nil && checkNodeHasKeyword(curNode.Parent, w) {
			return curNode.Parent
		}

		for _, child := range curNode.Children {
			if checkNodeHasKeyword(child, w) {
				return child
			}
		}
	}
	return nil
}

func checkNodeHasKeyword(node *TreeNode, keyword string) bool {
	return node.keywords[keyword] == true
}

func findNodeByRegex(curNode *TreeNode, treeRootNode *TreeNode, message string) *TreeNode {

	if checkNodeByRegex(treeRootNode, message) {
		return treeRootNode
	}

	if curNode.Parent != nil && checkNodeByRegex(curNode.Parent, message) {
		return curNode.Parent
	}

	for _, v := range curNode.Children {
		if checkNodeByRegex(v, message) {
			return v
		}
	}

	return nil
}

func checkNodeByRegex(node *TreeNode, message string) bool {
	if len(node.regexp) > 0 {
		r := regexp.MustCompile(node.regexp)
		if r.MatchString(message) {
			return true
		}
	}

	return false
}

func clearMessageWord(word string) string {
	r := regexp.MustCompile(`\P{L}`)

	word = string(r.ReplaceAllString(word, ""))
	return strings.ToLower(word)
}

func (processor *TreeProcessor) findCurrentNodeForUser(userID int) *TreeNode {
	curNode := processor.userPositions[userID]
	if curNode == nil {
		curNode = processor.tree.Root
		processor.updateCurrentNodeForUser(userID, curNode)
	}

	return curNode
}

func (processor *TreeProcessor) updateCurrentNodeForUser(userID int, node *TreeNode) {
	mutex := sync.Mutex{}
	mutex.Lock()
	processor.userPositions[userID] = node
	mutex.Unlock()
}
