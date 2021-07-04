package uri

import (
	"context"
	"testing"
)

func TestRewriteURI(t *testing.T) {

	ctx := context.Background()
	
	candidates := []string{
		"rewrite:///tmp/bob.jpg?target=alice.jpg",
	}

	for _, str_uri := range candidates {

		u, err := NewURI(ctx, str_uri)

		if err != nil {
			t.Fatalf("Failed to create new IIIF URI for '%s', %v", str_uri, err)
		}

		if u.String() != "rewrite:///tmp/bob.jpg?target=alice.jpg" {
			t.Fatalf("Unexpected string value for '%s': '%s'", str_uri, u.String())
		}

		target, err := u.Target(nil)

		if err != nil {
			t.Fatalf("Unable to determine target for '%s', %v", str_uri, err)
		}

		if target != "alice.jpg" {
			t.Fatalf("Unexpected target for '%s': '%s'", str_uri, target)
		}
	}
	
}
