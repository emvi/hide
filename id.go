package hide

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strings"
)

// ID type that can be used as an replacement for int64.
// It is converted to/from a hash value when marshalled to/from JSON.
// Value 0 is considered null.
type ID int64

// Scan implements the Scanner interface.
func (hideid *ID) Scan(value interface{}) error {
	if value == nil {
		*hideid = 0
		return nil
	}

	id, ok := value.(int64)

	if !ok {
		return errors.New("unexpected type")
	}

	*hideid = ID(id)
	return nil
}

// Value implements the driver Valuer interface.
func (hideid ID) Value() (driver.Value, error) {
	if hideid == 0 {
		return nil, nil
	}

	return int64(hideid), nil
}

// MarshalJSON implements the encoding json interface.
func (hideid ID) MarshalJSON() ([]byte, error) {
	if hideid == 0 {
		return json.Marshal(nil)
	}

	result, err := hash.Encode(hideid)

	if err != nil {
		return nil, err
	}

	return json.Marshal(string(result))
}

// UnmarshalJSON implements the encoding json interface.
func (hideid *ID) UnmarshalJSON(data []byte) error {
	// convert null to 0
	if strings.TrimSpace(string(data)) == "null" {
		*hideid = 0
		return nil
	}

	// remove quotes
	if len(data) >= 2 {
		data = data[1 : len(data)-1]
	}

	result, err := hash.Decode(data)

	if err != nil {
		return err
	}

	*hideid = ID(result)
	return nil
}
