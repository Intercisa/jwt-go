package util

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"bsipiczki.com/jwt-go/model"
	"github.com/atotto/clipboard"
	"golang.org/x/term"
)

func GetTermWidth() int {
	width, _, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting terminal size: %v\n", err)
		return model.TERM_WIDTH_DEFAULT
	}
	return width
}

func GetEnv(tokenKey string) string {
	value, exists := os.LookupEnv(tokenKey)
	if !exists {
		return model.DefaultToken
	}
	return value
}

func CheckBoolEnv(envKey string) bool {
	value, exists := os.LookupEnv(envKey)
	if !exists {
		return false
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}
	return boolValue
}

func CopyToClippboard(value string) {
	err := clipboard.WriteAll(value)
	if err != nil {
		fmt.Println("Failed to copy to clipboard:", err)
		return
	}
}

func ClearTerminal() {
	switch runtime.GOOS {
	case "linux", "darwin":
		fmt.Print("\033[H\033[2J")
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Unsupported platform!")
	}
}
