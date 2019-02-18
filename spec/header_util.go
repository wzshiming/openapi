package spec

// WithDescription sets the description on this response, allows for chaining
func (h *Header) WithDescription(description string) *Header {
	h.Description = description
	return h
}

// AllowsEmptyValues flags this parameter as being ok with empty values
func (h *Header) AllowsEmptyValues() *Header {
	h.AllowEmptyValue = true
	return h
}

// NoEmptyValues flags this parameter as not liking empty values
func (h *Header) NoEmptyValues() *Header {
	h.AllowEmptyValue = false
	return h
}

// AsOptional flags this parameter as optional
func (h *Header) AsOptional() *Header {
	h.Required = false
	return h
}
