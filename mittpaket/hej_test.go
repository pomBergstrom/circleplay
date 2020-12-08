package mittpaket

import "testing"

func TestHej(t *testing.T) {

	expected := "This will fail"
	if ret := SayHej(); ret != expected {

		t.Errorf("Hej() = %q, want %q", ret, expected)

	}

}
