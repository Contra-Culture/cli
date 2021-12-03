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
		params      map[string]*Param
	}
)

func (c *Command) execute(r *report.RContext, givenParams map[string]string) {

}
