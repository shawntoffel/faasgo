package main

import (
	"log"

	"github.com/shawntoffel/faasgo"
)

func main() {
	g := faasgo.New(faasgo.DefaultUrl)

	// service, image, and envprocess are required.
	f := faasgo.Function{
		Service:    "nodeinfo",
		Image:      "functions/nodeinfo:latest",
		EnvProcess: "node main.js",
	}

	err := g.DeployFunction(f)
	if err != nil {
		log.Fatal(err)
	}

	function, err := g.Function("nodeinfo")
	if err != nil {
		log.Fatal(err)
	}

	log.Print(function)
}
