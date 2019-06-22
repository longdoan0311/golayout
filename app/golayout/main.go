package main

import (
	"github.com/techcomsecurities/golayout"
	"github.com/techcomsecurities/golayout/app/golayout/cmd"
	"github.com/gobuffalo/packr/v2"
)

func init() {
	golayout.TemplateBox = packr.New("myBox", "./template")
}

func main() {
	cmd.Execute()
}
