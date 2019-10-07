package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	result := Greeting(10)
	if result != "Good Morning" {
		t.Error("not correct greeting")
	}
}
