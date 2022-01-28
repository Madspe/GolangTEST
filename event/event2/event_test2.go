package event

import "testing"

func TestPutIØst(t *testing.T) {
    // Hva forventer jeg?
    wanted := "[kylling rev korn ---\\ \\__/ _________________/---korn]"
    got := PutIØst("korn") // Hva fikk jeg?
    if got != wanted {
             t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
    }

    }