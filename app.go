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
		report         report.Node
	}
)

func New(cfg func(*AppCfgr)) (app *App, r report.Node) {
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
func (a *App) Handle() (r report.Node) {
	var (
		cmd         *Command
		cmdName     string
		params      = map[string]string{}
		paramsPairs = []string{}
	)
	r = a.report
	r.Info("called with params: \"%#v\"", os.Args)
	if len(os.Args) == 1 {
		cmd = a.defaultCommand
		r.Info("default command picked")
	} else {
		cmdName = os.Args[1]
		if cmdName[0] == '-' {
			r.Info("default command picked")
			cmd = a.defaultCommand
			paramsPairs = os.Args[1:]
		} else {
			for _, _cmd := range a.commands {
				if _cmd.name == cmdName {
					cmd = _cmd
					r.Info("command \"%s\" picked", cmd.title)
					paramsPairs = os.Args[2:]
					break
				}
			}
			if cmd == nil {
				r.Error("unknown command \"%s\"", cmdName)
				return
			}
		}
		for _, _pp := range paramsPairs {
			pp := strings.Split(_pp, "=")
			switch len(pp) {
			case 2:
				params[pp[0][1:]] = pp[1]
			default:
				r.Error("wrong parameter format \"%s\"", _pp)
				return
			}
		}
	}
	cmd.execute(
		r.Structure("command \"%s\" execution", cmd.title),
		params,
		a.errorHandler,
	)
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
				if p.defaultValue != nil {
					sb.WriteString(" (optional, default: ")
					sb.WriteString(*p.defaultValue)
					sb.WriteString(")")
				}
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
				p.writeDoctringFragment(&sb)
			}
		}
	}
	sb.WriteRune('\n')
	return sb.String()
}
