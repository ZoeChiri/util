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
