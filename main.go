package main

import (
	_ "embed"

	//"github.com/LosAngeles971/cba-tool/cmd"
	"github.com/LosAngeles971/cba-tool/business/ui"
)

//go:embed example.yaml
var example []byte


func main() {
	app := ui.Build()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
