// +build ignore
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"text/template"
)

var s = flag.String("s", "", "src type name")
var d = flag.String("d", "", "dist type name")
var f = flag.String("f", "", "output file")
var p = flag.String("p", "spec", "package name")

func main() {
	flag.Parse()
	if *s == "" || *d == "" {
		flag.PrintDefaults()
		return
	}
	t, _ := template.New("").Parse(temp)
	buf := bytes.NewBuffer(nil)
	t.Execute(buf, map[string]string{
		"Src":  *s,
		"Dist": *d,
		"Pkg":  *p,
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

// {{ .Dist }} It's the Union type of {{ .Src }} and Array
type {{ .Dist }} []{{ .Src }}

// MarshalJSON returns m as the JSON encoding of m.
func (m {{ .Dist }}) MarshalJSON() ([]byte, error) {
	return json.Marshal([]{{ .Src }}(m))
}

// UnmarshalJSON sets *m to a copy of data.
func (m *{{ .Dist }}) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("spec.{{ .Dist }}: UnmarshalJSON on nil pointer")
	}
	if len(data) == 0 {
		return nil
	}
	if data[0] != '[' {
		var s {{ .Src }}
		err := json.Unmarshal(data, &s)
		if err != nil {
			return err
		}
		*m = []{{ .Src }}{s}
		return nil
	}

	return json.Unmarshal(data, []{{ .Src }}(*m))
}

`
