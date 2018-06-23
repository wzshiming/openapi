package spec

// MediaType Each Media Type Object provides schema and examples for the media type identified by its key.
type MediaType struct {
	// The schema defining the type used for the request body.
	Schema *Schema `json:"schema,omitempty"`

	// Any Example of the media type.
	// The example object SHOULD be in the correct format as specified by the media type.
	// The example field is mutually exclusive of the examples field.
	// Furthermore, if referencing a schema which contains an example, the example value SHALL override the example provided by the schema.
	Example Any `json:"example,omitempty"`

	// Examples of the media type. Each example object SHOULD match the media type and specified schema if present.
	// The examples field is mutually exclusive of the example field. Furthermore, if referencing a schema which contains an example, the examples value SHALL override the example provided by the schema.
	Examples map[string]*Example `json:"examples,omitempty"`

	// A map between a property name and its encoding information.
	// The key, being the property name, MUST exist in the schema as a property.
	// The encoding object SHALL only apply to requestBody objects when the media type is multipart or application/x-www-form-urlencoded.
	Encoding map[string]*Encoding `json:"encoding,omitempty"`
}
