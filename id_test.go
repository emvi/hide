package hide

import (
	"database/sql/driver"
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

func TestScan(t *testing.T) {
	var id ID
	value := int64(123)

	if err := id.Scan(value); err != nil {
		t.Fatal(err)
	}

	if id != 123 {
		t.Fatalf("ID must have been set to value by scan, but was: %v", id)
	}
}

func TestValue(t *testing.T) {
	id := ID(123)
	driverValue, err := id.Value()

	if err != nil {
		t.Fatal(err)
	}

	_, ok := driverValue.(int64)

	if !ok {
		t.Fatal("Driver value must be of type int64")
	}
}

func TestNull(t *testing.T) {
	var id ID
	out, _ := id.MarshalJSON()
	expected := "null"

	if string(out) != expected {
		t.Fatalf("Expected null ID to be '%v', but was: %v", expected, string(out))
	}

	value, _ := id.Value()

	if value != driver.Value(nil) {
		t.Fatalf("Expected null ID to be driver.Value nil, but was: %v", value)
	}
}

func TestUnmarshalNull(t *testing.T) {
	in := `{"id":null}`
	out := &struct {
		Id ID `json:"id"`
	}{}

	if err := json.Unmarshal([]byte(in), out); err != nil || out.Id != 0 {
		t.Fatalf("Expected null to be unmarshalled to 0, but was: %v %v", err, out.Id)
	}
}
