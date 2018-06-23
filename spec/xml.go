package spec

// XML A metadata object that allows for more fine-tuned XML model definitions.
// When using arrays, XML element names are not inferred (for singular/plural forms) and the name property SHOULD be used to add that information.
// See examples for expected behavior.
type XML struct {

	// Replaces the name of the element/attribute used for the described schema property.
	// When defined within items, it will affect the name of the individual XML elements within the list.
	// When defined alongside type being array (outside the items), it will affect the wrapping element and only if wrapped is true.
	// If wrapped is false, it will be ignored.
	Name string `json:"name"`

	// The URI of the namespace definition.
	// Value MUST be in the form of an absolute URI.
	Namespace string `json:"namespace,omitempty"`

	// The prefix to be used for the name.
	Prefix string `json:"prefix,omitempty"`

	// Declares whether the property definition translates to an attribute instead of an element.
	// Default value is false.
	Attribute bool `json:"attribute,omitempty"`

	// MAY be used only for an array definition.
	// Signifies whether the array is wrapped (for example, <books><book/><book/></books>) or unwrapped (<book/><book/>).
	// Default value is false.
	// The definition takes effect only when defined alongside type being array (outside the items).
	Wrapped bool `json:"wrapped,omitempty"`
}

// WithName sets the xml name for the object
func (x *XML) WithName(name string) *XML {
	x.Name = name
	return x
}

// WithNamespace sets the xml namespace for the object
func (x *XML) WithNamespace(namespace string) *XML {
	x.Namespace = namespace
	return x
}

// WithPrefix sets the xml prefix for the object
func (x *XML) WithPrefix(prefix string) *XML {
	x.Prefix = prefix
	return x
}

// AsAttribute flags this object as xml attribute
func (x *XML) AsAttribute() *XML {
	x.Attribute = true
	return x
}

// AsElement flags this object as an xml node
func (x *XML) AsElement() *XML {
	x.Attribute = false
	return x
}

// AsWrapped flags this object as wrapped, this is mostly useful for array types
func (x *XML) AsWrapped() *XML {
	x.Wrapped = true
	return x
}

// AsUnwrapped flags this object as an xml node
func (x *XML) AsUnwrapped() *XML {
	x.Wrapped = false
	return x
}
