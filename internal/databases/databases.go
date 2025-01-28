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

var databases = []list.Item{
	list.Item(MongoDB),
	list.Item(PostgreSQL),
	list.Item(MySQL),
}

func NewDatabaseSelector() *selector.Selector {
	return selector.NewSelector("\n😎 Choose a database", databases)
}
