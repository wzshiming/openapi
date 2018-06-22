package spec

const (
	InQuery  = "query"
	InHeader = "header"
	InPath   = "path"
	InCookie = "cookie"
)

// QueryParam creates a query parameter
func QueryParam(name string) *Parameter {
	p := &Parameter{}
	p.In = InQuery
	p.Name = name
	return p
}

// HeaderParam creates a header parameter, this is always required by default
func HeaderParam(name string) *Parameter {
	p := &Parameter{}
	p.In = InHeader
	p.Name = name
	return p
}

// PathParam creates a path parameter, this is always required
func PathParam(name string) *Parameter {
	p := &Parameter{}
	p.In = InPath
	p.Name = name
	return p
}

// CookieParam creates a path parameter, this is always required
func CookieParam(name string) *Parameter {
	p := &Parameter{}
	p.In = InCookie
	p.Name = name
	return p
}

// WithDescription sets the description on this response, allows for chaining
func (h *Parameter) WithDescription(description string) *Parameter {
	h.Description = description
	return h
}

// WithName a fluent builder method to override the name of the parameter
func (p *Parameter) WithName(name string) *Parameter {
	p.Name = name
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
