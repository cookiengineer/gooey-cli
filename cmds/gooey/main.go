package main

import "github.com/cookiengineer/gooey-cli/actions"
import "github.com/cookiengineer/gooey-cli/structs"
import "os"
import "strings"

func showInitUsage(console *structs.Console) {

	console.Info("")
	console.Info("gooey toolchain")
	console.Info("")
	console.Group("Usage: gooey [Action] [--parameter=value...]")
	console.Log("")
	console.Log("Execute this program inside your Gooey Project's root folder.")
	console.Log("If a required parameter is not set, the program will prompt for user input.")
	console.Log("")
	console.Log("Custom Views require a custom Controller with the identical \"name\".")
	console.Log("")
	console.GroupEnd("------")

	console.Group("Action          | Parameters            | Description                                             \\")
	console.Log("----------------|-----------------------|---------------------------------------------------------|")
	console.Log("init            | --name=example-app    | The unique identifier of the Web App                    |")
	console.Log("                | --title=\"Example App\" | The title of the Web App                                |")
	console.Log("----------------|-----------------------|---------------------------------------------------------|")
	console.Log("init-component  | --name=Example        | The unique identifier of the Component                  |")
	console.Log("                | --source=ui/Table     | The Gooey Component that the Component is based on      |")
	console.Log("----------------|-----------------------|---------------------------------------------------------|")
	console.Log("init-controller | --name=example        | The unique identifier of the Controller                 |")
	console.Log("                | --label=Example       | The label that is displayed in the Header               |")
	console.Log("                | --path=/example.html  | The unique path for routing the Controller and its View |")
	console.Log("----------------|-----------------------|---------------------------------------------------------|")
	console.Log("init-view       | --name=example        | The unique identifier of the View                       |")
	console.Log("----------------|-----------------------|---------------------------------------------------------|")
	console.Log("init-webview    | --vendored=yes|no     | Vendor webview module                                   |")
	console.Log("                | --patched=yes|no      | Patch webview module CGo files to installed version     |")
	console.GroupEnd("----------------|-----------------------|---------------------------------------------------------/")

}

func main() {

	console := structs.NewConsole(os.Stdout, os.Stderr, 0)

	if len(os.Args) == 2 {

		action := strings.TrimSpace(strings.ToLower(os.Args[1]))

		if action == "init" {

			result := actions.Init(console)

			if result == true {
				os.Exit(0)
			} else {
				os.Exit(1)
			}

		} else if action == "init-component" {

			// TODO

		} else if action == "init-controller" {

			// TODO

		} else if action == "init-view" {

			// TODO

		} else if action == "init-webview" {

			// TODO

		} else {

			showInitUsage(console)
			os.Exit(1)

		}

	} else {

		showInitUsage(console)
		os.Exit(1)

	}

}
