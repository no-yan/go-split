package core

import "testing"

func Test(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"", ""},
		{"hello", "hello"},
	}
	for _, c := range cases {
		got := Split(c.in)
		if got != c.want {
			t.Errorf("Split(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
