package repeatedString

func RepeatedStirngs(s string, n int64) int64 {

	lenghtString := int64(len(s))

	if lenghtString == 1 {
		return n
	}

	// 1. Mamy nieskonczony ciag znakow np abca

	// 2  n jest liczba gdze trzeba znaleźć ile razy sie powtaza pierwsza litera nieskonczonego stringu

	//znalezc pierwsza litere stringu
	firstLetter := s[0]
	fullString := ContactString(s, n)
	var counts int64
	//ile razy powtarza sie pierwsza litera§
	for _, v := range fullString {
		if string(v) == string(firstLetter) {
			counts++
		}
	}

	//fullString := ContactString(s, int(n))

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
