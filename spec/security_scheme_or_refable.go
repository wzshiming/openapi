// Code generated; DO NOT EDIT.

package spec

import (
	"encoding/json"
	"errors"
)

// SecurityScheme It's the Union type of securityScheme and Refable
type SecurityScheme struct {
	Refable
	securityScheme
}

// MarshalJSON returns m as the JSON encoding of securityScheme or Refable.
func (m SecurityScheme) MarshalJSON() ([]byte, error) {
	if m.Ref != "" {
		return json.Marshal(m.Refable)
	}
	return json.Marshal(m.securityScheme)
}

// UnmarshalJSON sets securityScheme or Refable to data.
func (m *SecurityScheme) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("spec.SecurityScheme: UnmarshalJSON on nil pointer")
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
	return json.Unmarshal(data, &m.securityScheme)
}
