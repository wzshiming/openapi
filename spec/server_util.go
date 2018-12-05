package spec

import (
	"net/url"
	"sort"
)

// NewServers Will list all possible server options based on all servers
func NewServers(uris ...string) ([]*Server, error) {
	uris = append(uris, "/")
	uris = removeDuplicates(uris)
	ss := []*Server{}
	for _, v := range uris {
		ss = append(ss, &Server{
			URL: v,
		})
	}

	schemes := []string{"http://", "https://"}
	hosts := []string{"localhost"}
	paths := []string{"/"}
	ports := []string{""}
	for _, v := range uris {
		uri, err := url.Parse(v)
		if err != nil {
			return nil, err
		}

		if host := uri.Hostname(); host != "" {
			hosts = append(hosts, host)
		}

		if uri.Path != "" {
			paths = append(paths, uri.Path)
		}

		if port := uri.Port(); port != "" {
			ports = append(ports, ":"+port)
		}
	}
	hosts = removeDuplicates(hosts)
	ports = removeDuplicates(ports)
	paths = removeDuplicates(paths)
	ss = append(ss, &Server{
		URL: "{scheme}{host}{port}{path}",
		Variables: map[string]*ServerVariable{
			"scheme": &ServerVariable{
				Enum:    schemes,
				Default: schemes[0],
			},
			"host": &ServerVariable{
				Enum:    hosts,
				Default: hosts[0],
			},
			"port": &ServerVariable{
				Enum:    ports,
				Default: ports[0],
			},
			"path": &ServerVariable{
				Enum:    paths,
				Default: paths[0],
			},
		},
	})
	return ss, nil
}

func removeDuplicates(a []string) []string {
	if len(a) <= 1 {
		return a
	}
	sort.Strings(a)
	ret := []string{a[0]}
	for i := 1; i != len(a); i++ {
		if a[i-1] != a[i] {
			ret = append(ret, a[i])
		}
	}
	return ret
}
