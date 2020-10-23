package main

import "testing"

func TestMakeMsg(t *testing.T) {
	actual := MakeMsg("Casey", 3)
	expected := "Hello Casey\nHello Casey\nHello Casey"

	if actual != expected {
		t.Errorf("MakeMsg(Casey, 3) = %s; want %s", actual, expected)
	}
}
