package spec

// Parameter Locations
// There are four possible parameter locations specified by the in field:
const (
	InPath   = "path"   // Used together with Path Templating, where the parameter value is actually part of the operation's URL. This does not include the host or base path of the API. For example, in /items/{itemId}, the path parameter is itemId.
	InQuery  = "query"  // Parameters that are appended to the URL. For example, in /items?id=###, the query parameter is id.
	InHeader = "header" // Custom headers that are expected as part of the request. Note that RFC7230 states header names are case insensitive.
	InCookie = "cookie" // Used to pass a specific cookie value to the API.
)

// QueryParam creates a query parameter
func QueryParam(name string, schema *Schema) *Parameter {
	p := &Parameter{}
	p.In = InQuery
	p.Name = name
	p.Schema = schema
	return p
}

// HeaderParam creates a header parameter, this is always required by default
func HeaderParam(name string, schema *Schema) *Parameter {
	p := &Parameter{}
	p.In = InHeader
	p.Name = name
	p.Schema = schema
	return p
}

// PathParam creates a path parameter, this is always required
func PathParam(name string, schema *Schema) *Parameter {
	p := &Parameter{}
	p.In = InPath
	p.Name = name
	p.Schema = schema
	p.Required = true
	return p
}

// CookieParam creates a path parameter, this is always required
func CookieParam(name string, schema *Schema) *Parameter {
	p := &Parameter{}
	p.In = InCookie
	p.Name = name
	p.Schema = schema
	return p
}

// WithDescription sets the description on this response, allows for chaining
func (p *Parameter) WithDescription(description string) *Parameter {
	p.Description = description
	return p
}

// WithName a fluent builder method to override the name of the parameter
func (p *Parameter) WithName(name string) *Parameter {
	p.Name = name
	return p
}

// WithSchema a fluent builder method to override the schema of the parameter
func (p *Parameter) WithSchema(schema *Schema) *Parameter {
	p.Schema = schema
	return p
}

// WithLocation a fluent builder method to override the location of the parameter
func (p *Parameter) WithLocation(in string) *Parameter {
	p.In = in
	return p
}

// AllowsEmptyValues flags this parameter as being ok with empty values
func (p *Parameter) AllowsEmptyValues() *Parameter {
	p.AllowEmptyValue = true
	return p
}

// NoEmptyValues flags this parameter as not liking empty values
func (p *Parameter) NoEmptyValues() *Parameter {
	p.AllowEmptyValue = false
	return p
}

// AsOptional flags this parameter as optional
func (p *Parameter) AsOptional() *Parameter {
	p.Required = false
	return p
}
