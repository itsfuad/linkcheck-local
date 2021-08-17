package main

import (
	"encoding/json"
	"testing"
)

func TestJSONResult(t *testing.T) {
	_, err := json.Marshal(Analyze(""))
	if err != nil {
		t.Fatal(err)
	}
}
