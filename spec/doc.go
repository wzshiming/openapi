package spec

//go:generate go run ./hack/gen_x_or_refable.go -x Schema -f schema_or_refable.go
//go:generate go run ./hack/gen_x_or_refable.go -x Response -f response_or_refable.go
//go:generate go run ./hack/gen_x_or_refable.go -x Parameter -f parameter_or_refable.go
//go:generate go run ./hack/gen_x_or_refable.go -x Example -f example_or_refable.go
//go:generate go run ./hack/gen_x_or_refable.go -x RequestBody -f request_body_or_refable.go
//go:generate go run ./hack/gen_x_or_refable.go -x Header -f header_or_refable.go
//go:generate go run ./hack/gen_x_or_refable.go -x SecurityScheme -f security_scheme_or_refable.go
//go:generate go run ./hack/gen_x_or_refable.go -x Link -f link_or_refable.go
//go:generate go run ./hack/gen_x_or_refable.go -x Callback -f callback_or_refable.go
