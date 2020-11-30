package util

import (
	"testing"
)

func TestNewUT(t *testing.T) {
	ut := NewUT()
	if ut.Name != "GreyGrayPuff"{
		t.Fatalf("Expecting something else!")
	}
}

func TestList_Insert(t *testing.T) {
	link := List{}
	link.Insert("Rolls")
}