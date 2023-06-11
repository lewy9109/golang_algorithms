package repeatedString

import "testing"


func TestContactString(t *testing.T) {
	text := "abc"
	want := "abcabc"

	result := ContactString(text, 6)
	
	if want != result {
		t.Errorf("got %q, wanted %q", result, want)
	}
}

func TestRepetedString(t *testing.T) {
	text := "aba"
	var n int64 = 10

	var want int64 = 7

	result := RepeatedStirngs(text, n)

	if want != result {
		t.Errorf("got %q, wanted %q", result, want)
	}
}

func TestTestRepetedStringSecond(t *testing.T) {
	text := "abcac"
	var n int64 = 10

	var want int64 = 4

	result := RepeatedStirngs(text, n)

	if want != result {
		t.Errorf("got %q, wanted %q", result, want)
	}
}