// Package for testing events in the riverworld
// Possible events: put(item), getin(), getout(), crossriver(), takeout(item)
// start() or setup(), reset() osv.
package event

import "testing"

func TestPut(t *testing.T) {
    // Hva forventer jeg?
    wanted := "[kylling rev korn ---\\ \\_hs_/ _________________/---]"
    got := Put("hs") // Hva fikk jeg?
    if got != wanted {
        t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
    }


}

func TestPut2(t *testing.T) {
    // Hva forventer jeg?
    wanted := "[kylling rev korn ---\\ \\__/ _________________/---hs]"
    got := Put2("hs") // Hva fikk jeg?
    if got != wanted {
        t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
    }


}

func TestCrossriver(t *testing.T) {
    // Hva forventer jeg?
    wanted := "[kylling rev korn ---\\  _________________\\__//---hs]"
    got := Crossriver("hs") // Hva fikk jeg?
    if got != wanted {
        t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
    }


}