package core

import (
	"bufio"
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestSplitByByte(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want []string
		size int
	}{
		{
			name: "empty input",
			in:   "",
			want: []string{},
			size: 1,
		},
		{
			name: "single line",
			in:   "hello\n",
			want: []string{"h", "e", "l", "l", "o", "\n"},
			size: 1,
		},
		{
			name: "single line(divided by 2 bytes)",
			in:   "aa\n",
			want: []string{"aa", "\n"},
			size: 2,
		},
		{
			name: "10 rune",
			in:   "0123456789",
			want: []string{"012", "345", "678", "9"},
			size: 3,
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
			SplitByByte(r, writerFunc, c.size)

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

func TestSplitByByteLargeInput(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	cases := []struct {
		name string
		in   int
		want []string
		size int
	}{
		{
			name: "10MB / 512",
			in:   10 * 1024 * 1024,
			want: []string{},
			size: 512,
		},
		{
			name: "10MB / 1000",
			in:   10 * 1024 * 1024,
			want: []string{},
			size: 1000,
		},
		{
			name: "1 GB / 512",
			in:   1 * 1024 * 1024 * 1024,
			want: []string{},
			size: 512,
		},
		{
			name: "10GB / 1MB",
			in:   1 * 1024 * 1024 * 1024,
			want: []string{},
			size: 10 * 1024,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			largeInput := bytes.Repeat([]byte{'a'}, c.in)
			// (c.size - 1) / c.size は 0
			// c.size / c.size とすると余りの有無が考慮できない
			expectedSplits := (c.in + c.size - 1) / c.size

			r := bytes.NewReader(largeInput)
			var buffers []*bytes.Buffer
			writerFunc := func() io.WriteCloser {
				buf := &bytes.Buffer{}
				buffers = append(buffers, buf)
				return &BufferWriteCloser{
					Buffer: buf,
				}
			}

			SplitByByte(r, writerFunc, c.size)

			if len(buffers) != expectedSplits {
				t.Errorf("Expected %d splits, but got %d", expectedSplits, len(buffers))
			}

			for _, buf := range buffers {
				if buf.Len() != c.size && buf.Len() != c.in%c.size { // 最後のブロックは小さい可能性がある
					t.Errorf("Expected buffer of size %d or %d, but got %d", c.size, c.in%c.size, buf.Len())
				}
			}
		})
	}
}

func TestCustomSplitFunc(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want []string
		size int
	}{
		{name: "empty input", in: "", want: []string{}, size: 5},
		{name: "simple input", in: "hello\n", want: []string{"hello", "\n"}, size: 5},
		{name: "no trailing newline", in: "Lacking EOF new line", want: []string{"Lacki", "ng EO", "F new", " line"}, size: 5},
		{name: "multiple lines", in: "123456789\n123456789\n123456789\n123456789\n123456789\n",
			want: []string{"12345", "6789\n", "12345", "6789\n", "12345", "6789\n", "12345", "6789\n", "12345", "6789\n"}, size: 5},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r := bytes.NewReader([]byte(c.in))
			got := make([]string, 0)
			scanner := bufio.NewScanner(r)
			scanner.Split(split(c.size))

			for scanner.Scan() {
				got = append(got, scanner.Text())
			}
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("SplitBySize(%q)\n expected: %q\n got: %q", c.in, c.want, got)
			}
		})
	}
}
