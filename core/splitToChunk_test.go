package core

import (
	"bytes"
	"fmt"
	"io"
	"os"
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
			writerFunc := func() io.WriteCloser {
				buf := &bytes.Buffer{}
				buffers = append(buffers, buf)
				return &BufferWriteCloser{
					Buffer: buf,
				}
			}
			if err := SplitToChunk(r, writerFunc, c.chunk, len(c.in)); err != nil {
				t.Errorf("SplitToChunk(%q)\n expected: %q\n got: %q", c.in, c.want, err)
			}

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

func TestSplitToChunkLargeInput(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	cases := []struct {
		name  string
		in    int
		want  []string
		chunk int
	}{
		{
			name:  "10MB / 512",
			in:    10 * 1024 * 1024,
			want:  []string{},
			chunk: 512,
		},
		{
			name:  "10MB / 1000",
			in:    10 * 1024 * 1024,
			want:  []string{},
			chunk: 1000,
		},
		{
			name:  "1 GB / 512",
			in:    1 * 1024 * 1024 * 1024,
			want:  []string{},
			chunk: 512,
		},
		{
			name:  "10GB / 1MB",
			in:    1 * 1024 * 1024 * 1024,
			want:  []string{},
			chunk: 10 * 1024,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			largeInput := bytes.Repeat([]byte{'a'}, c.in)
			// (c.size - 1) / c.size は 0
			// c.size / c.size とすると余りの有無が考慮できない

			r := bytes.NewReader(largeInput)
			var buffers []*bytes.Buffer
			writerFunc := func() io.WriteCloser {
				buf := &bytes.Buffer{}
				buffers = append(buffers, buf)
				return &BufferWriteCloser{
					Buffer: buf,
				}
			}

			err := SplitToChunk(r, writerFunc, c.chunk, c.in)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if len(buffers) != c.chunk {
				t.Errorf("Expected %d chunks, but got %d", c.chunk, len(buffers))
			}
		})
	}
}

type BufferWriteCloser struct {
	*bytes.Buffer
}

func (bwc *BufferWriteCloser) Close() error {
	return nil
}
