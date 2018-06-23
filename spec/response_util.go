package spec

// JSONResponse creates a request body
func JSONResponse(schema *Schema) *Response {
	return NewResponse(MimeJSON, schema)
}

// XMLResponse creates a request body
func XMLResponse(schema *Schema) *Response {
	return NewResponse(MimeXML, schema)
}

// TextPlainResponse creates a request body
func TextPlainResponse(schema *Schema) *Response {
	return NewResponse(MimeTextPlain, schema)
}

// OctetStreamResponse creates a request body
func OctetStreamResponse(schema *Schema) *Response {
	return NewResponse(MimeOctetStream, schema)
}

// URLEncodedResponse creates a request body
func URLEncodedResponse(schema *Schema) *Response {
	return NewResponse(MimeURLEncoded, schema)
}

// FormDataResponse creates a request body
func FormDataResponse(schema *Schema) *Response {
	return NewResponse(MimeFormData, schema)
}

// NewResponse creates a response
func NewResponse(typ string, schema *Schema) *Response {
	rb := &Response{}
	rb.Content = map[string]*MediaType{
		typ: &MediaType{
			Schema: schema,
		},
	}
	return rb
}

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
