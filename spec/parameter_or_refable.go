// Code generated; DO NOT EDIT.

package spec

import (
	"encoding/json"
	"errors"
)

// Parameter It's the Union type of parameter and Refable
type Parameter struct {
	Refable
	parameter
}

// MarshalJSON returns m as the JSON encoding of parameter or Refable.
func (m Parameter) MarshalJSON() ([]byte, error) {
	if m.Ref != "" {
		return json.Marshal(m.Refable)
	}
	return json.Marshal(m.parameter)
}

// UnmarshalJSON sets parameter or Refable to data.
func (m *Parameter) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("spec.Parameter: UnmarshalJSON on nil pointer")
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
	return json.Unmarshal(data, &m.parameter)
}
