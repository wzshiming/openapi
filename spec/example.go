package spec

// example
type example struct {

	// Short description for the example.
	Summary string `json:"summary,omitempty"`

	// Long description for the example.
	// CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`

	// Embedded literal example.
	// The value field and externalValue field are mutually exclusive.
	// To represent examples of media types that cannot naturally represented in JSON or YAML, use a string value to contain the example, escaping where necessary.
	Value Any `json:"value"`

	// A URL that points to the literal example.
	// This provides the capability to reference examples that cannot easily be included in JSON or YAML documents.
	// The value field and externalValue field are mutually exclusive.
	ExternalValue string `json:"externalValue,omitempty"`
}
