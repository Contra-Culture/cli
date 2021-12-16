package cli

import (
	"os"

	"github.com/Contra-Culture/report"
)

type (
	Param struct {
		name         string
		description  string
		defaultValue string
		check        func(*report.RContext, string) bool
		parent       *Param
		params       map[string]*Param
	}
)

func (p *Param) prepare(r *report.RContext, gv string) (v string, ok bool) {
	switch len(gv) {
	case 0:
		r.Infof("trying to get param \"%s\" from env variable", p.name)
		v = os.Getenv(p.name)
		switch len(v) {
		case 0:
			r.Infof("env variable does not provide \"%s\" parameter value", p.name)
			r.Infof("default value \"%s\" for \"%s\" parameter picked", p.defaultValue, p.name)
			v = p.defaultValue
			return v, len(v) > 0
		default:
			ok = p.check(r.Contextf("check param \"%s\"=\"%s\"", p.name, v), v)
		}
	default:
		ok = p.check(r.Contextf("check param \"%s\"=\"%s\"", p.name, gv), gv)
		if ok {
			v = gv
		}
	}
	return
}
