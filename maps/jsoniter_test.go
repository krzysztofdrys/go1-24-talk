package maps

import (
	jsoniter "github.com/json-iterator/go"
	"testing"
)

type Metadata struct {
	Type string `json:"type"`
}

func TestJsonIter(t *testing.T) {
	m := map[string][]Metadata{"item1": []Metadata{{Type: "gauge"}}, "item2": []Metadata{{Type: "summary"}}}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	s, err := json.Marshal(m)
	if err != nil {
		t.FailNow()
	}
	if len(s) < 40 {
		t.Errorf("Result too short (%d): %q", len(s), s)
	}
}
