package main

import (
	"fmt"

	"github.com/Contra-Culture/cli"
	"github.com/Contra-Culture/report"
)

func main() {
	app, r := cli.New(
		func(app *cli.AppCfgr) {
			app.Title("testapp")
			app.Version("0.0.1 (test)")
			app.Description("testapp is a test application which is an example of use of github.com/Contra-Culture/cli library.")
			app.HandleErrorsWith(
				func(err error) {
					fmt.Printf("error: %s", err.Error())
				})
			app.Default(
				func(cmd *cli.CommandCfgr) {
					cmd.Description("prints its params")
					cmd.Title("no-title")
					cmd.HandleWith(func(params map[string]string) error {
						fmt.Printf("\n\nparams: %#v\n", params)
						return nil
					})
					cmd.Param(
						func(p *cli.ParamCfgr) {
							p.Name("filePath")
							p.Description("path to file")
							p.CheckWith(
								func(r report.Node, v string) bool {
									return true
								})
						})
					cmd.Param(
						func(p *cli.ParamCfgr) {
							p.Name("port")
							p.Description("port to listen")
							p.CheckWith(
								func(r report.Node, v string) bool {
									return true
								})
						})
					cmd.Param(
						func(p *cli.ParamCfgr) {
							p.Name("verbose")
							p.Description("verbose mode in which more detailed output is presented")
							p.Default("y")
							p.CheckWith(
								func(r report.Node, v string) bool {
									return true
								})
						})
				})
			app.Command(
				"echo",
				func(cmd *cli.CommandCfgr) {
					cmd.Description("prints your text")
					cmd.Title("echo")
					cmd.HandleWith(func(params map[string]string) (err error) {
						fmt.Printf("\necho > %s\n", params["message"])
						return
					})
					cmd.Param(
						func(p *cli.ParamCfgr) {
							p.Name("message")
							p.Description("returns your message back")
							p.CheckWith(
								func(r report.Node, v string) bool {
									return true
								})
						})
				})
			app.Command(
				"hello",
				func(cmd *cli.CommandCfgr) {
					cmd.Description("prints hello message")
					cmd.Title("hello")
					cmd.HandleWith(func(params map[string]string) error {
						if len(params["name.lastname"]) > 0 {
							fmt.Printf("\nHello %s %s!\n\n", params["name"], params["name.lastname"])
						} else {
							fmt.Printf("\nHello %s!\n\n", params["name"])
						}
						return nil
					})
					cmd.Param(
						func(p *cli.ParamCfgr) {
							p.Name("name")
							p.Description("name for hello")
							p.CheckWith(
								func(r report.Node, v string) bool {
									return true
								})
							p.Param(
								func(p *cli.ParamCfgr) {
									p.Name("lastname")
									p.Description("lastname for hello")
									p.Default("")
									p.CheckWith(
										func(r report.Node, v string) bool {
											return true
										})
								})
						})
					cmd.Param(
						func(p *cli.ParamCfgr) {
							p.Name("upcase")
							p.Description("if passed upcaes the text")
							p.Default("y")
							p.CheckWith(
								func(r report.Node, v string) bool {
									return true
								})
						})
				})
		})
	fmt.Print(report.ToString(r))
	r = app.Handle()
	fmt.Print(report.ToString(r))
}
