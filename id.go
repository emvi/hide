package hide

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strings"
)

// ID type that can be used as an replacement for uint64.
// It is converted to/from a hash value when marshalled to/from JSON.
// Value 0 is considered null.
type ID uint64

// Scan implements the Scanner interface.
func (hide *ID) Scan(value interface{}) error {
	if value == nil {
		*hide = 0
		return nil
	}

	id, ok := value.(uint64)

	if !ok {
		return errors.New("unexpected type")
	}

	*hide = ID(id)
	return nil
}

// Value implements the driver Valuer interface.
func (hide ID) Value() (driver.Value, error) {
	if hide == 0 {
		return nil, nil
	}

	return uint64(hide), nil
}

// MarshalJSON implements the encoding json interface.
func (hide ID) MarshalJSON() ([]byte, error) {
	if hide == 0 {
		return json.Marshal(nil)
	}

	result, err := hash.Encode(hide)

	if err != nil {
		return nil, err
	}

	return json.Marshal(string(result))
}

// UnmarshalJSON implements the encoding json interface.
func (hide *ID) UnmarshalJSON(data []byte) error {
	// convert null to 0
	if strings.TrimSpace(string(data)) == "null" {
		*hide = 0
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

	*hide = result
	return nil
}
