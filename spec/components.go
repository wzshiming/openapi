package spec

// Components Holds a set of reusable objects for different aspects of the OAS.
// All objects defined within the components object will have no effect on the API unless they are explicitly referenced from properties outside the components object.
type Components struct {

	// An object to hold reusable Schema Objects.
	Schemas map[string]*Schema `json:"schemas,omitempty"`

	// An object to hold reusable Response Objects.
	Responses map[string]*Response `json:"responses,omitempty"`

	// An object to hold reusable Parameter Objects.
	Parameters map[string]*Parameter `json:"parameters,omitempty"`

	// An object to hold reusable Example Objects.
	Examples map[string]*Example `json:"examples,omitempty"`

	// An object to hold reusable Request Body Objects.
	RequestBodies map[string]*RequestBody `json:"requestBodies,omitempty"`

	// An object to hold reusable Header Objects.
	Headers map[string]*Header `json:"headers,omitempty"`

	// An object to hold reusable Security Scheme Objects.
	SecuritySchemes map[string]*SecurityScheme `json:"securitySchemes,omitempty"`

	// An object to hold reusable Link Objects.
	Links map[string]*Link `json:"links,omitempty"`

	// An object to hold reusable Callback Objects.
	Callbacks map[string]*Callback `json:"callbacks,omitempty"`
}

// NewComponents create a new new components
func NewComponents() *Components {
	return &Components{
		Schemas:         map[string]*Schema{},
		Responses:       map[string]*Response{},
		Parameters:      map[string]*Parameter{},
		Examples:        map[string]*Example{},
		RequestBodies:   map[string]*RequestBody{},
		Headers:         map[string]*Header{},
		SecuritySchemes: map[string]*SecurityScheme{},
		Links:           map[string]*Link{},
		Callbacks:       map[string]*Callback{},
	}
}
