// Package for testing events in the riverworld
// possible events: put(items), getin(). getout(). crossriver(), takeout(items)
// start() or setup(), reset() osv.

package event

import "testing"

func TestPut(t *testing.T) {
    wanted := "[kylling rev korn ---\\ \\_hs_/ ________________/---]"
    got := Put(korn)
    if state != wanted {
        t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
        }


    }