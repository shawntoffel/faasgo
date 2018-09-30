package faasgo

import (
	"testing"
)

func TestListFunctions(t *testing.T) {
	t.Skip()

	g := New(DefaultUrl)

	entries, err := g.ListFunctions()
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", entries)
}

func TestDeployFunction(t *testing.T) {
	t.Skip()

	g := New(DefaultUrl)

	f := Function{
		Service:    "test",
		Image:      "functions/nodeinfo:latest",
		EnvProcess: "node main.js",
	}

	err := g.DeployFunction(f)
	if err != nil {
		t.Error(err)
	}

	r := DeleteFunctionRequest{
		FunctionName: "test",
	}

	err = g.DeleteFunction(r)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateFunction(t *testing.T) {
	t.Skip()

	g := New(DefaultUrl)

	f := Function{
		Service:    "test",
		Image:      "functions/nodeinfo:latest",
		EnvProcess: "ls",
	}

	err := g.DeployFunction(f)
	if err != nil {
		t.Error(err)
	}

	err = g.UpdateFunction(f)
	if err != nil {
		t.Error(err)
	}

	r := DeleteFunctionRequest{
		FunctionName: "test",
	}

	err = g.DeleteFunction(r)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteFunction(t *testing.T) {
	t.Skip()

	g := New(DefaultUrl)

	f := Function{
		Service:    "test",
		Image:      "functions/nodeinfo:latest",
		EnvProcess: "node main.js",
	}

	err := g.DeployFunction(f)
	if err != nil {
		t.Error(err)
	}

	r := DeleteFunctionRequest{
		FunctionName: "test",
	}

	err = g.DeleteFunction(r)
	if err != nil {
		t.Error(err)
	}
}

func TestScaleFunction(t *testing.T) {
	t.Skip()

	g := New(DefaultUrl)

	f := Function{
		Service:    "test",
		Image:      "functions/nodeinfo:latest",
		EnvProcess: "node main.js",
	}

	err := g.DeployFunction(f)
	if err != nil {
		t.Error(err)
	}

	r := ScaleFunctionRequest{
		Service:  "test",
		Replicas: 1,
	}

	err = g.ScaleFunction("test", r)
	if err != nil {
		t.Error(err)
	}

	d := DeleteFunctionRequest{
		FunctionName: "test",
	}

	err = g.DeleteFunction(d)
	if err != nil {
		t.Error(err)
	}
}

func TestGetFunction(t *testing.T) {
	t.Skip()

	g := New(DefaultUrl)

	f := Function{
		Service:    "test",
		Image:      "functions/nodeinfo:latest",
		EnvProcess: "node main.js",
	}

	err := g.DeployFunction(f)
	if err != nil {
		t.Error(err)
	}

	result, err := g.Function("test")
	if err != nil {
		t.Error(err)
	}

	t.Log(result)

	d := DeleteFunctionRequest{
		FunctionName: "test",
	}

	err = g.DeleteFunction(d)
	if err != nil {
		t.Error(err)
	}
}

func TestSystemInfo(t *testing.T) {
	t.Skip()

	g := New(DefaultUrl)

	info, err := g.SystemInfo()
	if err != nil {
		t.Error(err)
	}

	t.Log(info)
}
