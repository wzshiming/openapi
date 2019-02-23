package util

import (
	"encoding/json"
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

func JSON2YAML(d []byte) ([]byte, error) {
	ms := yaml.MapSlice{}
	err := yaml.Unmarshal(d, &ms)
	if err != nil {
		return nil, err
	}
	return yaml.Marshal(ms)
}

func YAML2JSON(d []byte) ([]byte, error) {
	ms := yaml.MapSlice{}
	err := yaml.Unmarshal(d, &ms)
	if err != nil {
		return nil, err
	}

	v := yaml2json(ms)
	return json.MarshalIndent(v, "", "  ")
}

func yaml2json(v interface{}) interface{} {
	switch t := v.(type) {
	case yaml.MapSlice:
		m := map[string]interface{}{}
		for _, v := range t {
			key := fmt.Sprint(v.Key)
			val := yaml2json(v.Value)
			m[key] = val
		}
		return m
	case []interface{}:
		for i, v := range t {
			t[i] = yaml2json(v)
		}
		return t
	default:
		return v
	}
}
