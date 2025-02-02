package cli_model

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/codersgyan/expressify/internal/coding_styles"
	"github.com/codersgyan/expressify/internal/configs"
	"github.com/codersgyan/expressify/internal/databases"
	"github.com/codersgyan/expressify/internal/languages"
	"github.com/codersgyan/expressify/internal/loggers"
	"github.com/codersgyan/expressify/internal/orms"
	"github.com/codersgyan/expressify/internal/package_managers"
	"github.com/codersgyan/expressify/internal/selector"
	"github.com/codersgyan/expressify/internal/structure"
	"github.com/codersgyan/expressify/internal/test_frameworks"
)

var quitTextStyle = lipgloss.NewStyle().Margin(1, 0, 2, 4)

type AppState int

const (
	StateWelcome AppState = iota
	StateProjectName
	StateLanguage
	StatePackageManager
	StateTestFramework
	StateLoggerLibrary
	StateDatabase
	StateORM
	StateConfig
	StateCodingStyle
	StateFolderStructure
)

type CliModel struct {
	CurrentState           AppState
	ProjectNameInput       textinput.Model
	WelcomeMessage         string
	LanguageList           list.Model
	SelectedLanguage       string
	PackageManagerList     list.Model
	SelectedPackageManager string
	TestFrameworkList      list.Model
	SelectedTestFramework  string
	LoggerLibraryList      list.Model
	SelectedLoggerLibrary  string
	DatabaseList           list.Model
	SelectedDatabase       string
	ORMList                list.Model
	SelectedORM            string
	ConfigList             list.Model
	SelectedConfig         string
	CodingStyleList        list.Model
	SelectedCodingStyle    string
	FolderStructureCreated bool
	Error                  error
}

func (m CliModel) Init() tea.Cmd {
	return nil
}

type errMsg error

func (m *CliModel) handleEnter() tea.Cmd {
	switch m.CurrentState {
	case StateWelcome:
		m.CurrentState = StateProjectName
		m.ProjectNameInput.Focus()
	case StateProjectName:
		m.CurrentState = StateLanguage
	case StateLanguage:
		i, ok := m.LanguageList.SelectedItem().(selector.Item)
		if ok {
			m.SelectedLanguage = string(i)
		}
		m.CurrentState = StatePackageManager
	case StatePackageManager:
		i, ok := m.PackageManagerList.SelectedItem().(selector.Item)
		if ok {
			m.SelectedPackageManager = string(i)
		}
		m.CurrentState = StateTestFramework
	case StateTestFramework:
		i, ok := m.TestFrameworkList.SelectedItem().(selector.Item)
		if ok {
			m.SelectedTestFramework = string(i)
		}
		m.CurrentState = StateLoggerLibrary
	case StateLoggerLibrary:
		i, ok := m.LoggerLibraryList.SelectedItem().(selector.Item)
		if ok {
			m.SelectedLoggerLibrary = string(i)
		}
		m.CurrentState = StateDatabase
	case StateDatabase:
		i, ok := m.DatabaseList.SelectedItem().(selector.Item)
		if ok {
			m.SelectedDatabase = string(i)
		}
		m.CurrentState = StateORM
	case StateORM:
		i, ok := m.ORMList.SelectedItem().(selector.Item)
		if ok {
			m.SelectedORM = string(i)
		}
		m.CurrentState = StateConfig
	case StateConfig:
		i, ok := m.ConfigList.SelectedItem().(selector.Item)
		if ok {
			m.SelectedConfig = string(i)
		}
		m.CurrentState = StateCodingStyle
	case StateCodingStyle:
		i, ok := m.CodingStyleList.SelectedItem().(selector.Item)
		if ok {
			m.SelectedCodingStyle = string(i)
		}
		m.CurrentState = StateFolderStructure
	case StateFolderStructure:
		err := structure.CreateBaseFileStructure(m.ProjectNameInput.Value(), m.SelectedLanguage)
		if err != nil {
			fmt.Printf("error creating folder structure: %v", err)
			m.Error = err
			return tea.Quit
		}
		// todo: transition to next state
	}
	return nil
}

func (m CliModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.LanguageList.SetWidth(msg.Width)
		m.ProjectNameInput.Width = msg.Width
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			cmd = m.handleEnter()
		case tea.KeyEsc, tea.KeyCtrlC:
			return m, tea.Quit
		}
	case errMsg:
		m.Error = msg
		return m, tea.Quit
	}

	switch m.CurrentState {
	case StateProjectName:
		m.ProjectNameInput, cmd = m.ProjectNameInput.Update(msg)
	case StateLanguage:
		m.LanguageList, cmd = m.LanguageList.Update(msg)
	case StatePackageManager:
		m.PackageManagerList, cmd = m.PackageManagerList.Update(msg)
	case StateTestFramework:
		m.TestFrameworkList, cmd = m.TestFrameworkList.Update(msg)
	case StateLoggerLibrary:
		m.LoggerLibraryList, cmd = m.LoggerLibraryList.Update(msg)
	case StateDatabase:
		m.DatabaseList, cmd = m.DatabaseList.Update(msg)
	case StateORM:
		m.ORMList, cmd = m.ORMList.Update(msg)
	case StateConfig:
		m.ConfigList, cmd = m.ConfigList.Update(msg)
	case StateCodingStyle:
		m.CodingStyleList, cmd = m.CodingStyleList.Update(msg)
	}

	return m, cmd
}

func (m CliModel) View() string {
	var s string

	switch m.CurrentState {
	case StateWelcome:
		s = m.WelcomeMessage
	case StateProjectName:
		s = "\n📗 Enter your project name:\n\n" + m.ProjectNameInput.View()
	case StateLanguage:
		if m.SelectedLanguage != "" {
			var str string
			if m.SelectedLanguage == string(languages.JavaScript) {
				str = "🎉 Awesome choice! JavaScript brings flexibility and dynamism to your project. Let's get coding! 🚀"
			} else if m.SelectedLanguage == string(languages.TypeScript) {
				str = "👍 Great pick! TypeScript adds type safety and robustness to your application. Time to build! 🏗️"
			}
			return quitTextStyle.Render(fmt.Sprintf(str))
		}
		return m.LanguageList.View()
	case StatePackageManager:
		if m.SelectedPackageManager != "" {
			var str string
			if m.SelectedPackageManager == string(package_managers.NPM) {
				str = "🎉 Awesome choice! NPM is the world's most popular package manager. Let's get coding! 🚀"
			} else if m.SelectedPackageManager == string(package_managers.PNPM) {
				str = "👍 Great pick! PNPM is a fast, disk space efficient package manager. Time to build! 🏗️"
			}
			return quitTextStyle.Render(fmt.Sprintf(str))
		}
		return m.PackageManagerList.View()
	case StateTestFramework:
		if m.SelectedTestFramework != "" {
			var str string
			if m.SelectedTestFramework == string(test_frameworks.SuperTestWithJest) {
				str = "🎉 Awesome choice! Supertest is best framework 🚀"
			} else if m.SelectedTestFramework == string(test_frameworks.MochaWithChaiHTTP) {
				str = "👍 Great pick! Mocha with Chai is powerful framework"
			}
			return quitTextStyle.Render(fmt.Sprintf(str))
		}
		return m.TestFrameworkList.View()
	case StateLoggerLibrary:
		if m.SelectedLoggerLibrary != "" {
			var str string
			if m.SelectedLoggerLibrary == string(loggers.Winston) {
				str = "🎉 Awesome choice! Winston is best logger out there 🚀"
			} else if m.SelectedLoggerLibrary == string(loggers.Bunyan) {
				str = "👍 Great pick! Bunyan is powerful logger"
			} else if m.SelectedLoggerLibrary == string(loggers.Pino) {
				str = "👍 Great pick! Pino is powerful logger"
			}
			return quitTextStyle.Render(fmt.Sprintf(str))
		}
		return m.LoggerLibraryList.View()
	case StateDatabase:
		if m.SelectedDatabase != "" {
			var str string
			if m.SelectedDatabase == string(databases.MongoDB) {
				str = "🎉 Awesome choice! MongoDB is powerful NoSQL database 🚀"
		