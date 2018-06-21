package spec

// Servers
type Servers []Server

// Server An object representing a Server.
type Server struct {

	// REQUIRED.
	// A URL to the target host.
	// This URL supports Server Variables and MAY be relative, to indicate that the host location is relative to the location where the OpenAPI document is being served.
	// Variable substitutions will be made when a variable is named in {brackets}.
	URL string `json:"url,omitempty"`

	// An optional string describing the host designated by the URL.
	// CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`

	// A map between a variable name and its value.
	// The value is used for substitution in the server's URL template.
	Variables map[string]*ServerVariable `json:"variables,omitempty"`
}

// ServerVariable An object representing a Server Variable for server URL template substitution.
type ServerVariable struct {

	// An enumeration of string values to be used if the substitution options are from a limited set.
	Enum []string `json:"enum,omitempty"`

	// REQUIRED.
	// The default value to use for substitution, and to send, if an alternate value is not supplied.
	// Unlike the Schema Object's default, this value MUST be provided by the consumer.
	Default string `json:"default,omitempty"`

	// An optional description for the server variable.
	// CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`
}
