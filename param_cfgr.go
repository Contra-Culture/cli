package cli

import (
	"fmt"

	"github.com/Contra-Culture/report"
)

type (
	ParamCfgr struct {
		param  *Param
		parent *Param
		report report.Node
	}
)

func (c *ParamCfgr) Name(n string) {
	if len(c.param.name) > 0 {
		c.report.Error("command param name already specified")
		return
	}
	c.report.Info("param name \"%s\"", n)
	if c.parent != nil {
		c.param.name = fmt.Sprintf("%s_%s", c.parent.name, n)
	} else {
		c.param.name = n
	}
}
func (c *ParamCfgr) Default(v string) {
	if c.param.defaultValue != nil {
		c.report.Error("default value already specified: %#v", *c.param.defaultValue)
		return
	}
	c.report.Info("param default value \"%s\"", v)
	c.param.defaultValue = &v
}
func (c *ParamCfgr) Description(d string) {
	if len(c.param.description) > 0 {
		c.report.Error("command param description already specified")
		return
	}
	c.report.Info("param description \"%s\"", d)
	c.param.description = d
}
func (c *ParamCfgr) CheckWith(checker func(report.Node, string) bool) {
	if c.param.check != nil {
		c.report.Error("checker already specified")
		return
	}
	c.report.Info("param checker specified")
	c.param.check = checker
}
func (c *ParamCfgr) Param(cfg func(*ParamCfgr)) {
	var (
		param = &Param{
			params: map[string]*Param{},
		}
		paramCfgr = &ParamCfgr{
			param:  param,
			parent: c.param,
			report: c.report.Structure("dependent parameter"),
		}
	)
	cfg(paramCfgr)
	if !paramCfgr.check() {
		return
	}
	_, exists := c.param.params[param.name]
	if exists {
		c.report.Error("parameter \"-%s\" already specified", param.name)
		return
	}
	c.param.params[param.name] = param
}
func (c *ParamCfgr) check() bool {
	errCount := 0
	if len(c.param.name) == 0 {
		c.report.Error("param name is not specified")
		errCount++
	}
	if len(c.param.description) == 0 {
		c.report.Error("param description is not specified")
		errCount++
	}
	if c.param.check == nil {
		c.report.Error("param value checker is not specified")
		errCount++
	}
	return errCount == 0
}
