// Code generated; DO NOT EDIT.

package spec

import (
	"encoding/json"
	"errors"
)

// ExampleOrRefable It's the Union type of Example and Refable
type ExampleOrRefable struct {
	Refable
	Example *Example
}

// MarshalJSON returns m as the JSON encoding of Example or Refable.
func (m ExampleOrRefable) MarshalJSON() ([]byte, error) {
	if m.Ref != "" {
		return json.Marshal(m.Refable)
	}
	return json.Marshal(m.Example)
}

// UnmarshalJSON sets Example or Refable to data.
func (m *ExampleOrRefable) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("spec.ExampleOrRefable: UnmarshalJSON on nil pointer")
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
	return json.Unmarshal(data, &m.Example)
}
