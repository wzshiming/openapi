package spec

import (
	"encoding/json"
	"errors"
)

// AnyOrExpressions It's the Union type of any and expressions
type AnyOrExpressions struct {
	Expressions Expressions
	Any         Any
}

// MarshalJSON returns m as the JSON encoding of m.
func (m AnyOrExpressions) MarshalJSON() ([]byte, error) {
	if m.Expressions != "" {
		return json.Marshal(m.Expressions)
	}
	return json.Marshal(m.Any)
}

// UnmarshalJSON sets *m to a copy of data.
func (m *AnyOrExpressions) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("spec.AnyOrExpressions: UnmarshalJSON on nil pointer")
	}
	if len(data) == 0 {
		return nil
	}
	if data[0] == '"' {
		return json.Unmarshal(data, &m.Expressions)
	}
	return json.Unmarshal(data, &m.Any)
}
