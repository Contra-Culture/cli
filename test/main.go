package main

import (
	"fmt"

	"github.com/Contra-Culture/cli"
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
					cmd.Description("")
					cmd.Title("")
					cmd.HandleWith(func(map[string]string) error {
						return nil
					})
					cmd.Param(
						func(p *cli.ParamCfgr) {
							p.Name("filePath")
							p.Description("path to file")
							p.Question("Enter the file path")
						})
					cmd.Param(
						func(p *cli.ParamCfgr) {
							p.Name("port")
							p.Description("port to listen")
							p.Question("Enter the port number")
						})
					cmd.Param(
						func(p *cli.ParamCfgr) {
							p.Name("verbose")
							p.Description("verbose mode in which more detailed output is presented")
							p.Question("Do you want the verbose mode? y/n")
							p.Default("y")
						})
				})
			app.Command(
				"echo",
				func(cmd *cli.CommandCfgr) {
					cmd.Description("")
					cmd.Title("")
					cmd.HandleWith(func(map[string]string) error {
						return nil
					})
					cmd.Param(
						func(p *cli.ParamCfgr) {
							p.Name("message")
							p.Description("returns your message back")
						})
				})
			app.Command(
				"hello",
				func(cmd *cli.CommandCfgr) {
					cmd.Description("")
					cmd.Title("")
					cmd.HandleWith(func(map[string]string) error {
						return nil
					})
					cmd.Param(
						func(p *cli.ParamCfgr) {
							p.Name("name")
							p.Description("name for welcome")
							p.Question("Enter your name")
							p.Param(
								func(p *cli.ParamCfgr) {
									p.Name("lastname")
									p.Description("lastname for welcome")
									p.Question("Enter your lastname")
								})
						})
					cmd.Param(
						func(p *cli.ParamCfgr) {
							p.Name("upcase")
							p.Description("if passed upcaes the text")
							p.Question("Upcase name? y/n")
							p.Default("y")
						})
				})
		})
	if app != nil {
		fmt.Println("ok")
	}
	fmt.Print(r.String())
}
