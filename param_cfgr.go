package cli

import (
	"github.com/Contra-Culture/report"
)

type (
	ParamCfgr struct {
		param  *Param
		report *report.RContext
	}
)

func (c *ParamCfgr) Name(n string) {
	if len(c.param.name) > 0 {
		c.report.Error("command param name already specified")
		return
	}
	c.param.name = n
}
func (c *ParamCfgr) Default(v string) {
	if len(c.param.description) > 0 {
		c.report.Error("default value already specified")
		return
	}
	c.param.defaultValue = v
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
	if !paramCfgr.check() {
		return
	}
	_, exists := c.param.params[param.name]
	if exists {
		c.report.Errorf("parameter \"-%s\" already specified", param.name)
		return
	}
	c.param.params[param.name] = param
}
func (c *ParamCfgr) check() (ok bool) {
	errCount := 0
	ok = len(c.param.name) > 0
	if !ok {
		c.report.Error("param name is not specified")
		errCount++
	}
	ok = len(c.param.description) > 0
	if !ok {
		c.report.Error("param description is not specified")
		errCount++
	}
	ok = len(c.param.question) > 0
	if !ok {
		c.report.Error("param question is not specified")
		errCount++
	}
	ok = c.param.check != nil
	if !ok {
		c.report.Error("param value checker is not specified")
		errCount++
	}
	ok = errCount == 0
	return
}
