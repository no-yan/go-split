package core

import "testing"

var empty []string

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func Test(t *testing.T) {
	cases := []struct {
		in   string
		want []string
	}{
		{"", []string{}},
		{"hello", []string{"hello"}},
	}
	for _, c := range cases {
		got := Split(c.in)
		if !Equal(got, c.want) {
			t.Errorf("Split(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
