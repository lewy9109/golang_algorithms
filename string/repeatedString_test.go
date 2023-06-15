package string

import "testing"

func TestRepeatedString(t *testing.T) {
	text := "abcac"
	var n int64 = 10
	var want int64 = 4
	result := RepeatedStirngs(text, n)

	if want != result {
		t.Errorf("got %d, wanted %d", result, want)
	}
}

func TestRepeatedStringSecond(t *testing.T) {
	text := "a"
	var n int64 = 10
	var want int64 = 10
	result := RepeatedStirngs(text, n)

	if want != result {
		t.Errorf("got %d, wanted %d", result, want)
	}
}

func TestRepeatedStringThird(t *testing.T) {
        text := "aba"
        var n int64 = 10
        var want int64 = 7
        result := RepeatedStirngs(text, n)

        if want != result {
                t.Errorf("got %d, wanted %d", result, want)
        }
}


func TestAbString(t *testing.T) {
        text := "aba"
        var n int64 = 10
        var want string = "abaabaabaaba"
        result := ContactString(text, n)

        if want != result {
                t.Errorf("got %s, wanted %s", result, want)
        }
}

