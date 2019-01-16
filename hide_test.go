package hide

import (
	"encoding/json"
	"testing"
)

type testStruct struct {
	Id   ID     `json:"id"`
	Test string `json:"test"`
}

func TestMarshalID(t *testing.T) {
	id := ID(123)
	expected := `"beJarVNaQM"`
	out, err := json.Marshal(id)

	if err != nil {
		t.Fatal(err)
	}

	if string(out) != expected {
		t.Fatalf("Expected marshalled ID to be %v, but was: %v", expected, string(out))
	}
}

func TestUnmarshalID(t *testing.T) {
	in := `"beJarVNaQM"`
	var id ID
	expected := ID(123)
	err := id.UnmarshalJSON([]byte(in))

	if err != nil {
		t.Fatal(err)
	}

	if id != expected {
		t.Fatalf("Expected unmarshalled ID to be %v, but was: %v", expected, id)
	}
}

func TestMarshalUnmarshalMatch(t *testing.T) {
	id := ID(123)
	marshalled, err := id.MarshalJSON()

	if err != nil {
		t.Fatal(err)
	}

	id = ID(0)

	if err := id.UnmarshalJSON(marshalled); err != nil {
		t.Fatal(err)
	}

	if id != ID(123) {
		t.Fatalf("Marshal unmarshal mismatch, expected %v but was: %v", ID(123), id)
	}
}

func TestMarshalIDStruct(t *testing.T) {
	in := testStruct{ID(123), "struct"}
	expected := `{"id":"beJarVNaQM","test":"struct"}`
	out, err := json.Marshal(in)

	if err != nil {
		t.Fatal(err)
	}

	if string(out) != expected {
		t.Fatalf("Expected marshalled struct to be %v, but was: %v", expected, string(out))
	}
}

func TestUnmarshalIDStruct(t *testing.T) {
	in := `{"id":"beJarVNaQM","test":"struct"}`
	var out testStruct
	expected := testStruct{123, "struct"}
	err := json.Unmarshal([]byte(in), &out)

	if err != nil {
		t.Fatal(err)
	}

	if out != expected {
		t.Fatalf("Expected unmarshalled struct to be %v, but was: %v", expected, out)
	}
}
