package nodeutils

import "testing"

func TestList_Insert(t *testing.T) {
	link := List{}

	headValue := "Spinach"
	tailValue := "Nuts"

	link.Insert(tailValue)
	link.Insert("Vegan Ice Cream")
	link.Insert(headValue)

	if link.head.key != headValue {
		t.Fatalf("Expected head value to be %s, got %s\n", headValue, link.head.key)
	}

	if link.tail.key != tailValue {
		t.Fatalf("Expected head value to be %s, got %s\n", tailValue, link.tail.key)
	}

}
