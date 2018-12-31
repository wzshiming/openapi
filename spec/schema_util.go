package spec

// Primitive data types in the OAS are based on the types supported by the JSON Schema Specification Wright Draft 00.
// Note that integer as a type is also supported and is defined as a JSON number without a fraction or exponent part.
// null is not supported as a type (see nullable for an alternative solution).
// Models are defined using the Schema Object, which is an extended subset of JSON Schema Specification Wright Draft 00.
const (
	TypeBoolean = "boolean"
	TypeNumber  = "number"
	TypeInteger = "integer"
	TypeString  = "string"
	TypeObject  = "object"
	TypeArray   = "array"
)

var timeExample = NewAny("2006-01-02T15:04:05+08:00")

// BooleanProperty creates a boolean property
func BooleanProperty() *Schema {
	so := &Schema{}
	so.Type = TypeBoolean
	return so
}

// Float64Property creates a float64/double property
func Float64Property() *Schema {
	return FloatFmtProperty("double")
}

// Float32Property creates a float32/float property
func Float32Property() *Schema {
	return FloatFmtProperty("float")
}

// FloatFmtProperty creates a property for the named float format
func FloatFmtProperty(format string) *Schema {
	so := &Schema{}
	so.Type = TypeNumber
	so.Format = format
	return so
}

// Int8Property creates an int8 property
func Int8Property() *Schema {
	return IntFmtProperty("int8")
}

// Int16Property creates an int16 property
func Int16Property() *Schema {
	return IntFmtProperty("int16")
}

// Int32Property creates an int32 property
func Int32Property() *Schema {
	return IntFmtProperty("int32")
}

// Int64Property creates an int64 property
func Int64Property() *Schema {
	return IntFmtProperty("int64")
}

// IntFmtProperty creates a property for the named int format
func IntFmtProperty(format string) *Schema {
	so := &Schema{}
	so.Type = TypeInteger
	so.Format = format
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
	return StrFmtProperty("data").WithExample(timeExample)
}

// DateTimeProperty creates a date time property
func DateTimeProperty() *Schema {
	return StrFmtProperty("data-time").WithExample(timeExample)
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

// AllSchema creates a schema with allOf
func AllSchema(schemas ...*Schema) *Schema {
	so := &Schema{}
	so.AllOf = schemas
	return so
}

// OneSchema creates a schema with oneOf
func OneSchema(schemas ...*Schema) *Schema {
	so := &Schema{}
	so.OneOf = schemas
	return so
}

// AnySchema creates a schema with anyOf
func AnySchema(schemas ...*Schema) *Schema {
	so := &Schema{}
	so.AnyOf = schemas
	return so
}

// NotSchema creates a schema with not
func NotSchema(schema *Schema) *Schema {
	so := &Schema{}
	so.Not = schema
	return so
}

// WithTitle sets the title for this schema, allows for chaining
func (s *Schema) WithTitle(title string) *Schema {
	s.Title = title
	return s
}

// WithDescription sets the description for this schema, allows for chaining
func (s *Schema) WithDescription(description string) *Schema {
	s.Description = description
	return s
}

// WithProperties sets the properties for this schema
func (s *Schema) WithProperties(schemas map[string]*Schema) *Schema {
	s.Properties = schemas
	return s
}

// WithProperty sets a property on this schema
func (s *Schema) WithProperty(name string, schema *Schema) *Schema {
	if s.Properties == nil {
		s.Properties = make(map[string]*Schema)
	}
	s.Properties[name] = schema
	return s
}

// WithAllOf sets the all of property
func (s *Schema) WithAllOf(schemas ...*Schema) *Schema {
	s.AllOf = schemas
	return s
}

// WithMaxProperties sets the max number of properties an object can have
func (s *Schema) WithMaxProperties(max int64) *Schema {
	s.MaxProperties = &max
	return s
}

// WithMinProperties sets the min number of properties an object must have
func (s *Schema) WithMinProperties(min int64) *Schema {
	s.MinProperties = &min
	return s
}

// WithType sets the type of this schema for a single value item
func (s *Schema) WithType(typ, format string) *Schema {
	s.Type = typ
	s.Format = format
	return s
}

// WithFormat sets the format of this schema for a single value item
func (s *Schema) WithFormat(format string) *Schema {
	s.Format = format
	return s
}

// WithDefault sets the default value on this parameter
func (s *Schema) WithDefault(defaultValue Any) *Schema {
	s.Default = defaultValue
	return s
}

// WithRequired flags this parameter as required
func (s *Schema) WithRequired(items ...string) *Schema {
	s.Required = items
	return s
}

// AddRequired  adds field names to the required properties array
func (s *Schema) AddRequired(items ...string) *Schema {
	s.Required = append(s.Required, items...)
	return s
}

// WithMaxLength sets a max length value
func (s *Schema) WithMaxLength(max int64) *Schema {
	s.MaxLength = &max
	return s
}

// WithMinLength sets a min length value
func (s *Schema) WithMinLength(min int64) *Schema {
	s.MinLength = &min
	return s
}

// WithPattern sets a pattern value
func (s *Schema) WithPattern(pattern string) *Schema {
	s.Pattern = pattern
	return s
}

// WithMultipleOf sets a multiple of value
func (s *Schema) WithMultipleOf(number float64) *Schema {
	s.MultipleOf = &number
	return s
}

// WithMaximum sets a maximum number value
func (s *Schema) WithMaximum(max float64, exclusive bool) *Schema {
	s.Maximum = &max
	s.ExclusiveMaximum = exclusive
	return s
}

// WithMinimum sets a minimum number value
func (s *Schema) WithMinimum(min float64, exclusive bool) *Schema {
	s.Minimum = &min
	s.ExclusiveMinimum = exclusive
	return s
}

// WithEnum sets a the enum values (replace)
func (s *Schema) WithEnum(values ...Any) *Schema {
	s.Enum = values
	return s
}

// WithMaxItems sets the max items
func (s *Schema) WithMaxItems(size int64) *Schema {
	s.MaxItems = &size
	return s
}

// WithMinItems sets the min items
func (s *Schema) WithMinItems(size int64) *Schema {
	s.MinItems = &size
	return s
}

// UniqueValues dictates that this array can only have unique items
func (s *Schema) UniqueValues() *Schema {
	s.UniqueItems = true
	return s
}

// AllowDuplicates this array can have duplicates
func (s *Schema) AllowDuplicates() *Schema {
	s.UniqueItems = false
	return s
}

// AddToAllOf adds a schema to the allOf property
func (s *Schema) AddToAllOf(schemas ...*Schema) *Schema {
	s.AllOf = append(s.AllOf, schemas...)
	return s
}

// AsReadOnly flags this schema as readonly
func (s *Schema) AsReadOnly() *Schema {
	s.ReadOnly = true
	return s
}

// AsWritable flags this schema as writeable (not read-only)
func (s *Schema) AsWritable() *Schema {
	s.ReadOnly = false
	return s
}

// WithExample sets the example for this schema
func (s *Schema) WithExample(example Any) *Schema {
	s.Example = example
	return s
}

// WithExternalDocs sets/removes the external docs for/from this schema.
// When you pass empty strings as params the external documents will be removed.
// When you pass non-empty string as one value then those values will be used on the external docs object.
// So when you pass a non-empty description, you should also pass the url and vice versa.
func (s *Schema) WithExternalDocs(description, url string) *Schema {
	if description == "" && url == "" {
		s.ExternalDocs = nil
		return s
	}

	if s.ExternalDocs == nil {
		s.ExternalDocs = &ExternalDocumentation{}
	}
	s.ExternalDocs.Description = description
	s.ExternalDocs.URL = url
	return s
}

// WithXMLName sets the xml name for the object
func (s *Schema) WithXMLName(name string) *Schema {
	if s.XML == nil {
		s.XML = &XML{}
	}
	s.XML.Name = name
	return s
}

// WithXMLNamespace sets the xml namespace for the object
func (s *Schema) WithXMLNamespace(namespace string) *Schema {
	if s.XML == nil {
		s.XML = &XML{}
	}
	s.XML.Namespace = namespace
	return s
}

// WithXMLPrefix sets the xml prefix for the object
func (s *Schema) WithXMLPrefix(prefix string) *Schema {
	if s.XML == nil {
		s.XML = &XML{}
	}
	s.XML.Prefix = prefix
	return s
}

// AsXMLAttribute flags this object as xml attribute
func (s *Schema) AsXMLAttribute() *Schema {
	if s.XML == nil {
		s.XML = &XML{}
	}
	s.XML.Attribute = true
	return s
}

// AsXMLElement flags this object as an xml node
func (s *Schema) AsXMLElement() *Schema {
	if s.XML == nil {
		s.XML = &XML{}
	}
	s.XML.Attribute = false
	return s
}

// AsWrappedXML flags this object as wrapped, this is mostly useful for array types
func (s *Schema) AsWrappedXML() *Schema {
	if s.XML == nil {
		s.XML = &XML{}
	}
	s.XML.Wrapped = true
	return s
}

// AsUnwrappedXML flags this object as an xml node
func (s *Schema) AsUnwrappedXML() *Schema {
	if s.XML == nil {
		s.XML = &XML{}
	}
	s.XML.Wrapped = false
	return s
}
