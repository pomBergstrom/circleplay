package mittpaket

import "testing"

// SayHej : hälsa på alla
func TestSayHej(t *testing.T) {
	//return "Hej hopp"
	svar := SayHej()
	if svar == "kalle anka" {
		t.Error("Fel hälsning")
	}
}
