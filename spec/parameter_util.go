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
