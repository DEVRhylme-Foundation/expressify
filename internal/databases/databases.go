package databases

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/codersgyan/expressify/internal/selector"
)

type Database string

const (
	MongoDB    Database = "MongoDB"
	PostgreSQL Database = "PostgreSQL"
	MySQL      Database = "MySQL"
)

// Implement the list.Item interface for Database type
func (d Database) FilterValue() string {
	return string(d)
}

var databases = []list.Item{
	MongoDB,
	PostgreSQL,
	MySQL,
}

func NewDatabaseSelector() *selector.Selector {
	return selector.NewSelector("\n😎 Choose a database", databases)
}
