package algorithms

func RepeatedStirngs(s string, n int64) int64 {

	lenghtString := int64(len(s))

	if lenghtString == n {
		return n
	}

	fullString := contactString(s, int(n))

	
	return n
}

func contactString (s string, n int) string {

	if(len(s) <= n){
		s += s
		contactString(s, n)
	}

	if (len(s) % n != 0){
		
	}

	return s

}