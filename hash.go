package hide

var hash Hash

// Interface used to marshal/unmarshal hide.ID to/from JSON.
type Hash interface {
	Encode(ID) ([]byte, error)
	Decode([]byte) (ID, error)
}

// Sets the hide.Hash used to marshal/unmarshal hide.ID to/from JSON.
// hide.HashID is used by default.
func UseHash(hasher Hash) {
	hash = hasher
}
