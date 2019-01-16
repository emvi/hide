package hide

var hash Hash

type Hash interface {
	Encode(ID) ([]byte, error)
	Decode([]byte) (ID, error)
}

func UseHash(hasher Hash) {
	hash = hasher
}
