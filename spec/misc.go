package spec

type StringOrArray []string

type SchemaOrArray []Schema

type SchemaOrBool struct {
	Allows bool
	Schema *Schema
}

type RequestBodyOrString struct {
	Literal     string
	RequestBody *RequestBody
}

type Headers map[string]Header
type Header Parameter

type Callbacks map[string]Callback
type Callback PathItem
