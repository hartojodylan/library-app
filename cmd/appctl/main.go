package main

import (
	"github.com/dylanh/library-app/cli/commands"
	"github.com/gookit/gcli/v3"
)

// for test run: go build ./cmd/cliapp && ./cliapp
func main() {
	app := gcli.NewApp()
	app.Version = "1.0.3"

	// app.SetVerbose(gcli.VerbDebug)
	// app.DefaultCmd("exampl")

	app.Add(commands.GitCommand())
	app.Add(commands.InstallGoLintCommand())
	app.Add(commands.InstallSwagCommand())
	// app.Add(cmd.ColorCommand())
	//app.Add(builtin.GenAutoCompleteScript())
	// fmt.Printf("%+v\n", cliapp.CommandNames())
	app.Run(nil)
}
