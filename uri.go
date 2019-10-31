package uri

import (
	"fmt"
)

type URI interface {
	URL() string
	String() string
}

// DEPRECATED (20191031/thisisaaronland)

func NewURIWithType(str_uri string, str_type string) (URI, error) {

	uri := fmt.Sprintf("%s://%s", str_type, str_uri)
	return NewURI(uri)
}
