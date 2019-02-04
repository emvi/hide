package hide

// Returns a new ID from given hash by using the hasher or an error if it couldn't decode the hash.
func FromString(id string) (ID, error) {
	return hash.Decode([]byte(id))
}
