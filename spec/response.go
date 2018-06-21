package spec

// Responses A container for the expected responses of an operation.
// The container maps a HTTP response code to the expected response.
// The documentation is not necessarily expected to cover all possible HTTP response codes because they may not be known in advance.
// However, documentation is expected to cover a successful operation response and any known errors.
// The default MAY be used as a default response object for all HTTP codes that are not covered individually by the specification.
// The Responses Object MUST contain at least one response code, and it SHOULD be the response for a successful operation call.
type Responses map[string]ResponseOrRefable

// Response Describes a single response from an API Operation, including design-time, static links to operations based on the response.
type Response struct {

	// REQUIRED.
	// A short description of the response.
	// CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`

	// Maps a header name to its definition.
	// RFC7230 states header names are case insensitive.
	// If a response header is defined with the name "Content-Type", it SHALL be ignored.
	Headers map[string]HeaderOrRefable `json:"headers,omitempty"`

	// A map containing descriptions of potential response payloads.
	// The key is a media type or media type range and the value describes it.
	// For responses that match multiple keys, only the most specific key is applicable.
	// e.g. text/plain overrides text/*
	Content map[string]MediaType `json:"content,omitempty"`

	// A map of operations links that can be followed from the response.
	// The key of the map is a short name for the link, following the naming constraints of the names for Component Objects.
	Links map[string]LinkOrRefable `json:"links,omitempty"`
}
