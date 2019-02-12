package util

import (
	"github.com/wzshiming/openapi/spec"
)

// Merge Added API to root
func Merge(root, added *spec.OpenAPI) {
	for name, path := range added.Paths {
		rpath, ok := root.Paths[name]
		if !ok {
			root.Paths[name] = path
			continue
		}

		if path.Get != nil && rpath.Get == nil {
			rpath.Get = path.Get
		}

		if path.Post != nil && rpath.Post == nil {
			rpath.Post = path.Post
		}

		if path.Put != nil && rpath.Put == nil {
			rpath.Put = path.Put
		}

		if path.Delete != nil && rpath.Delete == nil {
			rpath.Delete = path.Delete
		}

		if path.Head != nil && rpath.Head == nil {
			rpath.Head = path.Head
		}

		if path.Options != nil && rpath.Options == nil {
			rpath.Options = path.Options
		}

		if path.Patch != nil && rpath.Patch == nil {
			rpath.Patch = path.Patch
		}

		if path.Trace != nil && rpath.Trace == nil {
			rpath.Trace = path.Trace
		}
	}

	root.Tags = append(root.Tags, added.Tags...)

	if added.Components != nil && root.Components == nil {
		root.Components = added.Components
	} else {
		rcom := root.Components
		com := added.Components
		for name, v := range com.Schemas {
			if _, ok := rcom.Schemas[name]; !ok {
				rcom.Schemas[name] = v
			}
		}

		for name, v := range com.Responses {
			if _, ok := rcom.Responses[name]; !ok {
				rcom.Responses[name] = v
			}
		}

		for name, v := range com.Parameters {
			if _, ok := rcom.Parameters[name]; !ok {
				rcom.Parameters[name] = v
			}
		}

		for name, v := range com.Examples {
			if _, ok := rcom.Examples[name]; !ok {
				rcom.Examples[name] = v
			}
		}

		for name, v := range com.RequestBodies {
			if _, ok := rcom.RequestBodies[name]; !ok {
				rcom.RequestBodies[name] = v
			}
		}

		for name, v := range com.Headers {
			if _, ok := rcom.Headers[name]; !ok {
				rcom.Headers[name] = v
			}
		}

		for name, v := range com.SecuritySchemes {
			if _, ok := rcom.SecuritySchemes[name]; !ok {
				rcom.SecuritySchemes[name] = v
			}
		}

		for name, v := range com.Links {
			if _, ok := rcom.Links[name]; !ok {
				rcom.Links[name] = v
			}
		}

		for name, v := range com.Callbacks {
			if _, ok := rcom.Callbacks[name]; !ok {
				rcom.Callbacks[name] = v
			}
		}
	}
}
