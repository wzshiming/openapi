package spec

const (
	MimeJSON        = "application/json"
	MimeXML         = "application/xml"
	MimeTextPlain   = "text/plain"
	MimeOctetStream = "application/octet-stream"
	MimeURLEncoded  = "application/x-www-form-urlencoded"
	MimeFormData    = "multipart/form-data"
)

// JSONRequestBody creates a request body
func JSONRequestBody(schema *SchemaOrRefable) *RequestBodyOrRefable {
	return NewRequestBody(MimeJSON, schema)
}

// XMLRequestBody creates a request body
func XMLRequestBody(schema *SchemaOrRefable) *RequestBodyOrRefable {
	return NewRequestBody(MimeXML, schema)
}

// TextPlainRequestBody creates a request body
func TextPlainRequestBody(schema *SchemaOrRefable) *RequestBodyOrRefable {
	return NewRequestBody(MimeTextPlain, schema)
}

// OctetStreamRequestBody creates a request body
func OctetStreamRequestBody(schema *SchemaOrRefable) *RequestBodyOrRefable {
	return NewRequestBody(MimeOctetStream, schema)
}

// URLEncodedRequestBody creates a request body
func URLEncodedRequestBody(schema *SchemaOrRefable) *RequestBodyOrRefable {
	return NewRequestBody(MimeURLEncoded, schema)
}

// FormDataRequestBody creates a request body
func FormDataRequestBody(schema *SchemaOrRefable) *RequestBodyOrRefable {
	return NewRequestBody(MimeFormData, schema)
}

// NewRequestBody creates a request body
func NewRequestBody(typ string, schema *SchemaOrRefable) *RequestBodyOrRefable {
	rb := &RequestBodyOrRefable{}
	rb.Content = map[string]MediaType{
		typ: MediaType{
			Schema: schema,
		},
	}
	return rb
}
