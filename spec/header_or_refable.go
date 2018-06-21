// Code generated; DO NOT EDIT.

package spec

import (
	"encoding/json"
	"errors"
)

// HeaderOrRefable It's the Union type of Header and Refable
type HeaderOrRefable struct {
	Refable
	Header *Header
}

// MarshalJSON returns m as the JSON encoding of Header or Refable.
func (m HeaderOrRefable) MarshalJSON() ([]byte, error) {
	if m.Ref != "" {
		return json.Marshal(m.Refable)
	}
	return json.Marshal(m.Header)
}

// UnmarshalJSON sets Header or Refable to data.
func (m *HeaderOrRefable) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("spec.HeaderOrRefable: UnmarshalJSON on nil pointer")
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
	return json.Unmarshal(data, &m.Header)
}
