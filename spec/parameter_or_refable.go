// Code generated; DO NOT EDIT.

package spec

import (
	"encoding/json"
	"errors"
)

// ParameterOrRefable It's the Union type of Parameter and Refable
type ParameterOrRefable struct {
	Refable
	Parameter *Parameter
}

// MarshalJSON returns m as the JSON encoding of Parameter or Refable.
func (m ParameterOrRefable) MarshalJSON() ([]byte, error) {
	if m.Ref != "" {
		return json.Marshal(m.Refable)
	}
	return json.Marshal(m.Parameter)
}

// UnmarshalJSON sets Parameter or Refable to data.
func (m *ParameterOrRefable) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("spec.ParameterOrRefable: UnmarshalJSON on nil pointer")
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
	return json.Unmarshal(data, &m.Parameter)
}
