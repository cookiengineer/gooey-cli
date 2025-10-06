package utils

import "golang.org/x/term"
import "os"
import "strings"

func Prompt(question string) string {

	fd := int(os.Stdin.Fd())
	old_state, err0 := term.MakeRaw(fd)

	if err0 == nil {
		defer term.Restore(fd, old_state)
	}

	terminal := term.NewTerminal(os.Stdin, "> ")
	terminal.SetPrompt(question)

	line, err1 := terminal.ReadLine()

	if err1 == nil {
		return strings.TrimSpace(line)
	}


	return ""

}
