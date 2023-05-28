package utils

import "testing"

type Case struct {
	in, want string
}

var cases = []Case{
	{"hello, world", "HELLO, WORLD"},
	{"HoW ArE YoU DoInG", "HOW ARE YOU DOING"},
	{"", ""},
	{"GOOD", "GOOD"},
}	

func TestToUpper(t *testing.T) {
	for _, c := range cases{
		got := ToUpper(c.in)

		if got != c.want {
			t.Errorf("ToUpper(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}