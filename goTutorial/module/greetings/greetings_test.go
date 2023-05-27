package greetings

import (
	"regexp"
	"testing"
)


func TestHelloName(t *testing.T) {
	name := "Cyril"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello(name)

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("%v") = %q, %v, want match for %#q, nil`, name, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")

	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

func TestHellosNames(t *testing.T) {
	names := []string{"Cyril", "Nicolas", "Amaury"}
	messages, err := Hellos(names)
	
	if err != nil {
		t.Fatalf(`Hellos(%v) = %q, %v, want matches for %v, nil`, names, messages, err, names)
	}

	if len(names) != len(messages) {
		t.Fatalf(
			`Hello(%v) = %q, %v, Length %v: %d, not matching length %v: %d.`,
			names,
			messages,
			err,
		 	names,
			len(names),
			messages,
			len(messages))
	}

	for _, name := range names {
		want := regexp.MustCompile(`\b`+name+`\b`)

		if !want.MatchString(messages[name]) {
			t.Fatalf(`Hellos(%v) = %q, %v, want matches for %v, nil`, names, messages, err, name)
		}
	}
}