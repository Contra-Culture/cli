package cli

import (
	"os"
	"strings"

	"github.com/Contra-Culture/report"
)

type (
	App struct {
		title          string
		description    string
		version        string
		errorHandler   func(error)
		commands       []*Command
		defaultCommand *Command
		report         *report.RContext
	}
)

func New(cfg func(*AppCfgr)) (app *App, r *report.RContext) {
	app = &App{
		commands: []*Command{},
		report:   report.New("app initialization"),
	}
	r = report.New("app configuration")
	appCfgr := &AppCfgr{
		app:    app,
		report: r,
	}
	cfg(appCfgr)
	appCfgr.check()
	return
}
func (a *App) command() (c *Command, ps map[string]string, ok bool) {
	a.report.Infof("checking command...")
	var args []string
	ps = map[string]string{}
	switch len(os.Args) {
	case 1:
		a.report.Info("default command picked")
		c = a.defaultCommand
		args = os.Args[1:]
	default:
		cmdName := os.Args[1]
		switch cmdName[0] {
		case '-':
			a.report.Info("default command picked")
			c = a.defaultCommand
			args = os.Args[1:]
		default:
			a.report.Infof("%s command picked", cmdName)
			for _, _c := range a.commands {
				if _c.name == cmdName {
					c = _c
					args = os.Args[2:]
					break
				}
			}
		}
	}
	for _, arg := range args {
		keyVal := strings.Split(arg, "=")
		switch len(keyVal) {
		case 2:
			key := keyVal[0]
			switch key[0] {
			case '-':
				key = key[1:]
			default:
				a.report.Errorf("wrong parameter key \"%s\"", key)
				c = nil
				ps = nil
				return
			}
			val := keyVal[1]
			ps[key] = val
		default:
			a.report.Errorf("wrong -<key>=<value> string: \"%s\"", arg)
			c = nil
			ps = nil
			return
		}
	}
	ps, ok = c.prepareParams(a.report, ps)
	return
}
func (a *App) Handle() (r *report.RContext) {
	r = a.report
	cmd, params, ok := a.command()
	if !ok {
		r.Info("exit because of error")
		return
	}
	ok = cmd.execute(
		r.Contextf("command %s execution", cmd.title),
		params,
		a.errorHandler,
	)
	if !ok {
		panic(r.String())
	}
	return
}
func (a *App) DocString() string {
	var sb strings.Builder
	sb.WriteString(a.title)
	sb.WriteString(" (")
	sb.WriteString(a.version)
	sb.WriteString(")\n\t")
	sb.WriteString(a.description)
	sb.WriteString("\n\n\t>> commands <<")
	if a.defaultCommand != nil {
		c := a.defaultCommand
		sb.WriteString("\n\n\t")
		sb.WriteString(c.name)
		sb.WriteString(" - ")
		sb.WriteString(c.title)
		sb.WriteString("\n\t\t")
		sb.WriteString(c.description)
		if len(c.params) > 0 {
			sb.WriteString("\n\n\t\t>> parameters <<")
			for _, p := range c.params {
				sb.WriteString("\n\n\t\t-")
				sb.WriteString(p.name)
				sb.WriteString("\n\t\t\t")
				sb.WriteString(p.description)
			}
		}
	}
	for _, c := range a.commands {
		sb.WriteString("\n\n\t")
		sb.WriteString(c.name)
		sb.WriteString(" - ")
		sb.WriteString(c.title)
		sb.WriteString("\n\t\t")
		sb.WriteString(c.description)
		if len(c.params) > 0 {
			sb.WriteString("\n\n\t\t>> parameters <<")
			for _, p := range c.params {
				sb.WriteString("\n\n\t\t-")
				sb.WriteString(p.name)
				sb.WriteString("\n\t\t\t")
				sb.WriteString(p.description)
			}
		}
	}
	sb.WriteRune('\n')
	return sb.String()
}
