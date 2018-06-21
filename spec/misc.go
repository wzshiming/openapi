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
	ss := []string{}
	err := json.Unmarshal(data, &ss)
	if err != nil {
		return err
	}
	*m = StringOrArray(ss)
	return nil
}

type SchemaOrArray []Schema

type SchemaOrBool struct {
	Allows bool
	Schema *Schema
}

type RequestBodyOrString struct {
	Literal     string
	RequestBody *RequestBody
}

type Header Parameter

type Callback map[string]PathItem
