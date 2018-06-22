package spec

// WithDescription sets the description on this response, allows for chaining
func (h *Header) WithDescription(description string) *Header {
	h.Description = description
	return h
}

// WithName a fluent builder method to override the name of the parameter
func (p *Header) WithName(name string) *Header {
	p.Name = name
	return p
}

// AllowsEmptyValues flags this parameter as being ok with empty values
func (p *Header) AllowsEmptyValues() *Header {
	p.AllowEmptyValue = true
	return p
}

// NoEmptyValues flags this parameter as not liking empty values
func (p *Header) NoEmptyValues() *Header {
	p.AllowEmptyValue = false
	return p
}

// AsOptional flags this parameter as optional
func (p *Header) AsOptional() *Header {
	p.Required = false
	return p
}
