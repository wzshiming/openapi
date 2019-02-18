package spec

//go:generate go run ./hack/gen_x_or_refable.go -s schema          -d Schema         -f schema_or_refable.go
//go:generate go run ./hack/gen_x_or_refable.go -s response        -d Response       -f response_or_refable.go
//go:generate go run ./hack/gen_x_or_refable.go -s parameter       -d Parameter      -f parameter_or_refable.go
//go:generate go run ./hack/gen_x_or_refable.go -s header          -d Header         -f header_or_refable.go
//go:generate go run ./hack/gen_x_or_refable.go -s example         -d Example        -f example_or_refable.go
//go:generate go run ./hack/gen_x_or_refable.go -s requestBody     -d RequestBody    -f request_body_or_refable.go
//go:generate go run ./hack/gen_x_or_refable.go -s securityScheme  -d SecurityScheme -f security_scheme_or_refable.go
//go:generate go run ./hack/gen_x_or_refable.go -s link            -d Link           -f link_or_refable.go
//go:generate go run ./hack/gen_x_or_refable.go -s callback        -d Callback       -f callback_or_refable.go
