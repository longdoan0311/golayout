package main

import (
	"github.com/3t-dev/golayout"
	"github.com/3t-dev/golayout/app/golayout/cmd"
	"github.com/gobuffalo/packr/v2"
)

func init() {
	golayout.TemplateBox = packr.New("myBox", "./template")
}

func main() {
	cmd.Execute()
}
