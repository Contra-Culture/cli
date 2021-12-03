package cli

import (
	"os"
	"strings"

	"github.com/Contra-Culture/report"
)

type (
	App struct {
		title           string
		description     string
		version         string
		errorHandler    func(error)
		shotdownHandler func()
		timeoutHandler  func()
		commands        map[string]*Command
		report          *report.RContext
	}
)

func New(cfg func(*AppCfgr)) (app *App, r *report.RContext) {
	app = &App{
		commands: map[string]*Command{},
	}
	r = report.New("app configuration")
	cfg(
		&AppCfgr{
			app:    app,
			report: r,
		})
	return
}
func (a *App) command() (c *Command, ps map[string]string, ok bool) {
	a.report.Infof("checking command...")
	args := []string{}
	switch len(os.Args) {
	case 1:
		a.report.Info("default command picked")
		c = a.commands[DEFAULT_COMMAND_NAME]
	default:
		cmdName := os.Args[1]
		switch cmdName[0] {
		case '-':
			a.report.Info("default command picked")
			c = a.commands[DEFAULT_COMMAND_NAME]
			args = os.Args[1:]
		default:
			a.report.Infof("%s command picked", cmdName)
			c = a.commands[cmdName]
			args = os.Args[2:]
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
	ok = true
	return
}
func (a *App) Handle() (r *report.RContext) {
	r = a.report
	cmd, givenParams, ok := a.command()
	if !ok {
		r.Info("exit because of error")
		return
	}
	cmd.execute(r.Contextf("command %s execution", cmd.title), givenParams)
	return
}
func (a *App) DocString() string {
	return ""
}
