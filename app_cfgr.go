package cli

import (
	"fmt"

	"github.com/Contra-Culture/report"
)

type (
	AppCfgr struct {
		app    *App
		report *report.RContext
	}
)

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

func (c *AppCfgr) check() (ok bool) {
	errCount := 0
	ok = len(c.app.title) > 0
	if !ok {
		c.report.Error("")
		errCount++
	}
	ok = len(c.app.description) > 0
	if !ok {
		c.report.Error("")
		errCount++
	}
	_, ok = c.app.commands[DEFAULT_COMMAND_NAME]
	if !ok {
		c.report.Error("")
		errCount++
	}
	c.app.commands[HELP_COMMAND_NAME] = &Command{}
	c.app.commands[VERSION_COMMAND_NAME] = &Command{}
	c.app.commands[CONSOLE_COMMAND_NAME] = &Command{}
	return errCount == 0
}
