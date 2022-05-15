package reverse_string

import (
	"testing"
)

func TestTask1SimpleString(t *testing.T) {
	sourceData := "Hello world!"
	res := ReverseString(sourceData)

	if res != "!dlrow olleH" {
		t.Errorf("Result was incorect, got: %s, want: %s.", res, "!dlrow olleH")
	}
}

func TestTask1MultilineString(t *testing.T) {
	sourceData := "This is a\nmultiline\nstring."
	res := ReverseString(sourceData)
	if res != ".gnirts\nenilitlum\na si sihT" {
		t.Errorf("Result was incorect, got: %s, want: %s.", res, ".gnirts\n        enilitlum\n        a si sihT")
	}
}
