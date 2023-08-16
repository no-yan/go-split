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
			name: "single line(devided by 2 bytes)",
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
