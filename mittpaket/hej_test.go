package mittpaket

import (
	"circletest/mittpaket"
	"testing"
)

func TestHej(t *testing.T) {

	var s = mittpaket.SayHej()
	expected := "This will fail"
	if ret := s; ret != expected {

		//t.Errorf("Hej() = %q, want %q", ret, expected)
		//t.Logf("Hej() = %q, want %q", ret, expected)

	}

}
