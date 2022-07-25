package main

import (
	"fmt"
	"time"

	"github.com/Contra-Culture/cli"
	"github.com/Contra-Culture/report"
)

func main() {
	now, err := time.Parse(time.RFC3339Nano, "2022-07-02T15:04:05.999999999-07:00")
	if err != nil {
		panic(err)
	}
	rc := report.ReportCreator(report.DumbTimer(now))
	app, r := cli.New(
		rc,
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
					cmd.Title("parameters")
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
					cmd.HandleWith(func(params map[string]string) (err error) {
						ln := params["name_lastname"]
						if len(ln) > 0 {
							fmt.Printf("\n\t > Hello %s %s!\n\n", params["name"], ln)
							return
						}
						fmt.Printf("\n\t > Hello %s!\n\n", params["name"])
						return
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
							p.Description("if passed upcase the text")
							p.Default("y")
							p.CheckWith(
								func(r report.Node, v string) bool {
									return true
								})
						})
				})
		})
	reportStr, _ := report.ToString(r, report.NoTime(), report.NoDuration())
	fmt.Print(reportStr)
	r = app.Handle()
	reportStr, _ = report.ToString(r, report.NoTime(), report.NoDuration())
	fmt.Print(reportStr)
}
