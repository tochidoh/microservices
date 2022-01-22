package data

import (
	"encoding/json"
	"io"
)

// ToJSON serializes the given interface into a string based JSON format
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)

	return e.Encode(i)
}

// FromJSON deserializes JSON string to some object
func FromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)

	return d.Decode(i)
}
