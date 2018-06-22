// Code generated; DO NOT EDIT.

package spec

import (
	"encoding/json"
	"errors"
)

// Response It's the Union type of response and Refable
type Response struct {
	Refable
	response
}

// MarshalJSON returns m as the JSON encoding of response or Refable.
func (m Response) MarshalJSON() ([]byte, error) {
	if m.Ref != "" {
		return json.Marshal(m.Refable)
	}
	return json.Marshal(m.response)
}

// UnmarshalJSON sets response or Refable to data.
func (m *Response) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("spec.Response: UnmarshalJSON on nil pointer")
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
	return json.Unmarshal(data, &m.response)
}
