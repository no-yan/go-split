package core

import "testing"

func TestNextAlphabet(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"", "a"}, {"a", "b"}, {"z", "aa"}, {"zz", "aaa"}, {"daz", "dba"},
	}
	for _, c := range cases {
		got := nextAlphabet(c.in)
		if got != c.want {
			t.Errorf("Split(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestNextAlphabetSequence(t *testing.T) {
	got := ""
	prev := ""
	want := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < 26; i++ {
		next := nextAlphabet(prev)
		got += next
		prev = next
	}
	if got != want {
		t.Errorf("nextAlphabet() couldn't generate all alphabets in order, want %q, got %q", want, got)
	}
}
