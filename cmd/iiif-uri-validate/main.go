package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-iiif/go-iiif-uri"
	"log"
)

func main() {

	flag.Parse()

	ctx := context.Background()

	for _, str_uri := range flag.Args() {

		u, err := uri.NewURI(ctx, str_uri)

		if err != nil {
			msg := fmt.Sprintf("Invalid URI (%s) %s", str_uri, err)
			log.Fatal(msg)
		}

		origin := u.Origin()
		target, err := u.Target(nil)

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("OK '%s' %s %s\n", u.String(), origin, target)
	}
}
