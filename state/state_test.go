// Package for testing the represention of the state of the riverworld
package state

import "testing"

// Test function implemented in line with the Golang's testing tool
func TestViewState(t *testing.T) {
    wanted := "[kylling rev korn hs ---\\ \\__/ _________________/---]"
    state := ViewState();
    if state != wanted {
         t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
    }
}

func TestViewState2(t *testing.T) {
    wanted := "[kylling rev korn  ---\\ \\_hs_/ _________________/---]"
    state := ViewState2();
    if state != wanted {
         t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
    }
}

func TestViewState3(t *testing.T) {
    wanted := "[kylling rev korn ---\\  _________________\\__//---hs]"
    state := ViewState3();
    if state != wanted {
         t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
    }
}