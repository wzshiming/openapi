package util

import (
	"encoding/json"
	"errors"
)

// Unmarshal Decode the format supported by openapi
func Unmarshal(d []byte, v interface{}) (err error) {
	d, err = YAML2JSON(d)
	if err != nil {
		return err
	}
	return json.Unmarshal(d, v)
}

// Marshal Encode the format supported by openapi
func Marshal(v interface{}, format Format) (d []byte, err error) {
	switch format {
	case JSON:
		return json.Marshal(v)
	case JSONIndent:
		return json.MarshalIndent(v, "", "  ")
	case YAML:
		d, err = json.Marshal(v)
		if err != nil {
			return nil, err
		}
		return YAML2JSON(d)
	}
	return nil, errors.New("Bad format")
}

// Format for openapi
type Format uint8

const (
	JSON Format = iota
	JSONIndent
	YAML
)
