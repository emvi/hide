package hide

// Returns a new ID from given hash by using the hasher or an error if it couldn't decode the hash.
func FromString(id string) (ID, error) {
	return hash.Decode([]byte(id))
}

// Returns a new hash from given ID by using the hasher or an error if it couldn't encode the ID.
// If ID is 0, "null" will be returned.
func ToString(id ID) (string, error) {
	if id == 0 {
		return "null", nil
	}

	hash, err := hash.Encode(id)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}
