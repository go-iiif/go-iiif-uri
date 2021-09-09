package uri

import (
	"context"
	"fmt"
	"testing"
)

func TestFileURI(t *testing.T) {

	ctx := context.Background()

	candidates := map[string]string{
		"/tmp/example.jpg":        "tmp/example.jpg",
		"tmp/example.jpg":         "tmp/example.jpg",
		"file:///tmp/example.jpg": "tmp/example.jpg",
		"file:///1746308155_248479.tif?target=174/630/815/5/tiles": "174/630/815/5/tiles",
	}

	for str_uri, expected := range candidates {

		u, err := NewURI(ctx, str_uri)

		if err != nil {
			t.Fatalf("Failed to create new IIIF URI for '%s', %v", str_uri, err)
		}

		target, err := u.Target(nil)

		if err != nil {
			t.Fatalf("Unable to determine target for '%s', %v", str_uri, err)
		}

		if target != expected {
			fmt.Printf("Unexpected target for '%s': target is: '%s' expected: '%s'\n", str_uri, target, expected)
		}
	}

}
