package main

import (
	"flag"
	"fmt"
	"github.com/go-iiif/go-iiif-uri"
	"log"
)

func main() {

	flag.Parse()

	for _, str_uri := range flag.Args() {

		u, err := uri.NewURI(str_uri)

		if err != nil {
			msg := fmt.Sprintf("Invalid URI (%s) %s", str_uri, err)
			log.Fatal(msg)
		}

		log.Printf("OK '%s' %s\n", u.String(), uri.Path(u))
	}
}
