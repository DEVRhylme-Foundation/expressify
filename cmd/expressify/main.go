package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/codersgyan/expressify/internal/cli_model"
	"github.com/codersgyan/expressify/internal/errors"
	"github.com/codersgyan/expressify/internal/structure"
)

func main() {
	handleCLI()
}

func handleCLI() {
	cwd, err := os.Getwd()
	if err != nil {
		handleError(errors.NewSystemError(
			"Failed to get working directory",
			fmt.Sprintf("Error details: %v", err),
		))
		return
	}

	srcPath := cwd + "/.templates/jsbase"
	dstPath := cwd + "/.expressify/auth-service"

	if err := structure.CopyDir(srcPath, dstPath); err != nil {
		handleError(err)
		return
	}

	p := tea.NewProgram(cli_model.InitialModel())
	if _, err := p.Run(); err != nil {
		handleError(errors.NewRuntimeError(
			"CLI program error",
			fmt.Sprintf("Error running CLI: %v", err),
		))
		return
	}
}

func handleError(err *errors.AppError) {
	fmt.Printf("\nError Type: %s\n", err.Type)
	fmt.Printf("Message: %s\n", err.Message)
	fmt.Printf("Details: %s\n", err.Detail)
	fmt.Printf("Code: %d\n", err.Code)
	os.Exit(1)
}
