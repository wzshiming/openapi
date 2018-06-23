package spec

import (
	"encoding/json"
	"errors"
)

// NewAny Create an any
func NewAny(v interface{}) Any {
	s, _ := json.Marshal(v)
	return Any(s)
}

// Any is a raw encoded JSON value.
// It implements Marshaler and Unmarshaler and can
// be used to delay JSON decoding or precompute a JSON encoding.
type Any []byte

// MarshalJSON returns m as the JSON encoding of m.
func (m Any) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *Any) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("Any: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[0:0], data...)
	return nil
}
