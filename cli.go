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
		params      map[string]*Param
	}
	CommandCfgr struct {
		command *Command
		report  *report.RContext
	}
	Param struct {
		name         string
		description  string
		question     string
		defaultValue string
		check        func(*report.RContext, string) bool
		params       map[string]*Param
	}
	ParamCfgr struct {
		param  *Param
		report *report.RContext
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
	r = report.New("app configuration")
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
				return
			}
			key = key[1 : len(key)-1]
			if val[0] == '"' && val[len(val)-1] == '"' || val[0] == '\'' && val[len(val)-1] == '\'' {
				val = val[1 : len(val)-2]
			}
			parsedParams[key] = val
		default:
			c.Errorf("wrong parameter -<key>=<value> pair: \"%s\"", keyVal)
			return
		}
	}
	c = r.Context("params parsing")
	for cn, ci := range cmd.params {
		pp, passed := parsedParams[cn]
		if !passed {
			pp = os.Getenv(ci.name)
		}
		if !ci.check(c, pp) {
			return
		}
		params[cn] = pp
		for cn, ci := range ci.params {
			pp, passed := parsedParams[cn]
			if !passed {
				pp = os.Getenv(ci.name)
			}
			if !ci.check(c, pp) {
				return
			}
			params[cn] = pp
		}
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
			params: map[string]*Param{},
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
			params: map[string]*Param{},
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
func (c *CommandCfgr) Param(cfg func(*ParamCfgr)) {
	var (
		param = &Param{
			params: map[string]*Param{},
		}
		paramCfgr = &ParamCfgr{
			param:  param,
			report: c.report.Context("parameter"),
		}
	)
	cfg(paramCfgr)
	_, exists := c.command.params[param.name]
	if exists {
		c.report.Error(fmt.Sprintf("command param \"%s\" already specified", param.name))
		return
	}
	c.command.params[param.name] = param
}
func (c *ParamCfgr) Default(v string) {
	if len(c.param.description) > 0 {
		c.report.Error("default value already specified")
		return
	}
	c.param.defaultValue = v
}
func (c *ParamCfgr) Name(n string) {
	if len(c.param.name) > 0 {
		c.report.Error("command param name already specified")
		return
	}
	c.param.name = n
}
func (c *ParamCfgr) Description(d string) {
	if len(c.param.description) > 0 {
		c.report.Error("command param description already specified")
		return
	}
	c.param.description = d
}
func (c *ParamCfgr) Question(q string) {
	if len(c.param.question) > 0 {
		c.report.Error("command param question already specified")
		return
	}
	c.param.question = q
}
func (c *ParamCfgr) CheckWith(checker func(*report.RContext, string) bool) {
	if c.param.check != nil {
		c.report.Error("checker already specified")
		return
	}
	c.param.check = checker
}
func (c *ParamCfgr) Param(cfg func(*ParamCfgr)) {
	var (
		param = &Param{
			params: map[string]*Param{},
		}
		paramCfgr = &ParamCfgr{
			param:  param,
			report: c.report.Context("parameter"),
		}
	)
	cfg(paramCfgr)
	_, exists := c.param.params[param.name]
	if exists {
		c.report.Error(fmt.Sprintf("parameter \"-%s\" already specified", param.name))
		return
	}
	c.param.params[param.name] = param
}
