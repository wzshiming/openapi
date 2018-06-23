package spec

// WithDescription sets the description on this response, allows for chaining
func (r *Response) WithDescription(description string) *Response {
	r.Description = description
	return r
}

// AddHeader adds a header to this response
func (r *Response) AddHeader(name string, header *Header) *Response {
	if header == nil {
		return r.RemoveHeader(name)
	}
	if r.Headers == nil {
		r.Headers = map[string]*Header{}
	}
	r.Headers[name] = header
	return r
}

// RemoveHeader removes a header from this response
func (r *Response) RemoveHeader(name string) *Response {
	delete(r.Headers, name)
	return r
}

// AddContent adds a content to this response
func (r *Response) AddContent(name string, content *MediaType) *Response {
	if content == nil {
		return r.RemoveHeader(name)
	}
	if r.Content == nil {
		r.Content = map[string]*MediaType{}
	}
	r.Content[name] = content
	return r
}

// RemoveContent removes a content from this response
func (r *Response) RemoveContent(name string) *Response {
	delete(r.Content, name)
	return r
}

// AddLink adds a link to this response
func (r *Response) AddLink(name string, link *Link) *Response {
	if link == nil {
		return r.RemoveHeader(name)
	}
	if r.Links == nil {
		r.Links = map[string]*Link{}
	}
	r.Links[name] = link
	return r
}

// RemoveLink removes a link from this response
func (r *Response) RemoveLink(name string) *Response {
	delete(r.Links, name)
	return r
}
