package spec

// Info The object provides metadata about the API.
// The metadata MAY be used by the clients if needed,
// and MAY be presented in editing or documentation generation tools for convenience.
type Info struct {

	// REQUIRED.
	// The title of the application.
	Title string `json:"title,omitempty"`

	// A short description of the application.
	// CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`

	// A URL to the Terms of Service for the API.
	// MUST be in the format of a URL.
	TermsOfService string `json:"termsOfService,omitempty"`

	// The contact information for the exposed API.
	Contact *Contact `json:"contact,omitempty"`

	// The license information for the exposed API.
	License *License `json:"license,omitempty"`

	// REQUIRED.
	// The version of the OpenAPI document (which is distinct from the OpenAPI Specification version or the API implementation version).
	Version string `json:"version,omitempty"`
}

// Contact information for the exposed API.
type Contact struct {

	// The identifying name of the contact person/organization.
	Name string `json:"name,omitempty"`

	// The URL pointing to the contact information.
	// MUST be in the format of a URL.
	URL string `json:"url,omitempty"`

	// The email address of the contact person/organization.
	// MUST be in the format of an email address.
	Email string `json:"email,omitempty"`
}

// License information for the exposed API.
type License struct {

	// REQUIRED.
	// The license name used for the API.
	Name string `json:"name,omitempty"`

	// A URL to the license used for the API.
	// MUST be in the format of a URL.
	URL string `json:"url,omitempty"`
}
