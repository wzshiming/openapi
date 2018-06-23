package spec

// License information for the exposed API.
type License struct {

	// REQUIRED.
	// The license name used for the API.
	Name string `json:"name"`

	// A URL to the license used for the API.
	// MUST be in the format of a URL.
	URL string `json:"url,omitempty"`
}
