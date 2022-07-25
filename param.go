package cli

import (
	"os"
	"strings"

	"github.com/Contra-Culture/report"
)

type (
	Param struct {
		check        func(report.Node, string) bool
		params       map[string]*Param
		name         string
		description  string
		defaultValue *string
	}
)

func (p *Param) prepare(r report.Node, given map[string]string, params map[string]string) bool {
	var v = given[p.name]
	if len(v) > 0 {
		r.Info("given param \"%s\": \"%s\"", p.name, v)
		if p.check != nil {
			if !p.check(r.Structure("check param \"%s\"=\"%s\"", p.name, v), v) {
				return false
			}
		}
		params[p.name] = v
	} else {
		r.Info("trying to get param \"%s\" from env variable", p.name)
		v = os.Getenv(p.name)
		if len(v) > 0 {
			if !p.check(r.Structure("check param \"%s\"=\"%s\"", p.name, v), v) {
				return false
			}
			params[p.name] = v
		} else {
			if p.defaultValue == nil {
				r.Error("param \"%s\" required", p.name)
				return false
			}
			r.Info("env variable does not provide \"%s\" parameter value", p.name)
			r.Info("default value \"%s\" for \"%s\" parameter picked", *p.defaultValue, p.name)
			params[p.name] = *p.defaultValue
		}
	}
	for _, child := range p.params {
		nr := r.Structure("prepare param \"%s\"", child.name)
		if !child.prepare(nr, given, params) {
			return false
		}
	}
	return true
}
func (p *Param) writeDoctringFragment(sb *strings.Builder) {
	sb.WriteString("\n\n\t\t-")
	sb.WriteString(p.name)
	sb.WriteString("\n\t\t\t")
	sb.WriteString(p.description)
	for _, p := range p.params {
		p.writeDoctringFragment(sb)
	}
}
