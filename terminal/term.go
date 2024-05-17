package terminal

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func GetTermWidth() int {
	width, _, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting terminal size: %v\n", err)
		return 150
	}
	return width
}
