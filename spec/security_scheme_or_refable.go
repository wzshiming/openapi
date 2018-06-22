// Code generated; DO NOT EDIT.

package spec

import (
	"encoding/json"
	"errors"
)

// SecuritySchemeOrRefable It's the Union type of SecurityScheme and Refable
type SecuritySchemeOrRefable struct {
	Refable
	SecurityScheme
}

// MarshalJSON returns m as the JSON encoding of SecurityScheme or Refable.
func (m SecuritySchemeOrRefable) MarshalJSON() ([]byte, error) {
	if m.Ref != "" {
		return json.Marshal(m.Refable)
	}
	return json.Marshal(m.SecurityScheme)
}

// UnmarshalJSON sets SecurityScheme or Refable to data.
func (m *SecuritySchemeOrRefable) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("spec.SecuritySchemeOrRefable: UnmarshalJSON on nil pointer")
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
	return json.Unmarshal(data, &m.SecurityScheme)
}
