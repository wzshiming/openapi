package util

import (
	"encoding/json"
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

func YAML2JSON(d []byte) ([]byte, error) {
	ms := yaml.MapSlice{}
	err := yaml.Unmarshal(d, &ms)
	if err != nil {
		return nil, err
	}

	v := _YAML2JSON(ms)
	return json.Marshal(v)
}

func _YAML2JSON(v interface{}) interface{} {
	switch t := v.(type) {
	case yaml.MapSlice:
		m := map[string]interface{}{}
		for _, v := range t {
			key := fmt.Sprint(v.Key)
			val := _YAML2JSON(v.Value)
			m[key] = val
		}
		return m
	case []interface{}:
		for i, v := range t {
			t[i] = _YAML2JSON(v)
		}
		return t
	default:
		return v
	}
}
