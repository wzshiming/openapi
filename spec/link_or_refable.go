// Code generated; DO NOT EDIT.

package spec

import (
	"encoding/json"
	"errors"
)

// Link It's the Union type of link and Refable
type Link struct {
	Refable
	link
}

// MarshalJSON returns m as the JSON encoding of link or Refable.
func (m Link) MarshalJSON() ([]byte, error) {
	if m.Ref != "" {
		return json.Marshal(m.Refable)
	}
	return json.Marshal(m.link)
}

// UnmarshalJSON sets link or Refable to data.
func (m *Link) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("spec.Link: UnmarshalJSON on nil pointer")
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
	return json.Unmarshal(data, &m.link)
}
