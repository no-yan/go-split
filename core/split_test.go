package core

import (
	"bytes"
	"io"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"hello\n", "hello\n"},
		{"Lacking EOF new line", "Lacking EOF new line\n"},
		{"some io.Reader stream to be read\nsome io.Reader stream to be read\nsome io.Reader stream to be read\nsome io.Reader stream to be read\n", "some io.Reader stream to be read\nsome io.Reader stream to be read\nsome io.Reader stream to be read\nsome io.Reader stream to be read\n"},
	}
	for _, c := range cases {
		r := bytes.NewReader([]byte(c.in))
		out := &bytes.Buffer{}
		Split(r, func() io.Writer { return out })

		if got := out.String(); got != c.want {
			t.Errorf("Split(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
