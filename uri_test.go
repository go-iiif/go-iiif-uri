package uri

import (
	"context"
	"testing"
)

func TestNewURI(t *testing.T) {

	ctx := context.Background()

	candidates := []string{
		"/tmp/example.jpg",
		"file:///tmp/example.jpg",
		"idsecret:///tmp/example.jpg?id=1234",
		"rewrite:///tmp/example.jpg?target=bob.jpg",
	}

	for _, str_uri := range candidates {

		_, err := NewURI(ctx, str_uri)

		if err != nil {
			t.Fatalf("Failed to create new IIIF URI for '%s', %v", str_uri, err)
		}
	}
}
