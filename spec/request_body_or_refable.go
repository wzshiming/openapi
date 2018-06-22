// Code generated; DO NOT EDIT.

package spec

import (
	"encoding/json"
	"errors"
)

// RequestBody It's the Union type of requestBody and Refable
type RequestBody struct {
	Refable
	requestBody
}

// MarshalJSON returns m as the JSON encoding of requestBody or Refable.
func (m RequestBody) MarshalJSON() ([]byte, error) {
	if m.Ref != "" {
		return json.Marshal(m.Refable)
	}
	return json.Marshal(m.requestBody)
}

// UnmarshalJSON sets requestBody or Refable to data.
func (m *RequestBody) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("spec.RequestBody: UnmarshalJSON on nil pointer")
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
	return json.Unmarshal(data, &m.requestBody)
}
