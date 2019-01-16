package hide

import (
	"encoding/json"
)

type ID int64

func (this ID) MarshalJSON() ([]byte, error) {
	result, err := hash.Encode(this)

	if err != nil {
		return nil, err
	}

	return json.Marshal(string(result))
}

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
