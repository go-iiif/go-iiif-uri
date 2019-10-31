package uri

import (
	"errors"
	"fmt"
	"github.com/aaronland/go-string/dsn"
	"github.com/aaronland/go-string/random"
	_ "log"
	"net/url"
)

const IdSecretDriverName string = "idsecret"

func init() {
	dr := NewIdSecretURIDriver()
	RegisterDriver(IdSecretDriverName, dr)
}

type IdSecretURIDriver struct {
	Driver
}

func NewIdSecretURIDriver() Driver {

	dr := IdSecretURIDriver{}
	return &dr
}

// idsecret://{SOURCE}?id={ID}&secret={SECRET}

func (dr *IdSecretURIDriver) NewURI(str_uri string) (URI, error) {

	return NewIdSecretURI(str_uri)
}

type IdSecretURI struct {
	URI
	source   string
	id       string
	secret   string
	secret_o string
}

func NewIdSecretURIFromDSN(dsn_raw string) (URI, error) {

	dsn_map, err := dsn.StringToDSNWithKeys(dsn_raw, "id", "uri")

	if err != nil {
		return nil, err
	}

	source := dsn_map["id"]
	id := dsn_map["uri"]

	q := url.Values{}
	q.Set("id", id)

	secret, ok := dsn_map["secret"]

	if ok {
		q.Set("secret", secret)
	}

	secret_o, ok := dsn_map["secret_o"]

	if ok {
		q.Set("secret_o", secret_o)
	}

	uri_str := fmt.Sprintf("%s://%s?%s", IdSecretDriverName, source, q.Encode())
	return NewIdSecretURI(uri_str)
}

func NewIdSecretURI(str_uri string) (URI, error) {

	u, err := url.Parse(str_uri)

	if err != nil {
		return nil, err
	}

	source := u.Path

	q := u.Query()

	id := q.Get("id")

	if id == "" {
		return nil, errors.New("Missing id")
	}

	secret := q.Get("secret")
	secret_o := q.Get("secret_o")

	rnd_opts := random.DefaultOptions()
	rnd_opts.AlphaNumeric = true

	if secret == "" {

		s, err := random.String(rnd_opts)

		if err != nil {
			return nil, err
		}

		secret = s
	}

	if secret_o == "" {

		s, err := random.String(rnd_opts)

		if err != nil {
			return nil, err
		}

		secret_o = s
	}

	id_u := IdSecretURI{
		source:   source,
		id:       id,
		secret:   secret,
		secret_o: secret_o,
	}

	return &id_u, nil
}

func (u *IdSecretURI) URL() string {
	return u.id
}

func (u *IdSecretURI) String() string {

	q := url.Values{}
	q.Set("id", u.id)
	q.Set("secret", u.secret)
	q.Set("secret_o", u.secret_o)

	return fmt.Sprintf("%s://%s?%s", IdSecretDriverName, u.source, q.Encode())
}
