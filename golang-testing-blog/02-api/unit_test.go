package main

import "testing"

func TestMakeItProud(t *testing.T) {
	text := "sweet"
	expected := "SWEET!!!"
	makeItProud(&text)
	if text != expected {
		t.Errorf("Fail")
	}
}
