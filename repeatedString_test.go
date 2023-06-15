package golang_algorithms

import "testing"

func TestContactString(t *testing.T) {
	text := "abcac"
	var n int64 = 10
	var want int64 = 4
	result := RepeatedStirngs(text, n)

	if want != result {
		t.Errorf("got %q, wanted %q", result, want)
	}
}

func TestContactStringSecond(t *testing.T) {
	text := "a"
	var n int64 = 10
	var want int64 = 10
	result := RepeatedStirngs(text, n)

	if want != result {
		t.Errorf("got %q, wanted %q", result, want)
	}
}
