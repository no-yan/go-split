package core_test

import (
	"bytes"
	"io"
	"reflect"
	"testing"

	"github.com/no-yan/go-split/core"
)

func TestSplitByLine(t *testing.T) {
	cases := []struct {
		name  string
		in    string
		want  []string
		limit int
	}{
		{
			name:  "empty input",
			in:    "",
			want:  []string{},
			limit: 1,
		},
		{
			name:  "single line",
			in:    "hello\n",
			want:  []string{"hello\n"},
			limit: 1,
		},
		{
			name:  "no trailing newline",
			in:    "Lacking EOF new line",
			want:  []string{"Lacking EOF new line\n"},
			limit: 1,
		},
		{
			name: "multiple lines",
			in:   "some io.Reader stream to be read\nsome io.Reader stream to be read\n",
			// 本当は gnu split と挙動を合わせたい
			// want:  []string{"some io.Reader stream to be read\nsome io.Reader stream to be read\n", ""},
			want:  []string{"some io.Reader stream to be read\nsome io.Reader stream to be read\n"},
			limit: 2,
		},
		{
			name:  "multiple lines",
			in:    "a\nb\nc\nd\ne\n",
			want:  []string{"a\nb\n", "c\nd\n", "e\n"},
			limit: 2,
		},
		{
			name:  "empty lines",
			in:    "\n\n\n\n\n\n\n\n\n\n",
			want:  []string{"\n", "\n", "\n", "\n", "\n", "\n", "\n", "\n", "\n", "\n"},
			limit: 1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r := bytes.NewReader([]byte(c.in))
			var buffers []*bytes.Buffer
			writerFunc := func() io.WriteCloser {
				buf := &bytes.Buffer{}
				buffers = append(buffers, buf)
				return &BufferWriteCloser{
					Buffer: buf,
				}
			}
			if err := core.SplitByLine(r, writerFunc, c.limit); err != nil {
				t.Errorf("SplitByLine(%q)\n expected: %q\n got: %q", c.in, c.want, err)
			}

			got := make([]string, len(buffers))
			for i, buf := range buffers {
				got[i] = buf.String()
			}

			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("SplitByLine(%q)\n expected: %q\n got: %q", c.in, c.want, got)
			}
		})
	}
}
