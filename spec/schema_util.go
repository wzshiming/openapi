package spec

const (
	TypeBoolean = "boolean"
	TypeNumber  = "number"
	TypeInteger = "integer"
	TypeString  = "string"
	TypeObject  = "object"
	TypeArray   = "array"
)

// BooleanProperty creates a boolean property
func BooleanProperty() *SchemaOrRefable {
	so := &SchemaOrRefable{}
	so.Type = TypeBoolean
	return so
}

// Float64Property creates a float64/double property
func Float64Property() *SchemaOrRefable {
	so := &SchemaOrRefable{}
	so.Type = TypeNumber
	so.Format = "double"
	return so
}

// Float32Property creates a float32/float property
func Float32Property() *SchemaOrRefable {
	so := &SchemaOrRefable{}
	so.Type = TypeNumber
	so.Format = "float"
	return so
}

// Int8Property creates an int8 property
func Int8Property() *SchemaOrRefable {
	so := &SchemaOrRefable{}
	so.Type = TypeInteger
	so.Format = "int8"
	return so
}

// Int16Property creates an int16 property
func Int16Property() *SchemaOrRefable {
	so := &SchemaOrRefable{}
	so.Type = TypeInteger
	so.Format = "int16"
	return so
}

// Int32Property creates an int32 property
func Int32Property() *SchemaOrRefable {
	so := &SchemaOrRefable{}
	so.Type = TypeInteger
	so.Format = "int32"
	return so
}

// Int64Property creates an int64 property
func Int64Property() *SchemaOrRefable {
	so := &SchemaOrRefable{}
	so.Type = TypeInteger
	so.Format = "int64"
	return so
}

// StringProperty creates a string property
func StringProperty() *SchemaOrRefable {
	return StrFmtProperty("")
}

// ByteProperty creates a string property base64 encoded characters.
func ByteProperty() *SchemaOrRefable {
	return StrFmtProperty("byte")
}

// BinaryProperty creates a string property any sequence of octets.
func BinaryProperty() *SchemaOrRefable {
	return StrFmtProperty("binary")
}

// PasswordProperty creates a string property A hint to UIs to obscure input.
func PasswordProperty() *SchemaOrRefable {
	return StrFmtProperty("password")
}

// DateProperty creates a date property
func DateProperty() *SchemaOrRefable {
	return StrFmtProperty("data")
}

// DateTimeProperty creates a date time property
func DateTimeProperty() *SchemaOrRefable {
	return StrFmtProperty("data-time")
}

// StrFmtProperty creates a property for the named string format
func StrFmtProperty(format string) *SchemaOrRefable {
	so := &SchemaOrRefable{}
	so.Type = TypeString
	so.Format = format
	return so
}

// MapProperty creates a map property
func MapProperty(property *SchemaOrRefable) *SchemaOrRefable {
	so := &SchemaOrRefable{}
	so.Type = TypeObject
	so.AdditionalProperties = property
	return so
}

// ArrayProperty creates an array property
func ArrayProperty(items *SchemaOrRefable) *SchemaOrRefable {
	so := &SchemaOrRefable{}
	so.Type = TypeArray
	so.Items = items
	return so
}

// ComposedSchema creates a schema with allOf
func ComposedSchema(schemas ...SchemaOrRefable) *SchemaOrRefable {
	so := &SchemaOrRefable{}
	so.AllOf = schemas
	return so
}
