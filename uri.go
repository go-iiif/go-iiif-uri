package uri

import (
	"fmt"
	_ "log"
	"path/filepath"
)

type URI interface {
	String() string
	// URL() string
	Base() string
	Root() string
}

func Path(u URI) string {
	return filepath.Join(u.Root(), u.Base())
}

// DEPRECATED (20191031/thisisaaronland)

func NewURIWithType(str_uri string, str_type string) (URI, error) {
	uri := fmt.Sprintf("%s://%s", str_type, str_uri)
	return NewURI(uri)
}
