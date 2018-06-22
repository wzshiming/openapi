// Code generated; DO NOT EDIT.

package spec

import (
	"encoding/json"
	"errors"
)

// CallbackOrRefable It's the Union type of Callback and Refable
type CallbackOrRefable struct {
	Refable
	Callback
}

// MarshalJSON returns m as the JSON encoding of Callback or Refable.
func (m CallbackOrRefable) MarshalJSON() ([]byte, error) {
	if m.Ref != "" {
		return json.Marshal(m.Refable)
	}
	return json.Marshal(m.Callback)
}

// UnmarshalJSON sets Callback or Refable to data.
func (m *CallbackOrRefable) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("spec.CallbackOrRefable: UnmarshalJSON on nil pointer")
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
	return json.Unmarshal(data, &m.Callback)
}
