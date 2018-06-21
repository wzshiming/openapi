// Code generated; DO NOT EDIT.

package spec

import (
	"encoding/json"
	"errors"
)

// LinkOrRefable It's the Union type of Link and Refable
type LinkOrRefable struct {
	Refable
	Link *Link
}

// MarshalJSON returns m as the JSON encoding of Link or Refable.
func (m LinkOrRefable) MarshalJSON() ([]byte, error) {
	if m.Ref != "" {
		return json.Marshal(m.Refable)
	}
	return json.Marshal(m.Link)
}

// UnmarshalJSON sets Link or Refable to data.
func (m *LinkOrRefable) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("spec.LinkOrRefable: UnmarshalJSON on nil pointer")
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
	return json.Unmarshal(data, &m.Link)
}
