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
		params       map[string]*Param
	}
)

func (p *Param) prepare(r *report.RContext, gv string) (v string, ok bool) {
	switch len(gv) {
	case 0:
		v = os.Getenv(p.name)
		switch len(v) {
		case 0:
			v = p.defaultValue
			return v, len(v) > 0
		default:
			ok = p.check(r.Contextf("check param \"%s\"=\"%s\"", p.name, v), v)
			if ok {
				return
			}
		}
	default:
		ok = p.check(r.Contextf("check param \"%s\"=\"%s\"", p.name, v), v)
		if ok {
			return
		}
	}
	return "", false
}
