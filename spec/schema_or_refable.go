// Code generated; DO NOT EDIT.

package spec

import (
	"encoding/json"
	"errors"
)

// Schema It's the Union type of schema and Refable
type Schema struct {
	Refable
	schema
}

// MarshalJSON returns m as the JSON encoding of schema or Refable.
func (m Schema) MarshalJSON() ([]byte, error) {
	if m.Ref != "" {
		return json.Marshal(m.Refable)
	}
	return json.Marshal(m.schema)
}

// UnmarshalJSON sets schema or Refable to data.
func (m *Schema) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("spec.Schema: UnmarshalJSON on nil pointer")
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
	return json.Unmarshal(data, &m.schema)
}
