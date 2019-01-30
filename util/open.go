package util

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

// ReadFile Automatic distinction protocol switching between HTTP and local
func ReadFile(path string) ([]byte, error) {
	u, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	switch u.Scheme {
	case "http", "https":
		resp, err := http.Get(u.String())
		if err != nil {
			return nil, err
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		return data, nil

	case "file", "":
		data, err := ioutil.ReadFile(u.Path)
		if err != nil {
			return nil, err
		}
		return data, nil
	}

	return nil, errors.New("error path " + path)
}
