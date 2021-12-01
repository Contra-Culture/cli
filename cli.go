package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/Contra-Culture/report"
)

type (
	App struct {
		title           string
		description     string
		version         string
		errorHandler    func(error)
		shotdownHandler func()
		timeoutHandler  func()
		commands        map[string]*Command
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
		inputs      map[string]*CommandInput
	}
	CommandCfgr struct {
		command *Command
		report  *report.RContext
	}
	CommandInput struct {
		name         string
		description  string
		question     string
		defaultValue string
		check        func(*report.RContext, string) bool
		inputs       map[string]*CommandInput
	}
	CommandInputCfgr struct {
		commandInput *CommandInput
		report       *report.RContext
	}
)

const (
	DEFAULT_COMMAND_NAME = "[default]"
	HELP_COMMAND_NAME    = "help"
	VERSION_COMMAND_NAME = "version"
)

var reservedCommandNames = []string{
	DEFAULT_COMMAND_NAME,
	HELP_COMMAND_NAME,
	VERSION_COMMAND_NAME,
}

func New(cfg func(*AppCfgr)) (app *App, r *report.RContext) {
	app = &App{
		commands: map[string]*Command{},
	}
	r = report.New("app")
	cfg(
		&AppCfgr{
			app:    app,
			report: r,
		})
	return
}
func (a *App) Handle() (r *report.RContext) {
	r = report.New(a.title)
	var (
		cmdName      = os.Args[1]
		parsedParams = map[string]string{}
		params       = map[string]string{}
		cmd          *Command
	)
	c := r.Context("command parsing")
	switch cmdName[0] {
	case '-':
		c.Info("default command picked")
		cmd = a.commands[DEFAULT_COMMAND_NAME]
	default:
		c.Infof("%s command picked", cmdName)
		cmd = a.commands[cmdName]
	}
	c = r.Context("params parsing")
	for _, p := range os.Args[1:] {
		keyVal := strings.Split(p, "=")
		switch len(keyVal) {
		case 2:
			key := keyVal[0]
			val := keyVal[1]
			if key[0] != '-' {
				c.Errorf("wrong parameter name \"%s\": should start with \"-\"", key)
				panic("keys must start with \"-\"")
			}
			key = key[1 : len(key)-1]
			if val[0] == '"' && val[len(val)-1] == '"' || val[0] == '\'' && val[len(val)-1] == '\'' {
				val = val[1 : len(val)-2]
			}
			parsedParams[key] = val
		default:
			panic("wrong parameters")
		}
	}
	c = r.Context("params parsing")
	for cn, ci := range cmd.inputs {
		if !ci.check(c, parsedParams[cn]) {
			return
		}
		params[cn] = parsedParams[cn]
	}
	cmd.handler(params)
	return
}
func (a *App) DocString() string {
	return ""
}
func (c *AppCfgr) Timeout() {

}
func (c *AppCfgr) Shutdown() {

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
func (c *AppCfgr) HandleErrorsWith(h func(error)) {
	if c.app.errorHandler != nil {
		c.report.Error("errors handler already specified")
		return
	}
	c.app.errorHandler = h
}
func (c *AppCfgr) Command(n string, cfg func(*CommandCfgr)) {
	for _, rn := range reservedCommandNames {
		if n == rn {
			c.report.Error(fmt.Sprintf("command name \"%s\" is reserved", n))
			return
		}
	}
	if c.app.commands[n] != nil {
		c.report.Error(fmt.Sprintf("app has already \"%s\" command specified", n))
		return
	}
	var (
		command = &Command{
			name:   n,
			inputs: map[string]*CommandInput{},
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
			name:   DEFAULT_COMMAND_NAME,
			inputs: map[string]*CommandInput{},
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
func (c *CommandCfgr) Input(cfg func(*CommandInputCfgr)) {
	var (
		input = &CommandInput{
			inputs: map[string]*CommandInput{},
		}
		inputCfgr = &CommandInputCfgr{
			commandInput: input,
			report:       c.report.Context("parameter"),
		}
	)
	cfg(inputCfgr)
	_, exists := c.command.inputs[input.name]
	if exists {
		c.report.Error(fmt.Sprintf("command input \"%s\" already specified", input.name))
		return
	}
	c.command.inputs[input.name] = input
}
func (c *CommandInputCfgr) Default(v string) {
	if len(c.commandInput.description) > 0 {
		c.report.Error("default value already specified")
		return
	}
	c.commandInput.defaultValue = v
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
func (c *CommandInputCfgr) CheckWith(checker func(*report.RContext, string) bool) {
	if c.commandInput.check != nil {
		c.report.Error("checker already specified")
		return
	}
	c.commandInput.check = checker
}
func (c *CommandInputCfgr) Input(cfg func(*CommandInputCfgr)) {
	var (
		input = &CommandInput{
			inputs: map[string]*CommandInput{},
		}
		inputCfgr = &CommandInputCfgr{
			commandInput: input,
			report:       c.report.Context("parameter"),
		}
	)
	cfg(inputCfgr)
	_, exists := c.commandInput.inputs[input.name]
	if exists {
		c.report.Error(fmt.Sprintf("parameter \"-%s\" already specified", input.name))
		return
	}
	c.commandInput.inputs[input.name] = input
}
