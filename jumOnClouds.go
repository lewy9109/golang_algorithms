package golang_algorithms

func JumpingOnClouds(c []int32) int32 {
	var jump int32 = 0

	for i := 0; i < len(c); i++ {
		if i+2 < len(c) && c[i+2] == 0 {
			jump++
			i++
			continue
		}
		if i+2 < len(c) && c[i+2] == 1 {
			jump++
			continue
		}
		if i+1 < len(c) && c[i+1] == 0 {
			jump++
		}
	}

	return jump
}
