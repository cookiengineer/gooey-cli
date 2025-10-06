package main

import "github.com/cookiengineer/gooey-cli/actions"
import "github.com/cookiengineer/gooey-cli/structs"
import "os"
import "strings"

func getParameter(name string) string {

	var result string

	for _, arg := range os.Args {

		if strings.HasPrefix(arg, "--" + name + "=") {

			tmp := strings.TrimSpace(arg[len(name)+3:])

			if strings.HasPrefix(tmp, "\"") && strings.HasSuffix(tmp, "\"") {
				result = strings.TrimSpace(tmp[1:len(tmp)-1])
			} else if strings.HasPrefix(tmp, "'") && strings.HasSuffix(tmp, "'") {
				result = strings.TrimSpace(tmp[1:len(tmp)-1])
			} else {
				result = strings.TrimSpace(tmp)
			}

		}

	}

	return result

}

func showInitUsage(console *structs.Console) {

	console.Info("")
	console.Info("gooey bootstrapper")
	console.Info("")
	console.Group("Usage: gooey [Action] [--parameter=value...]")
	console.Log("")
	console.Log("Execute this program inside your Gooey Project's root folder.")
	console.Log("If a required parameter is not set, the program will prompt for user input.")
	console.Log("")
	console.Log("Custom Views require a custom Controller with the identical \"name\".")
	console.Log("")
	console.GroupEnd("------")

	console.Group("Action          | Parameters            | Description                                         \\")
	console.Log("----------------|-----------------------|-----------------------------------------------------|")
	console.Log("init            | --name=example-app    | Unique identifier of the Web App                    |")
	console.Log("                | --title=\"Example App\" | Title of the Web App                                |")
	console.Log("                | --themed=yes|no       | Use Gooey Default Theme                             |")
	console.Log("----------------|-----------------------|-----------------------------------------------------|")
	console.Log("init-component  | --name=Example        | Unique identifier of the Component                  |")
	console.Log("                | --source=ui/Table     | Gooey Component which the Component is based on     |")
	console.Log("----------------|-----------------------|-----------------------------------------------------|")
	console.Log("init-controller | --name=example        | Unique identifier of the Controller                 |")
	console.Log("                | --label=Example       | Label which is displayed in the Header              |")
	console.Log("                | --path=/example.html  | Unique path for routing the Controller and its View |")
	console.Log("----------------|-----------------------|-----------------------------------------------------|")
	console.Log("init-view       | --name=example        | The unique identifier of the View                   |")
	console.Log("----------------|-----------------------|-----------------------------------------------------|")
	console.Log("init-webview    | --version=4.0|4.1|6.0 | Patch webview files to selected WebKitGTK version   |")
	console.GroupEnd("----------------|-----------------------|-----------------------------------------------------/")

}

func main() {

	console := structs.NewConsole(os.Stdout, os.Stderr, 0)

	if len(os.Args) >= 2 {

		action := strings.TrimSpace(strings.ToLower(os.Args[1]))

		if action == "init" {

			app_name  := getParameter("name")
			app_title := getParameter("title")

			result := actions.Init(console, app_name, app_title)

			if result == true {
				os.Exit(0)
			} else {
				os.Exit(1)
			}

		} else if action == "init-component" {

			component_name   := getParameter("name")
			component_source := getParameter("source")

			result := actions.InitComponent(console, component_name, component_source)

			if result == true {
				os.Exit(0)
			} else {
				os.Exit(1)
			}

		} else if action == "init-controller" {

			controller_name  := getParameter("name")
			controller_label := getParameter("label")
			controller_path  := getParameter("path")

			result := actions.InitController(console, controller_name, controller_label, controller_path)

			if result == true {
				os.Exit(0)
			} else {
				os.Exit(1)
			}

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
