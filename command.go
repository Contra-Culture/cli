package cli

import (
	"github.com/Contra-Culture/report"
)

type (
	Command struct {
		name        string
		description string
		title       string
		handler     func(map[string]string) error
		params      []*Param
	}
)

func (c *Command) execute(r *report.RContext, ps map[string]string, fallback func(error)) (ok bool) {
	r.Infof("command \"%s\" called with: \"%#v\"", c.name, ps)
	params := map[string]string{}
	for _, p := range c.params {
		v := ps[p.name]
		v, ok = p.prepare(r.Contextf("prepare param \"%s\" with given: \"%s\"", p.name, v), v)
		if !ok {
			r.Errorf("parameter \"%s\" required", p.name)
			return
		}
		params[p.name] = v
	}
	err := c.handler(params)
	if err != nil {
		fallback(err)
		return false
	}
	return true
}
