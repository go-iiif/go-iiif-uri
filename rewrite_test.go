package uri

import (
	"context"
	"net/url"
	"testing"
)

func TestRewriteURI(t *testing.T) {

	ctx := context.Background()

	candidates := map[string]string{
		"rewrite:///tmp/bob.jpg?target=alice.jpg":                            "rewrite:///tmp/bob.jpg?target=alice.jpg",
		"rewrite:///usr/local/images/99400/99481/99481.jpg?target=99481.jpg": "rewrite:///usr/local/images/99400/99481/99481.jpg?target=99481.jpg",
	}

	for str_uri, expected_str := range candidates {

		u, err := NewURI(ctx, str_uri)

		if err != nil {
			t.Fatalf("Failed to create new IIIF URI for '%s', %v", str_uri, err)
		}

		if u.String() != expected_str {
			t.Fatalf("Unexpected string value for '%s': '%s'", str_uri, u.String())
		}

		target, err := u.Target(nil)

		if err != nil {
			t.Fatalf("Unable to determine target for '%s', %v", str_uri, err)
		}

		ru, err := url.Parse(str_uri)

		if err != nil {
			t.Fatalf("Failed to parse URI, %v", err)
		}

		rq := ru.Query()

		if target != rq.Get("target") {
			t.Fatalf("Unexpected target for '%s': '%s'", str_uri, target)
		}
	}

}
