package uri

import (
	"errors"
	"fmt"
	"github.com/aaronland/go-string/dsn"
	"github.com/aaronland/go-string/random"
	_ "log"
	"net/url"
	"path/filepath"
	"strconv"
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
	id       int64
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

	str_id := q.Get("id")

	if str_id == "" {
		return nil, errors.New("Missing id")
	}

	id, err := strconv.ParseInt(str_id, 10, 64)

	if err != nil {
		return nil, err
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

func (u *IdSecretURI) Root() string {
	return id2Path(u.id)
}

func (u *IdSecretURI) Base() string {
	return strconv.FormatInt(u.id, 10)
}

func (u *IdSecretURI) String() string {

	q := url.Values{}
	q.Set("id", strconv.FormatInt(u.id, 10))
	q.Set("secret", u.secret)
	q.Set("secret_o", u.secret_o)

	return fmt.Sprintf("%s://%s?%s", IdSecretDriverName, u.source, q.Encode())
}

func id2Path(id int64) string {

	parts := []string{}
	input := strconv.FormatInt(id, 10)

	for len(input) > 3 {

		chunk := input[0:3]
		input = input[3:]
		parts = append(parts, chunk)
	}

	if len(input) > 0 {
		parts = append(parts, input)
	}

	return filepath.Join(parts...)
}
