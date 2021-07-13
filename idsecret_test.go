package uri

import (
	"context"
	"fmt"
	"net/url"
	"path/filepath"
	"regexp"
	"testing"
)

func TestIdSecretURI(t *testing.T) {

	ctx := context.Background()

	candidates := []string{
		"idsecret:///tmp/example.jpg?id=1234",
	}

	for _, str_uri := range candidates {

		u, err := NewURI(ctx, str_uri)

		if err != nil {
			t.Fatalf("Failed to create new IIIF URI for '%s', %v", str_uri, err)
		}

		net_u, err := url.Parse(u.String())

		if err != nil {
			t.Fatalf("Failed to parse stringified IdSecreURI '%s', %v", u.String(), err)
		}

		net_q := net_u.Query()

		id := net_q.Get("id")
		secret := net_q.Get("secret")
		secret_o := net_q.Get("secret_o")

		if id != "1234" {
			t.Fatalf("Unexpected id value for '1234': '%s'", id)
		}

		if secret == "" {
			t.Fatalf("Missing secret value")
		}

		if secret_o == "" {
			t.Fatalf("Missing secret_o value")
		}

		format := "jpg"
		label := "x"

		values := &url.Values{}
		values.Set("format", format)
		values.Set("label", label)
		target, err := u.Target(values)

		if err != nil {
			t.Fatalf("Unable to determine target for '%s', %v", str_uri, err)
		}

		tree := Id2Path(id)

		root := filepath.Dir(target)
		fname := filepath.Base(target)

		if root != tree {
			t.Fatalf("Unexpected root: '%s'", root)
		}

		fname_pat := fmt.Sprintf("%s_([^_]+)_%s.%s", id, label, format)
		fname_re, err := regexp.Compile(fname_pat)

		if err != nil {
			t.Fatalf("Failed to compile fname pattern, %v", err)
		}

		if !fname_re.MatchString(fname) {
			t.Fatalf("Filename failed pattern match: '%s'", fname)
		}

	}

}
