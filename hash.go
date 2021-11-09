package hide

var hash Hash

// Hash is used to marshal/unmarshal hide.ID to/from JSON.
type Hash interface {
	Encode(ID) ([]byte, error)
	Decode([]byte) (ID, error)
}

// UseHash sets the hide.Hash used to marshal/unmarshal hide.ID to/from JSON.
// hide.HashID is used by default.
func UseHash(hashFunction Hash) {
	hash = hashFunction
}
