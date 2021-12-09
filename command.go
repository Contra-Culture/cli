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

func (c *Command) execute(r *report.RContext, params map[string]string, fallback func(error)) (ok bool) {
	err := c.handler(params)
	if err != nil {
		fallback(err)
		return false
	}
	return true
}
func (c *Command) prepareParams(r *report.RContext, gps map[string]string) (ps map[string]string, ok bool) {
	ps = map[string]string{}
	for _, p := range c.params {
		v := gps[p.name]
		v, ok = p.prepare(r.Contextf("prepare param \"%s\" with given: \"%s\"", p.name, v), v)
		if !ok {
			return nil, ok
		}
		ps[p.name] = v
	}
	return ps, true
}
