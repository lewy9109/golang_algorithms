package repeatedString

func RepeatedStirngs(s string, n int64) int64 {

	lenghtString := int64(len(s))

	if lenghtString == n {
		return n
	}

	fullString := ContactString(s, int(n))

	
	return n
}

func ContactString (s string, n int) string {

	if(len(s) <= n){
		s += s
		ContactString(s, n)
	}

	if (len(s) % n != 0){
		
	}
	return s
}