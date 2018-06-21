//// +build ignore
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"text/template"
)

var x = flag.String("x", "Schema", "type name")
var f = flag.String("f", "", "output file")
var p = flag.String("p", "spec", "package name")

func main() {
	flag.Parse()
	t, _ := template.New("").Parse(temp)
	buf := bytes.NewBuffer(nil)
	t.Execute(buf, map[string]string{
		"X":   *x,
		"Pkg": *p,
	})

	file := buf.Bytes()
	file, _ = format.Source(file)
	if *f == "" {
		fmt.Print(string(file))
		return
	}
	ioutil.WriteFile(*f, file, 0666)
	return
}

const temp = `
// Code generated; DO NOT EDIT.

package {{ .Pkg }}

import (
	"errors"
	"encoding/json"
)

// {{ .X }}OrRefable It's the Union type of {{ .X }} and Refable
type {{ .X }}OrRefable struct {
	Refable
	{{ .X }} *{{ .X }}
}

// MarshalJSON returns m as the JSON encoding of {{ .X }} or Refable.
func (m {{ .X }}OrRefable) MarshalJSON() ([]byte, error) {
	if m.Ref != "" {
		return json.Marshal(m.Refable)
	}
	return json.Marshal(m.{{ .X }})
}

// UnmarshalJSON sets {{ .X }} or Refable to data.
func (m *{{ .X }}OrRefable) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("spec.{{ .X }}OrRefable: UnmarshalJSON on nil pointer")
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
	return json.Unmarshal(data, &m.{{ .X }})
}

`
