package uri

import (
)

type URI interface {
	String() string
	Base() string
	Root() string
}

func NewURI(str_uri string) (URI, error) {

	return NewURIWithDriver(str_uri)
}
