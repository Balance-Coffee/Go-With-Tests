package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(1, 1)
	expected := 2
	if sum != expected {
		t.Errorf("Expected '%d' got '%d'", expected, sum)
	}
	// assertCorrectMessage(t, sum, expected)
}

// func assertCorrectMessage(t *testing.TB, sum, expected int) {
// 	// t.Helper()
// 	if sum != expected {
// 		t.Errorf("Expected '%d' got '%d'", expected, sum)
// 	}
// }
