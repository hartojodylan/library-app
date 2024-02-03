package main

import (
	"fmt"

	"github.com/gookit/gcli/v3"
)

var installOpts = struct {
	dirName    string
	fullName   string
	visualMode bool
	list       bool
	sample     bool
}{}

// test run: go build ./_examples/alone && ./alone -h
func main() {
	cmd := gcli.Command{
		Name:    "install",
		Aliases: []string{"ts"},
	}

	cmd.BoolOpt(
		&installOpts.visualMode,
		"visual", "v", false,
		"Prints the font name.",
	)
	cmd.StrOpt(
		&installOpts.dirName,
		"name", "n", "",
		"Choose a font name. Default is a random font.",
	)
	cmd.StrOpt(
		&installOpts.fullName,
		"full-name", "", "",
		"Choose a font name. Default is a random font.",
	)
	cmd.BoolOpt(
		&installOpts.list,
		"list", "", false,
		"Lists all available fonts.",
	)
	cmd.BoolOpt(
		&installOpts.sample,
		"sample",
		"",
		false,
		"Prints a sample with that font.",
	)

	cmd.Func = install

	// Alone Running
	cmd.MustRun(nil)
}

func install(_ *gcli.Command, args []string) error {
	gcli.Print("hello, in the alone command\n")

	fmt.Printf("opts %+v\n", installOpts)
	fmt.Printf("args is %v\n", args)

	return nil
}
