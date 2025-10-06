package actions

import "github.com/cookiengineer/gooey-cli/structs"
import "github.com/cookiengineer/gooey-cli/utils"
import "os"
import "strings"
import "time"

func Init(console *structs.Console, app_name string, app_title string) bool {

	var result bool

	console.Group("Init")

	if app_name == "" {
		app_name = utils.Prompt("Please enter the App Name (default: \"gooey-boilerplate\"): ")
	}

	if app_title == "" {
		app_title = utils.Prompt("Please enter the App Title (default: \"Gooey Boilerplate\"): ")
	}

	if app_name != "" && strings.ToLower(app_name) == app_name {
		console.Log("App Name: \"" + app_name + "\"")
	} else {
		app_name = "gooey-boilerplate"
		console.Warn("App Name: \"" + app_name + "\"")
	}

	if app_title != "" {
		console.Log("App Title: \"" + app_title + "\"")
	} else {
		app_title = "Gooey Boilerplate"
		console.Warn("App Title: \"" + app_title + "\"")
	}

	time_start := time.Now()

	if utils.IsProgram("env") && utils.IsProgram("go") {

		folder, err := os.Getwd()

		if err == nil {

			// TODO: If go.mod does not exist, then call go mod init app-name

			// TODO: Copy Boilerplate assets
			// TODO: Copy /public/favicon.ico
			// TODO: Copy /public/index.html
			// TODO: Copy /public/site.webmanifest

			// TODO: Copy /public/wasm_exec.js
			// TODO: Copy /public/wasm_init.js

			// TODO: Copy /public/design/app

			// TODO: If Gooey Theme is selected, copy design/gooey folder to public
			// TODO: If Gooey Theme is selected, then modify main.go (to register components)

			// TODO: Copy main.go
			// TODO: Modify main.go if Gooey Theme is selected

			console.Log("Folder is \"" + folder + "\"")

		}

	} else {

		if utils.IsProgram("env") == false {
			console.Error("Cannot find \"env\": Please execute \"sudo pacman -S coreutils\"")
		}

		if utils.IsProgram("go") == false {
			console.Error("Cannot find \"go\": Please execute \"sudo pacman -S go\"")
		}

	}

	if result == true {

		time_end := time.Now()
		duration := time_end.Sub(time_start)

		console.Log("Finished in " + duration.String())

	}

	console.GroupEnd("Init")

	return result

}
