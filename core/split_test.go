package core

import (
	"bytes"
	"testing"
)

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
		want string
	}{
		{"", ""},
		{"hello", "hello"},
	}
	for _, c := range cases {
		r := bytes.NewReader([]byte(c.in))
		out := &bytes.Buffer{}
		Split(r, out)

		if got := out.String(); got != c.want {
			t.Errorf("Split(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
