package cli

import (
	"fmt"

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
)

var reservedCommandNames = []string{
	DEFAULT_COMMAND_NAME,
	HELP_COMMAND_NAME,
	VERSION_COMMAND_NAME,
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
