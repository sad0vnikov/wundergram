package dialog

import (
	"testing"

	"gopkg.in/telegram-bot-api.v4"
)

func TestClearingMessageWordPuctuationSigns(t *testing.T) {
	word := "hello!.,/?:;%$@()=`~"
	expected := "hello"
	if cleaned := clearMessageWord(word); cleaned != expected {
		t.Errorf("error clearing word %v, got %v, expected %v", word, cleaned, expected)
	}
}

func TestClearingMessageWordPuctuationSignsUnicode(t *testing.T) {
	word := "привет!.,/?:;%$@()=`~"
	expected := "привет"
	if cleaned := clearMessageWord(word); cleaned != expected {
		t.Errorf("error clearing word %v, got %v, expected %v", word, cleaned, expected)
	}
}

func TestClearingMessageWordToLower(t *testing.T) {
	word := "HELLO"
	expected := "hello"
	if cleaned := clearMessageWord(word); cleaned != expected {
		t.Errorf("error clearing word %v, got %v, expected %v", word, cleaned, expected)
	}
}

func TestClearingMessageWordToLowerUnicode(t *testing.T) {
	word := "ПРИВЕТ"
	expected := "привет"
	if cleaned := clearMessageWord(word); cleaned != expected {
		t.Errorf("error clearing word %v, got %v, expected %v", word, cleaned, expected)
	}
}

func TestTraversingForward(t *testing.T) {

	treeRoot := NewConversationTreeNode(func(m *tgbotapi.Message, b *tgbotapi.BotAPI) {})

	childOne := NewConversationTreeNode(func(m *tgbotapi.Message, b *tgbotapi.BotAPI) {}).
		WithKeywords([]string{"spam"})

	childTwo := NewConversationTreeNode(func(m *tgbotapi.Message, b *tgbotapi.BotAPI) {}).
		WithKeywords([]string{"eggs"})

	treeRoot.AddChild(&childTwo)
	treeRoot.AddChild(&childOne)

	tree := NewConversationTree(&treeRoot)

	treeProcessor := NewTreeProcessor(&tree)

	message := makeMessage("green eggs")

	nodeToMoveIn := treeProcessor.GetNodeToMoveIn(&message, nil)

	if nodeToMoveIn == nil {
		t.Error("got nil node")

	}
	if nodeToMoveIn != &childTwo {
		t.Errorf("got wrong tree child, expected %v, got %v", &childTwo, nodeToMoveIn)
	}
}

func TestTraversingBackWithoutParent(t *testing.T) {
	treeRoot := NewConversationTreeNode(func(m *tgbotapi.Message, b *tgbotapi.BotAPI) {}).
		WithGoBackKeywords([]string{"hello"})

	tree := NewConversationTree(&treeRoot)
	treeProcessor := NewTreeProcessor(&tree)
	message := makeMessage("hello, world!")
	nodeToMoveIn := treeProcessor.GetNodeToMoveIn(&message, nil)

	if nodeToMoveIn != &treeRoot {
		t.Errorf("got wrong tree child, expected %v, got %v", &treeRoot, nodeToMoveIn)
	}
}

func makeMessage(text string) tgbotapi.Message {
	return tgbotapi.Message{Text: text, From: &tgbotapi.User{ID: 1}}
}
