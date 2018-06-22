// Code generated; DO NOT EDIT.

package spec

import (
	"encoding/json"
	"errors"
)

// Example It's the Union type of example and Refable
type Example struct {
	Refable
	example
}

// MarshalJSON returns m as the JSON encoding of example or Refable.
func (m Example) MarshalJSON() ([]byte, error) {
	if m.Ref != "" {
		return json.Marshal(m.Refable)
	}
	return json.Marshal(m.example)
}

// UnmarshalJSON sets example or Refable to data.
func (m *Example) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("spec.Example: UnmarshalJSON on nil pointer")
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
	return json.Unmarshal(data, &m.example)
}
