package ui

import (
	"bytes"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/wzshiming/go-bindata/fs"
	"github.com/wzshiming/openapi/spec"
	"github.com/wzshiming/openapi/util"
)

const index = "index.html"

// HandleWithFiles Replace the openapi file path
func HandleWithFiles(m map[string][]byte, asset func(path string) ([]byte, error)) http.Handler {
	return HandleWith(func(path string) ([]byte, error) {
		data, ok := m[path]
		if ok {
			return data, nil
		}
		return nil, os.ErrNotExist
	}, asset)
}

// HandleWithFile Replace the openapi file path
func HandleWithFile(name string, data []byte, asset func(path string) ([]byte, error)) http.Handler {
	return HandleWith(func(path string) ([]byte, error) {
		if path == name {
			return data, nil
		}
		return nil, os.ErrNotExist
	}, asset)
}

// HandleWith Try the contents of openapi from f
func HandleWith(f, asset func(path string) ([]byte, error)) http.Handler {
	fun := f
	if fun == nil {
		fun = asset
	} else {
		fun = func(path string) ([]byte, error) {
			data, err := f(path)
			if err == nil {
				return data, nil
			}
			return asset(path)
		}
	}
	afs := &fs.AssetFS{
		Asset: fun,
		Index: index,
	}
	return http.FileServer(afs)
}

// HandleURL Get the url address from the url parameter and modify the default address
func HandleURL(asset func(path string) ([]byte, error)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		asset0 := func(path string) ([]byte, error) {
			data, err := asset(path)
			if err != nil {
				return nil, err
			}

			if path == index {
				srcs := r.URL.Query()["src"]
				if len(srcs) != 0 {
					val := url.Values{}
					val["src"] = srcs
					data = bytes.Replace(data, []byte("./openapi.json"), []byte("./openapi.json?"+val.Encode()), 1)
				}
			}
			return data, nil
		}
		afs := &fs.AssetFS{
			Asset: asset0,
			Index: index,
		}
		http.FileServer(afs).ServeHTTP(w, r)
	})
}

// HandleOpenapi Get the url address from the url parameter to combine it
func HandleOpenapi(w http.ResponseWriter, r *http.Request) {
	srcs := r.URL.Query()["src"]

	var root *spec.OpenAPI
	for i, src := range srcs {
		var api *spec.OpenAPI
		file, err := util.ReadFile(src)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		err = util.Unmarshal(file, &api)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		if i == 0 {
			root = api
		} else {
			util.Merge(root, api)
		}
	}

	data, err := util.Marshal(root, util.JSON)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	http.ServeContent(w, r, "openapi.json", time.Now(), bytes.NewReader(data))
}
