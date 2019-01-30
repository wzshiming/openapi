package spec

import (
	"encoding/json"
	"errors"
	"strings"
)

// Refable A simple object to allow referencing other components in the specification, internally and externally.
// The Reference Object is defined by JSON Reference and follows the same structure, behavior and rules.
// For this specification, reference resolution is accomplished as defined by the JSON Reference specification and not by the JSON Schema specification.
type Refable struct {
	// REQUIRED.
	// The reference string.
	Ref        string
	Components string
}

// MarshalJSON returns m as the JSON encoding of m.
func (m Refable) MarshalJSON() ([]byte, error) {
	if m.Ref == "" {
		return []byte("{}"), nil
	}
	return json.Marshal(ref{m.Components + m.Ref})
}

// UnmarshalJSON sets *m to a copy of data.
func (m *Refable) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("spec.Refable: UnmarshalJSON on nil pointer")
	}
	ref := ref{}
	err := json.Unmarshal(data, &ref)
	if err != nil {
		return err
	}
	i := strings.LastIndex(ref.Ref, "/")
	m.Components = ref.Ref[:i+1]
	m.Ref = ref.Ref[i+1:]
	return nil
}

type ref struct {
	Ref string `json:"$ref,omitempty"`
}
