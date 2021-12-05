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

				})
			app.Default(
				func(cmd *cli.CommandCfgr) {
					cmd.Description("print its params")
					cmd.Title("no-title")
					cmd.HandleWith(func(map[string]string) error {
						return nil
					})
					cmd.Param(
						func(p *cli.ParamCfgr) {
							p.Name("filePath")
							p.Description("path to file")
							p.Question("Enter the file path")
							p.CheckWith(
								func(r *report.RContext, v string) bool {
									return true
								})
						})
					cmd.Param(
						func(p *cli.ParamCfgr) {
							p.Name("port")
							p.Description("port to listen")
							p.Question("Enter the port number")
							p.CheckWith(
								func(r *report.RContext, v string) bool {
									return true
								})
						})
					cmd.Param(
						func(p *cli.ParamCfgr) {
							p.Name("verbose")
							p.Description("verbose mode in which more detailed output is presented")
							p.Question("Do you want the verbose mode? y/n")
							p.Default("y")
							p.CheckWith(
								func(r *report.RContext, v string) bool {
									return true
								})
						})
				})
			app.Command(
				"echo",
				func(cmd *cli.CommandCfgr) {
					cmd.Description("prints your text")
					cmd.Title("echo")
					cmd.HandleWith(func(map[string]string) error {
						return nil
					})
					cmd.Param(
						func(p *cli.ParamCfgr) {
							p.Name("message")
							p.Description("returns your message back")
							p.CheckWith(
								func(r *report.RContext, v string) bool {
									return true
								})
						})

				})
			app.Command(
				"hello",
				func(cmd *cli.CommandCfgr) {
					cmd.Description("prints hello message")
					cmd.Title("hello")
					cmd.HandleWith(func(map[string]string) error {
						return nil
					})
					cmd.Param(
						func(p *cli.ParamCfgr) {
							p.Name("name")
							p.Description("name for welcome")
							p.Question("Enter your name")
							p.CheckWith(
								func(r *report.RContext, v string) bool {
									return true
								})
							p.Param(
								func(p *cli.ParamCfgr) {
									p.Name("lastname")
									p.Description("lastname for welcome")
									p.Question("Enter your lastname")
									p.CheckWith(
										func(r *report.RContext, v string) bool {
											return true
										})
								})
						})
					cmd.Param(
						func(p *cli.ParamCfgr) {
							p.Name("upcase")
							p.Description("if passed upcaes the text")
							p.Question("Upcase name? y/n")
							p.Default("y")
							p.CheckWith(
								func(r *report.RContext, v string) bool {
									return true
								})
						})
				})
		})
	if app != nil {
		fmt.Println("ok")
	}
	fmt.Print(r.String())
	r = app.Handle()
	fmt.Print(r.String())
}
