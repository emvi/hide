package hide

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// ID type that can be used as an replacement for int64.
// It is converted to/from a hash value when marshalled to/from JSON.
// Value 0 is considered null.
type ID int64

// Scan implements the Scanner interface.
func (this *ID) Scan(value interface{}) error {
	if value == nil {
		*this = 0
		return nil
	}

	id, ok := value.(int64)

	if !ok {
		return errors.New("unexpected type")
	}

	*this = ID(id)
	return nil
}

// Value implements the driver Valuer interface.
func (this ID) Value() (driver.Value, error) {
	if this == 0 {
		return nil, nil
	}

	return int64(this), nil
}

// MarshalJSON implements the encoding json interface.
func (this ID) MarshalJSON() ([]byte, error) {
	if this == 0 {
		return json.Marshal(nil)
	}

	result, err := hash.Encode(this)

	if err != nil {
		return nil, err
	}

	return json.Marshal(string(result))
}

// UnmarshalJSON implements the encoding json interface.
func (this *ID) UnmarshalJSON(data []byte) error {
	// remove quotes
	if len(data) >= 2 {
		data = data[1 : len(data)-1]
	}

	result, err := hash.Decode(data)

	if err != nil {
		return err
	}

	*this = ID(result)
	return nil
}
