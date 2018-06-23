package spec

// Refable A simple object to allow referencing other components in the specification, internally and externally.
// The Reference Object is defined by JSON Reference and follows the same structure, behavior and rules.
// For this specification, reference resolution is accomplished as defined by the JSON Reference specification and not by the JSON Schema specification.
type Refable struct {
	// REQUIRED.
	// The reference string.
	Ref string `json:"$ref,omitempty"`
}
