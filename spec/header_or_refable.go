// Code generated; DO NOT EDIT.

package spec

import (
	"encoding/json"
	"errors"
)

// Header It's the Union type of header and Refable
type Header struct {
	Refable
	header
}

// MarshalJSON returns m as the JSON encoding of header or Refable.
func (m Header) MarshalJSON() ([]byte, error) {
	if m.Ref != "" {
		return json.Marshal(m.Refable)
	}
	return json.Marshal(m.header)
}

// UnmarshalJSON sets header or Refable to data.
func (m *Header) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("spec.Header: UnmarshalJSON on nil pointer")
	}
	if len(data) == 0 {
		return nil
	}
	err := json.Unmarshal(data, &m.Refable)
	if err != nil {
		return err
	}
	if m.Ref != "" {
		return nil
	}
	return json.Unmarshal(data, &m.header)
}
