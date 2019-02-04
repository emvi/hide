package hide

import (
	"testing"
)

func TestFromString(t *testing.T) {
	if id, err := FromString("invalid"); err == nil || id != 0 {
		t.Fatalf("Must not return ID from string, but was: %v %v", err, id)
	}

	if id, err := FromString("beJarVNaQM"); err != nil || id != 123 {
		t.Fatalf("Must return ID from string, but was: %v %v", err, id)
	}
}

func TestFromStringEmpty(t *testing.T) {
	if id, err := FromString(""); err.Error() != "input value empty" || id != 0 {
		t.Fatalf("Must return error on empty string, but was: %v %v", err, id)
	}
}
