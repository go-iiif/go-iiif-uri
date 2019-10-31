package uri

import (
	"net/url"
)

func init() {
	dr := NewStringURIDriver()
	RegisterDriver("string", dr)
	RegisterDriver("file", dr)
}

type StringURIDriver struct {
	Driver
}

func NewStringURIDriver() Driver {

	dr := StringURIDriver{}
	return &dr
}

func (dr *StringURIDriver) NewURI(str_uri string) (URI, error) {

	u, err := url.Parse(str_uri)

	if err != nil {
		return nil, err
	}

	return NewStringURI(u.Path)
}

type StringURI struct {
	URI
	raw string
}

func NewStringURI(raw string) (URI, error) {

	u := StringURI{
		raw: raw,
	}

	return &u, nil
}

func (u *StringURI) URL() string {
	return u.raw
}

func (u *StringURI) String() string {
	return u.raw
}
