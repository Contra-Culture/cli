package cli

import (
	"github.com/Contra-Culture/report"
)

type (
	Command struct {
		handler     func(map[string]string) error
		params      []*Param
		name        string
		title       string
		description string
	}
)

func (c *Command) execute(r report.Node, given map[string]string, fallback func(error)) (ok bool) {
	r.Info("command \"%s\" called with: \"%#v\"", c.name, given)
	params := map[string]string{}
	for _, p := range c.params {
		nr := r.Structure("prepare param \"%s\"", p.name)
		ok = p.prepare(nr, given, params)
		if !ok {
			r.Error("param preparation failed")
			return
		}
	}
	err := c.handler(params)
	if err != nil {
		r.Error("command failed: %s", err.Error())
		r.Info("fallback called")
		fallback(err)
		return false
	}
	return true
}
