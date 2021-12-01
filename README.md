# cli - a toolkit for CLI apps

## Glossary

- **App** - is the entrypoint and a container for all the CLI commands of the the target you build.
- **Command** - is a first-class interface to your executable.
- **CommandInput** - is a parameter to a command
- **Handler** - is a function of command inputs that is responsible for command execution.

## Configuration

### Hierarchy

- **\*AppCfgr** - specifies CLI application
- - **\*CommandCfgr** - specifies a command, its handler and inputs.
- - - **\*CommandInputCfgr** - specifies command input.

###  Example

```Golang
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
						func(p *cli.CommandInputCfgr) {
							p.Name("filePath")
							p.Description("path to file")
							p.Question("Enter the file path")
						})
					cmd.Primary(
						func(p *cli.CommandInputCfgr) {
							p.Name("port")
							p.Description("port to listen")
							p.Question("Enter the port number")
						})
					cmd.Optional(
						func(p *cli.CommandInputCfgr) {
							p.Name("verbose")
							p.Description("verbose mode in which more detailed output is presented")
							p.Question("Do you want the verbose mode? y/n")
						})
				})
    })
}

```


## Conventions

1. **default command**
2. **flag and environment variable naming**
3. **version command**
4. **help command**
