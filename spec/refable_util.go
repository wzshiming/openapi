package spec

// RefSchemas creates a ref schemas
func RefSchemas(ref string) *SchemaOrRefable {
	so := &SchemaOrRefable{}
	so.Ref = "#/components/schemas/" + ref
	return so
}

// RefResponse creates a ref responses
func RefResponse(ref string) *ResponseOrRefable {
	p := &ResponseOrRefable{}
	p.Ref = "#/components/responses/" + ref
	return p
}

// RefParameter creates a ref rarameter
func RefParameter(ref string) *ParameterOrRefable {
	p := &ParameterOrRefable{}
	p.Ref = "#/components/parameters/" + ref
	return p
}

// RefExampley creates a ref example
func RefExample(ref string) *ExampleOrRefable {
	rb := &ExampleOrRefable{}
	rb.Ref = "#/components/examples/" + ref
	return rb
}

// RefRequestBody creates a ref request body
func RefRequestBody(ref string) *RequestBodyOrRefable {
	rb := &RequestBodyOrRefable{}
	rb.Ref = "#/components/requestBodies/" + ref
	return rb
}

// RefHeader creates a ref header
func RefHeader(ref string) *HeaderOrRefable {
	rb := &HeaderOrRefable{}
	rb.Ref = "#/components/headers/" + ref
	return rb
}

// RefSecurityScheme creates a ref security schemes
func RefSecurityScheme(ref string) *SecuritySchemeOrRefable {
	rb := &SecuritySchemeOrRefable{}
	rb.Ref = "#/components/securitySchemes/" + ref
	return rb
}

// RefLink creates a ref link
func RefLink(ref string) *LinkOrRefable {
	rb := &LinkOrRefable{}
	rb.Ref = "#/components/links/" + ref
	return rb
}

// RefCallback creates a ref callback
func RefCallback(ref string) *CallbackOrRefable {
	rb := &CallbackOrRefable{}
	rb.Ref = "#/components/callbacks/" + ref
	return rb
}
