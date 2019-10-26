package main

import (
	"log"
	"time"

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

	log.Println("deploying function")
	err := g.DeployFunction(f)
	if err != nil {
		log.Fatalf("failed to deploy function: %s", err.Error())
	}

	log.Println("Waiting a bit for function to finish deploying")
	time.Sleep(5 * time.Second)

	resp, err := g.Invoke("nodeinfo", nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(string(resp))

}
