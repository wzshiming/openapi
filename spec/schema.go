package spec

// schema The Schema Object allows the definition of input and output data types.
// These types can be objects, but also primitives and arrays.
// This object is an extended subset of the JSON Schema Specification Wright Draft 00.
// For more information about the properties, see JSON Schema Core and JSON Schema Validation.
// Unless stated otherwise, the property definitions follow the JSON Schema.
type schema struct {

	// Allows sending a null value for the defined schema.
	// Default value is false.
	Nullable bool `json:"nullable,omitempty"`

	// Adds support for polymorphism.
	// The discriminator is an object name that is used to differentiate between other schemas which may satisfy the payload description.
	// See Composition and Inheritance for more details.
	Discriminator *Discriminator `json:"discriminator,omitempty"`

	// Relevant only for Schema "properties" definitions.
	// Declares the property as "read only".
	// This means that it MAY be sent as part of a response but SHOULD NOT be sent as part of the request.
	// If the property is marked as readOnly being true and is in the required list, the required will take effect on the response only.
	// A property MUST NOT be marked as both readOnly and writeOnly being true.
	// Default value is false.
	ReadOnly bool `json:"readOnly,omitempty"`

	// Relevant only for Schema "properties" definitions.
	// Declares the property as "write only".
	// Therefore, it MAY be sent as part of a request but SHOULD NOT be sent as part of the response.
	// If the property is marked as writeOnly being true and is in the required list, the required will take effect on the request only.
	// A property MUST NOT be marked as both readOnly and writeOnly being true.
	// Default value is false.
	WriteOnly bool `json:"writeOnly,omitempty"`

	// This MAY be used only on properties schemas.
	// It has no effect on root schemas.
	// Adds additional metadata to describe the XML representation of this property.
	XML *XML `json:"xml,omitempty"`

	// Additional external documentation for this schema.
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty"`

	// Any	A free-form property to include an example of an instance for this schema.
	// To represent examples that cannot be naturally represented in JSON or YAML, a string value can be used to contain the example with escaping where necessary.
	Example Any `json:"example,omitempty"`

	// Specifies that a schema is deprecated and SHOULD be transitioned out of usage.
	// Default value is false.
	Deprecated bool `json:"deprecated,omitempty"`

	// The following properties are taken directly from the JSON Schema definition and follow the same specifications:

	Title string `json:"title,omitempty"`

	// Numbers
	MultipleOf       *float64 `json:"multipleOf,omitempty"`
	Maximum          *float64 `json:"maximum,omitempty"`
	ExclusiveMaximum bool     `json:"exclusiveMaximum,omitempty"`
	Minimum          *float64 `json:"minimum,omitempty"`
	ExclusiveMinimum bool     `json:"exclusiveMinimum,omitempty"`

	// Strings
	MaxLength *int64 `json:"maxLength,omitempty"`
	MinLength *int64 `json:"minLength,omitempty"`
	Pattern   string `json:"pattern,omitempty"`

	// Arrays
	MaxItems    *int64 `json:"maxItems,omitempty"`
	MinItems    *int64 `json:"minItems,omitempty"`
	UniqueItems bool   `json:"uniqueItems,omitempty"`

	// Objects
	MaxProperties *int64   `json:"maxProperties,omitempty"`
	MinProperties *int64   `json:"minProperties,omitempty"`
	Required      []string `json:"required,omitempty"`

	// All
	Enum []Any `json:"enum,omitempty"`

	// The following properties are taken from the JSON Schema definition but their definitions were adjusted to the OpenAPI Specification.

	// Value MUST be a string.
	// Multiple types via an array are not supported.
	Type string `json:"type,omitempty"`

	// Inline or referenced schema MUST be of a Schema Object and not a standard JSON Schema.
	AllOf []*Schema `json:"allOf,omitempty"`

	// Inline or referenced schema MUST be of a Schema Object and not a standard JSON Schema.
	OneOf []*Schema `json:"oneOf,omitempty"`

	// Inline or referenced schema MUST be of a Schema Object and not a standard JSON Schema.
	AnyOf []*Schema `json:"anyOf,omitempty"`

	// Inline or referenced schema MUST be of a Schema Object and not a standard JSON Schema.
	Not *Schema `json:"not,omitempty"`

	// Value MUST be an object and not an array.
	// Inline or referenced schema MUST be of a Schema Object and not a standard JSON Schema.
	// items MUST be present if the type is array.
	Items *Schema `json:"items,omitempty"`

	// Property definitions MUST be a Schema Object and not a standard JSON Schema (inline or referenced).
	Properties map[string]*Schema `json:"properties,omitempty"`

	// Value can be boolean or object.
	// Inline or referenced schema MUST be of a Schema Object and not a standard JSON Schema.
	AdditionalProperties *Schema `json:"additionalProperties,omitempty"`

	// CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`

	// See Data Type Formats for further details.
	// While relying on JSON Schema's defined formats, the OAS offers a few additional predefined formats.
	Format string `json:"format,omitempty"`

	// The default value represents what would be assumed by the consumer of the input as the value of the schema if one is not provided.
	// Unlike JSON Schema, the value MUST conform to the defined type for the Schema Object defined at the same level.
	// For example, if type is string, then default can be "foo" but cannot be 1.
	Default Any `json:"default,omitempty"`
}
