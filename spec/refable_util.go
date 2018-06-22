package spec

// RefSchemas creates a ref schemas
func RefSchemas(ref string) *Schema {
	so := &Schema{}
	so.Ref = "#/components/schemas/" + ref
	return so
}

// RefResponse creates a ref responses
func RefResponse(ref string) *Response {
	p := &Response{}
	p.Ref = "#/components/responses/" + ref
	return p
}

// RefParameter creates a ref rarameter
func RefParameter(ref string) *Parameter {
	p := &Parameter{}
	p.Ref = "#/components/parameters/" + ref
	return p
}

// RefExampley creates a ref example
func RefExample(ref string) *Example {
	rb := &Example{}
	rb.Ref = "#/components/examples/" + ref
	return rb
}

// RefRequestBody creates a ref request body
func RefRequestBody(ref string) *RequestBody {
	rb := &RequestBody{}
	rb.Ref = "#/components/requestBodies/" + ref
	return rb
}

// RefHeader creates a ref header
func RefHeader(ref string) *Header {
	rb := &Header{}
	rb.Ref = "#/components/headers/" + ref
	return rb
}

// RefSecurityScheme creates a ref security schemes
func RefSecurityScheme(ref string) *SecurityScheme {
	rb := &SecurityScheme{}
	rb.Ref = "#/components/securitySchemes/" + ref
	return rb
}

// RefLink creates a ref link
func RefLink(ref string) *Link {
	rb := &Link{}
	rb.Ref = "#/components/links/" + ref
	return rb
}

// RefCallback creates a ref callback
func RefCallback(ref string) *Callback {
	rb := &Callback{}
	rb.Ref = "#/components/callbacks/" + ref
	return rb
}
