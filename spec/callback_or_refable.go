// Code generated; DO NOT EDIT.

package spec

import (
	"encoding/json"
	"errors"
)

// Callback It's the Union type of callback and Refable
type Callback struct {
	Refable
	callback
}

// MarshalJSON returns m as the JSON encoding of callback or Refable.
func (m Callback) MarshalJSON() ([]byte, error) {
	if m.Ref != "" {
		return json.Marshal(m.Refable)
	}
	return json.Marshal(m.callback)
}

// UnmarshalJSON sets callback or Refable to data.
func (m *Callback) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("spec.Callback: UnmarshalJSON on nil pointer")
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
	return json.Unmarshal(data, &m.callback)
}
