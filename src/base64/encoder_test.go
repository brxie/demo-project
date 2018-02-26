package base64

import "testing"


func TestEncode(t *testing.T) {
	for _, c := range []struct {
		in, want string
	}{
		{"foo", 	"Zm9v"},
		{"bar", 	"YmFy"},
		{"bazzz", 	"YmF6eno="},
	} {
		got := Encode(&c.in)
		if *got != c.want {
			t.Errorf("Encode(%q) == %q, want %q", c.in, *got, c.want)
		}
	}
}