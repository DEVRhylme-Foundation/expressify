package structure

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/codersgyan/expressify/internal/errors"
	"github.com/codersgyan/expressify/internal/languages"
)

func CreateBaseFileStructure(projectName string, language string) *errors.AppError {
	cwd, err := os.Getwd()
	if err != nil {
		return errors.NewSystemError(
			"Failed to get working directory",
			fmt.Sprintf("Error details: %v", err),
		)
	}

	projectPath := filepath.Join(cwd, ".expressify", projectName)
	if _, err := os.Stat(projectPath); !os.IsNotExist(err) {
		return errors.NewValidationError(
			"Project already exists",
			fmt.Sprintf("A project with name '%s' already exists at path: %s", projectName, projectPath),
		)
	}

	if err := validateProjectName(projectName); err != nil {
		return errors.NewValidationError(
			"Invalid project name",
			err.Error(),
		)
	}

	mkdirProjectDirErr := os.MkdirAll(projectPath, 0755)
	if mkdirProjectDirErr != nil {
		return errors.NewSystemError(
			"Failed to create project directory",
			fmt.Sprintf("Error creating directory at %s: %v", projectPath, mkdirProjectDirErr),
		)
	}

	var languageWisePath string
	if language == string(languages.JavaScript) {
		languageWisePath = "jsbase"
	} else if language == string(languages.TypeScript) {
		languageWisePath = "tsbase"
	} else {
		return errors.NewValidationError(
			"Invalid language selection",
			fmt.Sprintf("Language '%s' is not supported. Use 'JavaScript' or 'TypeScript'", language),
		)
	}

	srcPath := filepath.Join(cwd, ".templates", languageWisePath)
	dstPath := filepath.Join(projectPath)

	if err := CopyDir(srcPath, dstPath); err != nil {
		return errors.NewSystemError(
			"Failed to copy template files",
			fmt.Sprintf("Error copying from %s to %s: %v", srcPath, dstPath, err),
		)
	}

	return nil
}

func validateProjectName(name string) error {
	if name == "" {
		return fmt.Errorf("project name cannot be empty")
	}
	if len(name) > 214 {
		return fmt.Errorf("project name too long (max 214 characters)")
	}
	// Add more validation rules as needed
	return nil
}

func CopyFile(src, dst string) *errors.AppError {
	sourceFile, err := os.Open(src)
	if err != nil {
		return errors.NewSystemError(
			"Failed to open source file",
			fmt.Sprintf("Error opening %s: %v", src, err),
		)
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return errors.NewSystemError(
			"Failed to create destination file",
			fmt.Sprintf("Error creating %s: %v", dst, err),
		)
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, sourceFile); err != nil {
		return errors.NewSystemError(
			"Failed to copy file contents",
			fmt.Sprintf("Error copying from %s to %s: %v", src, dst, err),
		)
	}

	return nil
}

func CopyDir(src, dst string) *errors.AppError {
	info, err := os.Stat(src)
	if err != nil {
		return errors.NewSystemError(
			"Failed to get source directory info",
			fmt.Sprintf("Error getting info for %s: %v", src, err),
		)
	}

	if err := os.MkdirAll(dst, info.Mode()); err != nil {
		return errors.NewSystemError(
			"Failed to create destination directory",
			fmt.Sprintf("Error creating directory at %s: %v", dst, err),
		)
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return errors.NewSystemError(
			"Failed to read source directory",
			fmt.Sprintf("Error reading directory %s: %v", src, err),
		)
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err := CopyDir(srcPath, dstPath); err != nil {
				return err // Already wrapped in AppError
			}
		} else {
			if err := CopyFile(srcPath, dstPath); err != nil {
				return err // Already wrapped in AppError
			}
		}
	}

	return nil
}
