package spec

// Contact information for the exposed API.
type Contact struct {

	// The identifying name of the contact person/organization.
	Name string `json:"name"`

	// The URL pointing to the contact information.
	// MUST be in the format of a URL.
	URL string `json:"url,omitempty"`

	// The email address of the contact person/organization.
	// MUST be in the format of an email address.
	Email string `json:"email,omitempty"`
}
