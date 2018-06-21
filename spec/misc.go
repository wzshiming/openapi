package spec

import (
	"encoding/json"
	"errors"
)

type StringOrArray []string

// MarshalJSON returns m as the JSON encoding of m.
func (m StringOrArray) MarshalJSON() ([]byte, error) {
	return json.Marshal([]string(m))
}

// UnmarshalJSON sets *m to a copy of data.
func (m *StringOrArray) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("spec.StringOrArray: UnmarshalJSON on nil pointer")
	}
	if len(data) == 0 {
		return nil
	}
	if data[0] != '[' {
		s := ""
		err := json.Unmarshal(data, &s)
		if err != nil {
			return err
		}
		*m = []string{s}
		return nil
	}

	return json.Unmarshal(data, []string(*m))
}

type SchemaOrArray []SchemaOrRefable

// MarshalJSON returns m as the JSON encoding of m.
func (m SchemaOrArray) MarshalJSON() ([]byte, error) {
	return json.Marshal([]SchemaOrRefable(m))
}

// UnmarshalJSON sets *m to a copy of data.
func (m *SchemaOrArray) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("spec.SchemaOrArray: UnmarshalJSON on nil pointer")
	}
	if len(data) == 0 {
		return nil
	}
	if data[0] != '[' {
		s := SchemaOrRefable{}
		err := json.Unmarshal(data, &s)
		if err != nil {
			return err
		}
		*m = []SchemaOrRefable{s}
		return nil
	}
	return json.Unmarshal(data, []SchemaOrRefable(*m))
}

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
		return errors.New("spec.StringOrArray: UnmarshalJSON on nil pointer")
	}
	if len(data) == 0 {
		return nil
	}
	if data[0] == '"' {
		return json.Unmarshal(data, &m.Expressions)
	}
	return json.Unmarshal(data, &m.Any)
}
