package ut

import "testing"

func TestNewUT(t *testing.T) {
	u := NewUT()
	if u.Name == "GreyGrayPuff" {
		t.Logf("Name correct\n")
	} else {
		t.Fatalf("Expected \"GreyGrayPuff\""+
			" got %s\n", u.Name)
	}
}

func TestNewUT2(t *testing.T) {
	u := NewUT()
	u.Function("Cat")

	if value, ok := u.Mymap["Cat"]; ok {
		if value == "Mouse" {
			t.Logf("Success")
		} else {
			t.Logf("Got: %s, expected \"Mouse\"\n", value)
		}

	}

}
