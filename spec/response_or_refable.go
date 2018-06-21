// Code generated; DO NOT EDIT.

package spec

import (
	"encoding/json"
	"errors"
)

// ResponseOrRefable It's the Union type of Response and Refable
type ResponseOrRefable struct {
	Refable
	Response *Response
}

// MarshalJSON returns m as the JSON encoding of Response or Refable.
func (m ResponseOrRefable) MarshalJSON() ([]byte, error) {
	if m.Ref != "" {
		return json.Marshal(m.Refable)
	}
	return json.Marshal(m.Response)
}

// UnmarshalJSON sets Response or Refable to data.
func (m *ResponseOrRefable) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("spec.ResponseOrRefable: UnmarshalJSON on nil pointer")
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
	return json.Unmarshal(data, &m.Response)
}
