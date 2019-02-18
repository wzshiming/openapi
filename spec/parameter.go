package spec

// parameter Describes a single operation parameter.
// A unique parameter is defined by a combination of a name and location.
// Parameter Locations
// There are four possible parameter locations specified by the in field:
// path - Used together with Path Templating, where the parameter value is actually part of the operation's URL.
// This does not include the host or base path of the API.
// For example, in /items/{itemId}, the path parameter is itemId.
// query - Parameters that are appended to the URL.
// For example, in /items?id=###, the query parameter is id.
// header - Custom headers that are expected as part of the request.
// Note that RFC7230 states header names are case insensitive.
// cookie - Used to pass a specific cookie value to the API.
type parameter struct {

	// REQUIRED.
	// The name of the parameter.
	// Parameter names are case sensitive.
	// If in is "path", the name field MUST correspond to the associated path segment from the path field in the Paths Object.
	// See Path Templating for further information.
	// If in is "header" and the name field is "Accept", "Content-Type" or "Authorization", the parameter definition SHALL be ignored.
	// For all other cases, the name corresponds to the parameter name used by the in property.
	Name string `json:"name"`

	// REQUIRED.
	// The location of the parameter.
	// Possible values are "query", "header", "path" or "cookie".
	In string `json:"in"`

	header
}
