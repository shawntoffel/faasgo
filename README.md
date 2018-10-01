# faasgo

[![Build Status](https://travis-ci.org/shawntoffel/faasgo.svg?branch=master)](https://travis-ci.org/shawntoffel/faasgo) [![GoDoc](https://godoc.org/github.com/shawntoffel/faasgo?status.svg)](https://godoc.org/github.com/shawntoffel/faasgo)  [![Go Report Card](https://goreportcard.com/badge/github.com/shawntoffel/faasgo)](https://goreportcard.com/report/github.com/shawntoffel/faasgo)

swaggerless golang api client for the [OpenFaaS gateway](https://github.com/openfaas/faas)

## Installation

```bash
go get github.com/shawntoffel/faasgo
```

## Environment Variables
OpenFaaS uses basic authentication. Credentials are read from the following environment variables:
```bash
export FAASGO_USER=admin
export FAASGO_PASS=<pass>
```

## Basic Usage

```go
gateway := faasgo.New(faasgo.DefaultUrl)

// service, image, and envprocess are required.
function := faasgo.Function{
	Service:    "nodeinfo",
	Image:      "functions/nodeinfo:latest",
	EnvProcess: "node main.js",
}

// Deploy a function
err := gateway.DeployFunction(function)
if err != nil {
	log.Fatal(err)
}

// List deployed functions
functions, err := gateway.ListFunctions()
if err != nil {
	log.Fatal(err)
}

// Invoke a function
resp, err := gateway.Invoke("nodeinfo", nil)
if err != nil {
	log.Fatal(err)
}

fmt.Println(string(resp))
```
