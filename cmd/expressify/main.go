package main

import (
	"fmt"
	"os"
	"github.com/charmbracelet/bubbletea"
	"github.com/codersgyan/expressify/internal/cli_model"
	"github.com/codersgyan/expressify/internal/structure" // Ensure this package is imported for CopyDir
)

func main() {

	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Unable to get current working directory: %v\n", err)
		os.Exit(1)
	}

	// Define source and destination paths for the directory copy
	srcPath := cwd + "/.templates/jsbase"
	dstPath := cwd + "/.expressify/auth-service"

	// Copy the directory
	cpErr := structure.CopyDir(srcPath, dstPath)
	if cpErr != nil {
		fmt.Printf("Error copying directory: %s\n", cpErr)
		os.Exit(1)
	} else {
		fmt.Println("Directory copied successfully.")
	}

	// Run the CLI program using bubbletea
	p := tea.NewProgram(cli_model.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
