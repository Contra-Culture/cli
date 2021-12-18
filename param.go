package cli

import (
	"os"

	"github.com/Contra-Culture/report"
)

type (
	Param struct {
		check        func(report.Node, string) bool
		params       map[string]*Param
		name         string
		description  string
		defaultValue string
	}
)

func (p *Param) prepare(r report.Node, given map[string]string, params map[string]string) (ok bool) {
	var v = given[p.name]
	switch len(v) {
	case 0:
		r.Info("trying to get param \"%s\" from env variable", p.name)
		v = os.Getenv(p.name)
		switch len(v) {
		case 0:
			r.Info("env variable does not provide \"%s\" parameter value", p.name)
			r.Info("default value \"%s\" for \"%s\" parameter picked", p.defaultValue, p.name)
			v = p.defaultValue
			ok = len(v) > 0
			if !ok {
				r.Error("praram \"%s\" required", p.name)
				return
			}
		default:
			ok = p.check(r.Structure("check param \"%s\"=\"%s\"", p.name, v), v)
		}
	default:
		ok = p.check(r.Structure("check param \"%s\"=\"%s\"", p.name, v), v)
		if ok {
			params[p.name] = v
		}
	}
	for _, child := range p.params {
		nr := r.Structure("prepare param \"%s\"", p.name)
		ok = child.prepare(nr, given, params)
		if !ok {
			return
		}
	}
	return
}
