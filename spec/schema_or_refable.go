// Code generated; DO NOT EDIT.

package spec

import (
	"encoding/json"
	"errors"
)

// SchemaOrRefable It's the Union type of Schema and Refable
type SchemaOrRefable struct {
	Refable
	Schema
}

// MarshalJSON returns m as the JSON encoding of Schema or Refable.
func (m SchemaOrRefable) MarshalJSON() ([]byte, error) {
	if m.Ref != "" {
		return json.Marshal(m.Refable)
	}
	return json.Marshal(m.Schema)
}

// UnmarshalJSON sets Schema or Refable to data.
func (m *SchemaOrRefable) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("spec.SchemaOrRefable: UnmarshalJSON on nil pointer")
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
	return json.Unmarshal(data, &m.Schema)
}
