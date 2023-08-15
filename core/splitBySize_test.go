package core

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"
)

// func TestBySize(t *testing.T) {
// 	cases := []struct {
// 		in   string
// 		want string
// 	}{
// 		{"", ""},
// 		{"hello\n", "hello\n\n"},
// 		{"Lacking EOF new line", "Lacki\nng EO\nF new\n line\n"},
// 		{"some io.Reader stream to be read\nsome io.Reader stream to be read\nsome io.Reader stream to be read\nsome io.Reader stream to be read\n", "some \nio.Re\nader \nstrea\nm to \nbe re\nad\nso\nme io\n.Read\ner st\nream to be read\nsome io.Reader stream to be read\nsome io.Reader stream to be read\n"},
// 	}
// 	for _, c := range cases {
// 		r := bytes.NewReader([]byte(c.in))
// 		out := &bytes.Buffer{}
// 		SplitBySize(r, func() io.Writer { return out }, 5)

// 		if got := out.String(); got != c.want {
// 			t.Errorf("Split(%q) == %q, want %q", c.in, got, c.want)
// 		}
// 	}
// }

func TestSplitFunction(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"hello\n", "hello\n\n\n"},
		{"Lacking EOF new line", "Lacki\nng EO\nF new\n line\n"},
		{"some io.Reader stream to be read\nsome io.Reader stream to be read\nsome io.Reader stream to be read\nsome io.Reader stream to be read\n", "some \nio.Re\nader \nstrea\nm to \nbe re\nad\nso\nme io\n.Read\ner st\nream \nto be\n read\n\nsome\n io.R\neader\n stre\nam to\n be r\nead\ns\nome i\no.Rea\nder s\ntream\n to b\ne rea\nd\n\n"},
	}
	for _, c := range cases {
		r := bytes.NewReader([]byte(c.in))
		out := &bytes.Buffer{}
		scanner := bufio.NewScanner(r)
		scanner.Split(split(5))

		for scanner.Scan() {
			txt := scanner.Text()
			if len(txt) > 0 {
				fmt.Fprintln(out, txt)
			}
		}
		if got := out.String(); got != c.want {
			t.Errorf("Split(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
