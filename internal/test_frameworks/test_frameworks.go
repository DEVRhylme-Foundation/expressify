package test_frameworks

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/codersgyan/expressify/internal/selector"
)

type TestFramework string

const (
	SuperTestWithJest TestFramework = "SuperTest with Jest"
	MochaWithChaiHTTP TestFramework = "Mocha with Chai HTTP"
)

var testFrameworks = []TestFramework{
	SuperTestWithJest,
	MochaWithChaiHTTP,
}

// NewTestFrameworkSelector creates a new test framework selector with available options
func NewTestFrameworkSelector() *selector.Selector {
	var items []list.Item
	for _, framework := range testFrameworks {
		items = append(items, list.Item(framework))
	}
	return selector.NewSelector("\n😎 Choose a test framework", items)
}
