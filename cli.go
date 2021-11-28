package cli

import (
	"fmt"

	"github.com/Contra-Culture/report"
)

type (
	App struct {
		title       string
		description string
		version     string
		commands    map[string]*Command
	}
	AppCfgr struct {
		app    *App
		report *report.RContext
	}
	Command struct {
		name        string
		description string
		title       string
		handler     func(map[string]string) error
		primary     []*CommandInput
		flagIndex   map[string]*CommandInput
	}
	CommandCfgr struct {
		command *Command
		report  *report.RContext
	}
	CommandInput struct {
		name        string
		description string
		question    string
		flag        string
		envVar      string
		checker     func(string) *report.RContext
		primary     []*CommandInput
		flagIndex   map[string]*CommandInput
	}
	CommandInputCfgr struct {
		commandInput *CommandInput
		report       *report.RContext
	}
)

const (
	DEFAULT_COMMAND_NAME = ""
	HELP_COMMAND_NAME    = "help"
)

var reserverCommandNames = []string{
	DEFAULT_COMMAND_NAME,
	HELP_COMMAND_NAME,
}

func New(cfg func(*AppCfgr)) *App {
	app := &App{
		commands: map[string]*Command{},
	}
	cfg(
		&AppCfgr{
			app:    app,
			report: report.New("app"),
		})
	return app
}
func (c *AppCfgr) Version(v string) {
	if len(c.app.version) > 0 {
		c.report.Error("app version already specified")
		return
	}
	c.app.version = v
}
func (c *AppCfgr) Title(n string) {
	if len(c.app.title) > 0 {
		c.report.Error("app title already specified")
		return
	}
	c.app.title = n
}
func (c *AppCfgr) Description(d string) {
	if len(c.app.description) > 0 {
		c.report.Error("app description already specified")
		return
	}
	c.app.description = d
}
func (c *AppCfgr) Command(n string, cfg func(*CommandCfgr)) {
	if c.app.commands[n] != nil {
		c.report.Error(fmt.Sprintf("app has already \"%s\" command specified", n))
		return
	}
	var (
		command = &Command{
			name:      n,
			primary:   []*CommandInput{},
			flagIndex: map[string]*CommandInput{},
		}
		commandCfgr = &CommandCfgr{
			command: command,
			report:  c.report.Context(fmt.Sprintf("command \"%s\"", n)),
		}
	)
	cfg(commandCfgr)
	c.app.commands[n] = command
}
func (c *AppCfgr) Default(cfg func(*CommandCfgr)) {
	if c.app.commands[DEFAULT_COMMAND_NAME] != nil {
		c.report.Error("app default command already specified")
		return
	}
	var (
		command = &Command{
			name:      DEFAULT_COMMAND_NAME,
			primary:   []*CommandInput{},
			flagIndex: map[string]*CommandInput{},
		}
		commandCfgr = &CommandCfgr{
			command: command,
			report:  c.report.Context("default command"),
		}
	)
	cfg(commandCfgr)
	c.app.commands[DEFAULT_COMMAND_NAME] = command
}
func (c *CommandCfgr) Description(d string) {
	if len(c.command.description) > 0 {
		c.report.Error("command description already specified")
		return
	}
	c.command.description = d
}
func (c *CommandCfgr) Title(t string) {
	if len(c.command.title) > 0 {
		c.report.Error("command title already specified")
		return
	}
	c.command.title = t
}
func (c *CommandCfgr) HandleWith(handler func(map[string]string) error) {
	if c.command.handler != nil {
		c.report.Error("handler already specified")
		return
	}
	c.command.handler = handler
}
func (c *CommandCfgr) Optional(cfg func(*CommandInputCfgr)) {
	input := &CommandInput{
		flagIndex: map[string]*CommandInput{},
	}
	inputCfgr := &CommandInputCfgr{
		commandInput: input,
		report:       c.report.Context("optional param"),
	}
	cfg(inputCfgr)
	_, exists := c.command.flagIndex[input.name]
	if exists {
		c.report.Error(fmt.Sprintf("command input \"%s\" already specified", input.name))
		return
	}
	c.command.flagIndex[input.name] = input
}
func (c *CommandCfgr) Primary(cfg func(*CommandInputCfgr)) {
	input := &CommandInput{
		flagIndex: map[string]*CommandInput{},
	}
	inputCfgr := &CommandInputCfgr{
		commandInput: input,
		report:       c.report.Context("primary param"),
	}
	cfg(inputCfgr)
	_, exists := c.command.flagIndex[input.name]
	if exists {
		c.report.Error(fmt.Sprintf("command input \"%s\" already specified", input.name))
		return
	}
	c.command.primary = append(c.command.primary, input)
	c.command.flagIndex[input.name] = input
}
func (c *CommandInputCfgr) Flag(f string) {
	if len(c.commandInput.flag) > 0 {
		c.report.Error("command input flag already specified")
		return
	}
	c.commandInput.flag = f
}
func (c *CommandInputCfgr) Name(n string) {
	if len(c.commandInput.name) > 0 {
		c.report.Error("command input name already specified")
		return
	}
	c.commandInput.name = n
}
func (c *CommandInputCfgr) Description(d string) {
	if len(c.commandInput.description) > 0 {
		c.report.Error("command input description already specified")
		return
	}
	c.commandInput.description = d
}
func (c *CommandInputCfgr) Question(q string) {
	if len(c.commandInput.question) > 0 {
		c.report.Error("command input question already specified")
		return
	}
	c.commandInput.question = q
}
func (c *CommandInputCfgr) CheckWith(checker func(string) *report.RContext) {
	if c.commandInput.checker != nil {
		c.report.Error("checker already specified")
		return
	}
	c.commandInput.checker = checker
}
func (c *CommandInputCfgr) Primary(cfg func(*CommandInputCfgr)) {
	input := &CommandInput{
		flagIndex: map[string]*CommandInput{},
	}
	inputCfgr := &CommandInputCfgr{
		commandInput: input,
		report:       c.report.Context("primary flag"),
	}
	cfg(inputCfgr)
	_, exists := c.commandInput.flagIndex[input.name]
	if exists {
		c.report.Error(fmt.Sprintf("command input \"%s\" already specified", input.name))
		return
	}
	c.commandInput.primary = append(c.commandInput.primary, input)
	c.commandInput.flagIndex[input.name] = input
}
func (c *CommandInputCfgr) Optional(cfg func(*CommandInputCfgr)) {
	input := &CommandInput{
		flagIndex: map[string]*CommandInput{},
	}
	inputCfgr := &CommandInputCfgr{
		commandInput: input,
		report:       c.report.Context("optional flag"),
	}
	cfg(inputCfgr)
	_, exists := c.commandInput.flagIndex[input.name]
	if exists {
		c.report.Error(fmt.Sprintf("flag \"%s\" already specified", input.name))
		return
	}
	c.commandInput.flagIndex[input.name] = input
}
