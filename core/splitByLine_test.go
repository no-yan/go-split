package core

import (
	"bytes"
	"io"
	"testing"
)

func TestSplitByLine(t *testing.T) {
	cases := []struct {
		name  string
		in    string
		want  string
		limit int
	}{
		{
			name:  "empty input",
			in:    "",
			want:  "",
			limit: 1,
		},
		{
			name:  "single line",
			in:    "hello\n",
			want:  "hello\n",
			limit: 1,
		},
		{
			name:  "no trailing newline",
			in:    "Lacking EOF new line",
			want:  "Lacking EOF new line\n",
			limit: 1,
		},
		{
			name:  "multiple lines",
			in:    "some io.Reader stream to be read\nsome io.Reader stream to be read\n",
			want:  "some io.Reader stream to be read\nsome io.Reader stream to be read\n",
			limit: 2,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r := bytes.NewReader([]byte(c.in))
			out := &bytes.Buffer{}
			SplitByLine(r, func() io.Writer { return out }, c.limit)

			if got := out.String(); got != c.want {
				t.Errorf("SplitByLine(%q)\n got: %q\n want: %q\n", c.in, got, c.want)
			}
		})
	}
}
