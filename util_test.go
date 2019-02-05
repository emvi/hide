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
	if id, err := FromString(""); err != nil || id != 0 {
		t.Fatalf("Must return 0 on empty string, but was: %v %v", err, id)
	}
}

func TestToString(t *testing.T) {
	if hash, err := ToString(123); err != nil || hash != "beJarVNaQM" {
		t.Fatalf("Must return hash from ID, but was: %v %v", err, hash)
	}
}

func TestToStringNull(t *testing.T) {
	if hash, err := ToString(0); err != nil || hash != "null" {
		t.Fatalf("Must return 0 from null ID, but was: %v %v", err, hash)
	}
}
