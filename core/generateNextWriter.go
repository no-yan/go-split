package core

func nextAlphabet(prev string) string {
	if prev == "" {
		return "a"
	}

	headRunes, trailingRune := prev[:len(prev)-1], prev[len(prev)-1]
	if trailingRune == 'z' {
		return nextAlphabet(headRunes) + "a"
	}
	return headRunes + string(trailingRune+1)
}
