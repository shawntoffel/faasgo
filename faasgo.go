package faasgo

import (
	"net/http"
)

const DefaultUrl = "http://127.0.0.1:8080"

type Gateway struct {
	BaseUrl string
	Client  *http.Client
}

func New(endpoint string) Gateway {
	return NewWithClient(endpoint, &http.Client{})
}

func NewWithClient(endpoint string, client *http.Client) Gateway {
	return Gateway{BaseUrl: endpoint, Client: client}
}

func (g Gateway) ListFunctions() ([]FunctionListEntry, error) {
	entries := []FunctionListEntry{}

	err := g.request("GET", g.BaseUrl+SystemFunctionsEndpoint, nil, &entries)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func (g Gateway) DeployFunction(f Function) error {
	err := g.request("POST", g.BaseUrl+SystemFunctionsEndpoint, f, nil)
	if err != nil {
		return err
	}

	return nil
}

func (g Gateway) UpdateFunction(f Function) error {
	err := g.request("PUT", g.BaseUrl+SystemFunctionsEndpoint, f, nil)
	if err != nil {
		return err
	}

	return nil
}

func (g Gateway) DeleteFunction(r DeleteFunctionRequest) error {
	err := g.request("DELETE", g.BaseUrl+SystemFunctionsEndpoint, r, nil)
	if err != nil {
		return err
	}

	return nil
}

func (g Gateway) Invoke(name string, body []byte) ([]byte, error) {
	return g.simpleRequest("POST", g.BaseUrl+FunctionEndpoint+"/"+name, body)
}

func (g Gateway) InvokeAsync(name string, body []byte) error {
	_, err := g.simpleRequest("POST", g.BaseUrl+AsyncFunctionEndpoint+"/"+name, body)

	return err
}

func (g Gateway) ScaleFunction(name string, r ScaleFunctionRequest) error {
	err := g.request("POST", g.BaseUrl+SystemScaleFunctionEndpoint+"/"+name, r, nil)
	if err != nil {
		return err
	}

	return nil
}

func (g Gateway) Function(name string) (FunctionListEntry, error) {
	f := FunctionListEntry{}

	err := g.request("GET", g.BaseUrl+SystemFunctionEndpoint+"/"+name, nil, &f)
	if err != nil {
		return f, err
	}

	return f, nil
}

func (g Gateway) SystemInfo() (Info, error) {
	i := Info{}

	err := g.request("GET", g.BaseUrl+SystemInfoEndpoint, nil, &i)
	if err != nil {
		return i, err
	}

	return i, nil
}
