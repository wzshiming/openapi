package spec

// ExternalDocumentation Allows referencing an external resource for extended documentation.
type ExternalDocumentation struct {

	// A short description of the target documentation.
	// CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description"`

	// REQUIRED.
	// The URL for the target documentation.
	// Value MUST be in the format of a URL.
	URL string `json:"url,omitempty"`
}
