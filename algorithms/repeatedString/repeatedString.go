package repeatedString

func RepeatedStirngs(s string, n int64) int64 {

	lenghtString := int64(len(s))

	if lenghtString == n {
		return n
	}

	//fullString := ContactString(s, int(n))

	
	return n
}

func ContactString (s string, n int) string {
	var cs string = ""
	var lengBasicString = len(s)

	for i := 0; i < n; i += lengBasicString {
		cs += s
	}
	rest := len(cs)
	
	if (rest % n != 0){
		
	}


	return cs
}