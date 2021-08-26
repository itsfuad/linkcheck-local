package main

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestSemanticExample(t *testing.T) {
	data, err := ioutil.ReadFile("examples/input.txt")
	if err != nil {
		t.Fatal(err)
	}
	got := Analyze(string(data))["external"]
	want := 1
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected value: %#v want %#v", got, want)
	}
}
