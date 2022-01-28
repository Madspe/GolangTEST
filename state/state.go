package state
// alle mulige states
//"[kylling rev korn hs ---\\ \\__/ _________________/---]"
//"[kylling rev  ---\\ \\_korn hs_/ _________________/---]"
//"[kylling  korn  ---\\ \\_rev hs_/ _________________/---]"
//"[kylling rev korn hs ---\\ \\_kylling hs_/ _________________/---]"
//"[kylling rev  ---\\ \\__/ _________________/---korn hs]"
//"[kylling korn ---\\ \\__/ _________________/---rev hs]"
//"[rev korn ---\\ \\__/ _________________/---kylling hs]"
//"[kylling rev korn hs ---\\ \\__/ _________________/---kylling rev korn hs]"






func ViewState() string {
    // Sjekke data som er lagret ("kylling til venstre", "rev til venstre")
    return "[kylling rev korn hs ---\\ \\__/ _________________/---]"
}

func ViewState2() string {
    // Prøver ut flere states
    return "[kylling rev korn  ---\\ \\_hs_/ _________________/---]"
}

func ViewState3() string {
    //prøver ut state flyttebåten
    return "[kylling rev korn ---\\  _________________\\__//---hs]"
}