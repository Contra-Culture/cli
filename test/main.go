package main

import (
	"fmt"

	"github.com/Contra-Culture/cli"
)

func main() {
	app := cli.New(
		func(app *cli.AppCfgr) {
			app.Title("testapp")
			app.Version("0.0.1 (test)")
			app.Description("testapp is a test application which is an example of use of github.com/Contra-Culture/cli library.")
			app.Default(
				func(cmd *cli.CommandCfgr) {
					cmd.Description("")
					cmd.Title("")
					cmd.HandleWith(func(map[string]string) error {
						return nil
					})
					cmd.Primary(
						"file-path",
						func(p *cli.CommandInputCfgr) {
							p.Name("filePath")
							p.Description("path to file")
							p.Question("Enter the file path")
						})
					cmd.Primary(
						"port",
						func(p *cli.CommandInputCfgr) {
							p.Name("port")
							p.Description("port to listen")
							p.Question("Enter the port number")
						})
					cmd.Optional(
						"verbose",
						func(p *cli.CommandInputCfgr) {
							p.Name("verbose")
							p.Description("verbose mode in which more detailed output is presented")
							p.Question("Do you want the verbose mode? y/n")
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
					cmd.Primary(
						"message",
						func(p *cli.CommandInputCfgr) {
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
					cmd.Primary(
						"name",
						func(p *cli.CommandInputCfgr) {
							p.Name("name")
							p.Description("name for welcome")
							p.Question("Enter your name")
						})
					cmd.Optional(
						"upcase",
						func(p *cli.CommandInputCfgr) {
							p.Name("upcase")
							p.Description("if passed upcaes the text")
							p.Question("Upcase name? y/n")
						})
				})
		})
	if app != nil {
		fmt.Println("ok")
	}
}
