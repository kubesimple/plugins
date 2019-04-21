package main

import (
	"log"

	"github.com/kubesimple/plugins/pkg/transform"
)

func main() {
	env, err := transform.GetEnvironment()
	if err != nil {
		log.Fatal(err)
	}

}
