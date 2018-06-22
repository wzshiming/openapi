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
func BooleanProperty() *Schema {
	so := &Schema{}
	so.Type = TypeBoolean
	return so
}

// Float64Property creates a float64/double property
func Float64Property() *Schema {
	so := &Schema{}
	so.Type = TypeNumber
	so.Format = "double"
	return so
}

// Float32Property creates a float32/float property
func Float32Property() *Schema {
	so := &Schema{}
	so.Type = TypeNumber
	so.Format = "float"
	return so
}

// Int8Property creates an int8 property
func Int8Property() *Schema {
	so := &Schema{}
	so.Type = TypeInteger
	so.Format = "int8"
	return so
}

// Int16Property creates an int16 property
func Int16Property() *Schema {
	so := &Schema{}
	so.Type = TypeInteger
	so.Format = "int16"
	return so
}

// Int32Property creates an int32 property
func Int32Property() *Schema {
	so := &Schema{}
	so.Type = TypeInteger
	so.Format = "int32"
	return so
}

// Int64Property creates an int64 property
func Int64Property() *Schema {
	so := &Schema{}
	so.Type = TypeInteger
	so.Format = "int64"
	return so
}

// StringProperty creates a string property
func StringProperty() *Schema {
	return StrFmtProperty("")
}

// ByteProperty creates a string property base64 encoded characters.
func ByteProperty() *Schema {
	return StrFmtProperty("byte")
}

// BinaryProperty creates a string property any sequence of octets.
func BinaryProperty() *Schema {
	return StrFmtProperty("binary")
}

// PasswordProperty creates a string property A hint to UIs to obscure input.
func PasswordProperty() *Schema {
	return StrFmtProperty("password")
}

// DateProperty creates a date property
func DateProperty() *Schema {
	return StrFmtProperty("data")
}

// DateTimeProperty creates a date time property
func DateTimeProperty() *Schema {
	return StrFmtProperty("data-time")
}

// StrFmtProperty creates a property for the named string format
func StrFmtProperty(format string) *Schema {
	so := &Schema{}
	so.Type = TypeString
	so.Format = format
	return so
}

// MapProperty creates a map property
func MapProperty(property *Schema) *Schema {
	so := &Schema{}
	so.Type = TypeObject
	so.AdditionalProperties = property
	return so
}

// ArrayProperty creates an array property
func ArrayProperty(items *Schema) *Schema {
	so := &Schema{}
	so.Type = TypeArray
	so.Items = items
	return so
}

// ComposedSchema creates a schema with allOf
func ComposedSchema(schemas ...Schema) *Schema {
	so := &Schema{}
	so.AllOf = schemas
	return so
}
