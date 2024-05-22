package util

import (
	"fmt"
	"os"

	"bsipiczki.com/jwt-go/model"
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
