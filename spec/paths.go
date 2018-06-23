package spec

// Paths Holds the relative paths to the individual endpoints and their operations.
// The path is appended to the URL from the Server Object in order to construct the full URL.
// The Paths MAY be empty, due to ACL constraints.
type Paths map[string]*PathItem

// PathItem Describes the operations available on a single path.
// A Path Item MAY be empty, due to ACL constraints.
// The path itself is still exposed to the documentation viewer but they will
// not know which operations and parameters are available.
type PathItem struct {

	// An optional, string summary, intended to apply to all operations in this path.
	Summary string `json:"summary,omitempty"`

	// An optional, string description, intended to apply to all operations in this path.
	// CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`

	// A definition of a GET operation on this path.
	Get *Operation `json:"get,omitempty"`

	// A definition of a PUT operation on this path.
	Put *Operation `json:"put,omitempty"`

	// A definition of a POST operation on this path.
	Post *Operation `json:"post,omitempty"`

	// A definition of a DELETE operation on this path.
	Delete *Operation `json:"delete,omitempty"`

	// A definition of a OPTIONS operation on this path.
	Options *Operation `json:"options,omitempty"`

	// A definition of a HEAD operation on this path.
	Head *Operation `json:"head,omitempty"`

	// A definition of a PATCH operation on this path.
	Patch *Operation `json:"patch,omitempty"`

	// A definition of a TRACE operation on this path.
	Trace *Operation `json:"trace,omitempty"`

	// An alternative server array to service all operations in this path.
	Servers []*Server `json:"servers,omitempty"`

	// A list of parameters that are applicable for all the operations described under this path.
	// These parameters can be overridden at the operation level, but cannot be removed there.
	// The list MUST NOT include duplicated parameters.
	// A unique parameter is defined by a combination of a name and location.
	// The list can use the Reference Object to link to parameters that are defined at the OpenAPI Object's components/parameters.
	Parameters []*Parameter `json:"parameters,omitempty"`
}
