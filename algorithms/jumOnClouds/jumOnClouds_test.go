package jumOnClouds

import "testing"

func TestJumpingOnClouds(t *testing.T) {
	var array [6]int32 = [6]int32{0,0,1,0,1,0}

	var want int32 = 3
	result := JumpingOnClouds(array[:])

	if want != result {
		t.Errorf("got %q, wanted %q", result, want)
	}
}