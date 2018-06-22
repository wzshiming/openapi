// Code generated; DO NOT EDIT.

package spec

import (
	"encoding/json"
	"errors"
)

// RequestBodyOrRefable It's the Union type of RequestBody and Refable
type RequestBodyOrRefable struct {
	Refable
	RequestBody
}

// MarshalJSON returns m as the JSON encoding of RequestBody or Refable.
func (m RequestBodyOrRefable) MarshalJSON() ([]byte, error) {
	if m.Ref != "" {
		return json.Marshal(m.Refable)
	}
	return json.Marshal(m.RequestBody)
}

// UnmarshalJSON sets RequestBody or Refable to data.
func (m *RequestBodyOrRefable) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("spec.RequestBodyOrRefable: UnmarshalJSON on nil pointer")
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
	return json.Unmarshal(data, &m.RequestBody)
}
