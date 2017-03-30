package dialog

import (
	"reflect"
	"testing"
)

func TestMakingKeywordsMap(t *testing.T) {
	keywords := []string{"a", "b", "c"}
	expectedMap := map[string]bool{
		"a": true,
		"b": true,
		"c": true,
	}

	calculatedMap := makeKeywordsMap(keywords)
	if !reflect.DeepEqual(calculatedMap, expectedMap) {
		t.Errorf("error making map from keywods %v: expected %v, got %v", keywords, expectedMap, calculatedMap)
	}
}
