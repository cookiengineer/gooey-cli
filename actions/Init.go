package actions

import "github.com/cookiengineer/gooey-cli/structs"
import "github.com/cookiengineer/gooey-cli/utils"
import "os"
import "time"

func Init(console *structs.Console) bool {

	var result bool

	console.Group("Init")

	time_start := time.Now()

	if utils.IsProgram("env") && utils.IsProgram("go") {

		folder, err := os.Getwd()

		if err == nil {

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
