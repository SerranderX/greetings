package greetings

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Juan"
	want := regexp.MustCompile(name)
	message, err := Hello("Juan")

	if !want.MatchString(message) || err != nil {
		t.Fatalf(`Hello("Juan") = %q, %v, quiere coincidencia para %#q, nil`, message, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	_, err := Hello("")
	if err == nil {
		t.Fatal("Hello(\"\") = nil, quiere error para \"\"")
	}
}

func TestHellos(t *testing.T) {
    // Test case: Multiple names
    names := []string{"Gladys", "Samantha", "Darrin"}
    messages, err := Hellos(names)
    if err != nil {
        t.Errorf("Hellos(%v) returned an error: %v", names, err)
    }

    // Check that each name has a message
    for _, name := range names {
        if _, ok := messages[name]; !ok {
            t.Errorf("Hellos(%v) did not return a message for %v", names, name)
        }
    }
}
