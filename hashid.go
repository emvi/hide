package hide

import (
	"errors"
	"github.com/speps/go-hashids"
)

// HashID implements the hide.Hash interface and uses github.com/speps/go-hashids to encode and decode hashes.
type HashID struct {
	Salt      string
	MinLength int
}

// Creates a new HashID with given salt and minimum hash length.
func NewHashID(salt string, minlen int) *HashID {
	return &HashID{salt, minlen}
}

// Encode implements the hide.Hash interface.
func (this *HashID) Encode(id ID) ([]byte, error) {
	hash, err := this.newHash()

	if err != nil {
		return nil, err
	}

	result, err := hash.EncodeInt64([]int64{int64(id)})

	if err != nil {
		return nil, err
	}

	return []byte(result), nil
}

// Decode implements the hide.Hash interface.
func (this *HashID) Decode(data []byte) (ID, error) {
	hash, err := this.newHash()

	if err != nil {
		return 0, err
	}

	result, err := hash.DecodeInt64WithError(string(data))

	if err != nil {
		return 0, err
	}

	if len(result) != 1 {
		return 0, errors.New("input value too long")
	}

	return ID(result[0]), nil
}

// Creates a new hashids.HashID object to encode/decode IDs.
func (this *HashID) newHash() (*hashids.HashID, error) {
	config := hashids.NewData()
	config.Salt = this.Salt
	config.MinLength = this.MinLength
	hash, err := hashids.NewWithData(config)

	if err != nil {
		return nil, err
	}

	return hash, nil
}
