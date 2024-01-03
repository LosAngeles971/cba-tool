package main

import (
	_ "embed"
	
	//"github.com/LosAngeles971/cba-tool/cmd"
	"github.com/LosAngeles971/cba-tool/business/ui"
	"github.com/LosAngeles971/cba-tool/business/cba"
)

//go:embed example.yaml
var example []byte

func main() {
	d := cba.NewCBA(example)
	app := &ui.CBAToolApp{
		Data: d,
	}
	app.Build()
	app.Run()
	//cmd.Execute()
}
