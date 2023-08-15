package core

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestSplitByChunk(t *testing.T) {
	cases := []struct {
		name  string
		in    string
		want  []string
		chunk int
	}{
		{
			name:  "empty input",
			in:    "",
			want:  []string{},
			chunk: 3,
		},
		{
			name:  "single line",
			in:    "hello\n",
			want:  []string{"he", "ll", "o\n"},
			chunk: 3,
		},
		{
			name:  "10 rune",
			in:    "0123456789",
			want:  []string{"012", "345", "6789"},
			chunk: 3,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r := bytes.NewReader([]byte(c.in))
			var buffers []*bytes.Buffer
			writerFunc := func() io.Writer {
				buf := &bytes.Buffer{}
				buffers = append(buffers, buf)
				return buf
			}
			SplitByChunk(r, writerFunc, c.chunk)

			got := make([]string, len(buffers))
			for i, buf := range buffers {
				got[i] = buf.String()
			}

			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("SplitBySize(%q)\n expected: %q\n got: %q", c.in, c.want, got)
			}
		})
	}
}
