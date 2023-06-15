package golang_algorithms

func RepeatedStirngs(s string, n int64) int64 {

	lenghtString := int64(len(s))

	if lenghtString == 1 {
		return n
	}

	fullString := ContactString(s, n)
	var counts int64

	for _, v := range fullString {
		if string(v) == string([]rune(s)[0]) {
			counts++
		}
	}

	return counts
}

func ContactString(s string, n int64) string {
	var cs string = ""
	var lengBasicString = len(s)

	for i := 0; int64(i) < n; i += lengBasicString {
		cs += s
	}

	return cs
}
