package core

import (
	"os"
	"strings"
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
		want []string
	}{
		{"", []string{}},
		{"hello", []string{}},
	}
	for _, c := range cases {
		r := strings.NewReader(c.in)
		got := Split(r, os.Stdout)
		if !Equal(got, c.want) {
			t.Errorf("Split(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
