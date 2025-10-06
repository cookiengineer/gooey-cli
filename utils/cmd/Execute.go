package cmd

import "os/exec"

func Execute(folder string, program string, args []string) bool {

	var result bool

	cmd := exec.Command(program, args...)
	cmd.Dir = folder

	err := cmd.Run()

	if err == nil {
		result = true
	}

	return result

}
