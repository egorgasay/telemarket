package client

import (
	"testing"
)

func TestNewKeyboardButtonURL(t *testing.T) {
	b := NewKeyboardButtonURL("text", "url")
	if b.Text != "text" {
		t.Errorf("NewKeyboardButtonURL() = %v, want %v", b.Text, "text")
	}
	if *b.URL != "url" {
		t.Errorf("NewKeyboardButtonURL() = %v, want %v", *b.URL, "url")
	}

	b = NewKeyboardButtonURL("text2", "url2")
	if b.Text != "text2" {
		t.Errorf("NewKeyboardButtonURL() = %v, want %v", b.Text, "text2")
	}
	if *b.URL != "url2" {
		t.Errorf("NewKeyboardButtonURL() = %v, want %v", *b.URL, "url2")
	}
}

func TestNewKeyboardButtonWithData(t *testing.T) {
	b := NewKeyboardButtonWithData("text", "data")
	if b.Text != "text" {
		t.Errorf("NewKeyboardButtonWithData() = %v, want %v", b.Text, "text")
	}
	if *b.CallbackData != "data" {
		t.Errorf("NewKeyboardButtonWithData() = %v, want %v", *b.CallbackData, "data")
	}

	b = NewKeyboardButtonWithData("text2", "data2")
	if b.Text != "text2" {
		t.Errorf("NewKeyboardButtonWithData() = %v, want %v", b.Text, "text2")
	}
	if *b.CallbackData != "data2" {
		t.Errorf("NewKeyboardButtonWithData() = %v, want %v", *b.CallbackData, "data2")
	}
}
