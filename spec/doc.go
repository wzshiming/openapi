package spec

//go:generate go run ./hack/gen_x_or_refable.go -s Schema          -d SchemaOrRefable         -f schema_or_refable.go
//go:generate go run ./hack/gen_x_or_refable.go -s Response        -d ResponseOrRefable       -f response_or_refable.go
//go:generate go run ./hack/gen_x_or_refable.go -s Parameter       -d ParameterOrRefable      -f parameter_or_refable.go
//go:generate go run ./hack/gen_x_or_refable.go -s Example         -d ExampleOrRefable        -f example_or_refable.go
//go:generate go run ./hack/gen_x_or_refable.go -s RequestBody     -d RequestBodyOrRefable    -f request_body_or_refable.go
//go:generate go run ./hack/gen_x_or_refable.go -s Header          -d HeaderOrRefable         -f header_or_refable.go
//go:generate go run ./hack/gen_x_or_refable.go -s SecurityScheme  -d SecuritySchemeOrRefable -f security_scheme_or_refable.go
//go:generate go run ./hack/gen_x_or_refable.go -s Link            -d LinkOrRefable           -f link_or_refable.go
//go:generate go run ./hack/gen_x_or_refable.go -s Callback        -d CallbackOrRefable       -f callback_or_refable.go
