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

func TestJumpOnCloudsSecondTest(t *testing.T) {
	var array [7]int32 = [7]int32{0,0,1,0,0,1,0}

	var want int32 = 4
	result := JumpingOnClouds(array[:])

	if want != result {
		t.Errorf("got %q, wanted %q", result, want)
	}
}

func TestJumpOnCloudOnlyZero(t *testing.T) {
	var array [7]int32 = [7]int32{0,0,0,0,0,0,0}

	var want int32 = 3
	result := JumpingOnClouds(array[:])

	if want != result {
		t.Errorf("got %q, wanted %q", result, want)
	}
}