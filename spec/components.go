package spec

// Components Holds a set of reusable objects for different aspects of the OAS.
// All objects defined within the components object will have no effect on the API unless they are explicitly referenced from properties outside the components object.
type Components struct {

	// An object to hold reusable Schema Objects.
	Schemas map[string]SchemaOrRefable `json:"schemas,omitempty"`

	// An object to hold reusable Response Objects.
	Responses map[string]ResponseOrRefable `json:"responses,omitempty"`

	// An object to hold reusable Parameter Objects.
	Parameters map[string]ParameterOrRefable `json:"parameters,omitempty"`

	// An object to hold reusable Example Objects.
	Examples map[string]ExampleOrRefable `json:"examples,omitempty"`

	// An object to hold reusable Request Body Objects.
	RequestBodies map[string]RequestBodyOrRefable `json:"requestBodies,omitempty"`

	// An object to hold reusable Header Objects.
	Headers map[string]HeaderOrRefable `json:"headers,omitempty"`

	// An object to hold reusable Security Scheme Objects.
	SecuritySchemes map[string]SecuritySchemeOrRefable `json:"securitySchemes,omitempty"`

	// An object to hold reusable Link Objects.
	Links map[string]LinkOrRefable `json:"links,omitempty"`

	// An object to hold reusable Callback Objects.
	Callbacks map[string]CallbackOrRefable `json:"callbacks,omitempty"`
}
