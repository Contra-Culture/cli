package cli

import (
	"github.com/Contra-Culture/report"
)

type (
	CommandCfgr struct {
		command *Command
		report  *report.RContext
	}
)

const (
	DEFAULT_COMMAND_NAME = "[default]"
	HELP_COMMAND_NAME    = "help"
	VERSION_COMMAND_NAME = "version"
	CONSOLE_COMMAND_NAME = "console"
)

var reservedCommandNames = []string{
	DEFAULT_COMMAND_NAME,
	HELP_COMMAND_NAME,
	VERSION_COMMAND_NAME,
	CONSOLE_COMMAND_NAME,
}

func (c *CommandCfgr) Title(t string) {
	if len(c.command.title) > 0 {
		c.report.Error("command title already specified")
		return
	}
	c.command.title = t
}
func (c *CommandCfgr) Description(d string) {
	if len(c.command.description) > 0 {
		c.report.Error("command description already specified")
		return
	}
	c.command.description = d
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
	for _, p := range c.command.params {
		if p.name == param.name {
			c.report.Errorf("command param \"%s\" already specified", param.name)
			return
		}
	}
	c.command.params = append(c.command.params, param)
}
func (c *CommandCfgr) check() (ok bool) {
	errCount := 0
	ok = len(c.command.title) > 0
	if !ok {
		c.report.Error("command title is not specified")
		errCount++
	}
	ok = len(c.command.description) > 0
	if !ok {
		c.report.Error("command description is not specified")
		errCount++
	}
	ok = len(c.command.name) > 0
	if !ok {
		c.report.Error("command name is not specified")
		errCount++
	}
	if c.command.handler == nil {
		c.report.Error("command handler is not specified")
		errCount++
	}
	return errCount == 0
}
