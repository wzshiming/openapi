package spec

// RefSchemas creates a ref schemas
func RefSchemas(ref string) *Schema {
	s := &Schema{}
	s.Components = "#/components/schemas/"
	s.Ref = ref
	return s
}

// RefResponse creates a ref responses
func RefResponse(ref string) *Response {
	r := &Response{}
	r.Components = "#/components/responses/"
	r.Ref = ref
	return r
}

// RefParameter creates a ref rarameter
func RefParameter(ref string) *Parameter {
	p := &Parameter{}
	p.Components = "#/components/parameters/"
	p.Ref = ref
	return p
}

// RefExample creates a ref example
func RefExample(ref string) *Example {
	e := &Example{}
	e.Components = "#/components/examples/"
	e.Ref = ref
	return e
}

// RefRequestBody creates a ref request body
func RefRequestBody(ref string) *RequestBody {
	r := &RequestBody{}
	r.Components = "#/components/requestBodies/"
	r.Ref = ref
	return r
}

// RefHeader creates a ref header
func RefHeader(ref string) *Header {
	h := &Header{}
	h.Components = "#/components/headers/"
	h.Ref = ref
	return h
}

// RefSecurityScheme creates a ref security schemes
func RefSecurityScheme(ref string) *SecurityScheme {
	s := &SecurityScheme{}
	s.Components = "#/components/securitySchemes/"
	s.Ref = ref
	return s
}

// RefLink creates a ref link
func RefLink(ref string) *Link {
	l := &Link{}
	l.Components = "#/components/links/"
	l.Ref = ref
	return l
}

// RefCallback creates a ref callback
func RefCallback(ref string) *Callback {
	c := &Callback{}
	c.Components = "#/components/callbacks/"
	c.Ref = ref
	return c
}
