package main

import (
	//"github.com/LosAngeles971/cba-tool/cmd"
	"github.com/LosAngeles971/cba-tool/business/ui"
)

func main() {
	app := ui.Build()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
