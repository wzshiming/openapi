package spec

// Media type definitions are spread across several resources. The media type definitions SHOULD be in compliance with RFC6838.
// Some examples of possible media type definitions:
const (
	MimeJSON        = "application/json"
	MimeXML         = "application/xml"
	MimeTextPlain   = "text/plain"
	MimeOctetStream = "application/octet-stream"
	MimeURLEncoded  = "application/x-www-form-urlencoded"
	MimeFormData    = "multipart/form-data"
)

// JSONRequestBody creates a request body
func JSONRequestBody(schema *Schema) *RequestBody {
	return NewRequestBody(MimeJSON, schema)
}

// XMLRequestBody creates a request body
func XMLRequestBody(schema *Schema) *RequestBody {
	return NewRequestBody(MimeXML, schema)
}

// TextPlainRequestBody creates a request body
func TextPlainRequestBody(schema *Schema) *RequestBody {
	return NewRequestBody(MimeTextPlain, schema)
}

// OctetStreamRequestBody creates a request body
func OctetStreamRequestBody(schema *Schema) *RequestBody {
	return NewRequestBody(MimeOctetStream, schema)
}

// URLEncodedRequestBody creates a request body
func URLEncodedRequestBody(schema *Schema) *RequestBody {
	return NewRequestBody(MimeURLEncoded, schema)
}

// FormDataRequestBody creates a request body
func FormDataRequestBody(schema *Schema) *RequestBody {
	return NewRequestBody(MimeFormData, schema)
}

// NewRequestBody creates a request body
func NewRequestBody(typ string, schema *Schema) *RequestBody {
	rb := &RequestBody{}
	rb.Content = map[string]*MediaType{
		typ: &MediaType{
			Schema: schema,
		},
	}
	return rb
}
