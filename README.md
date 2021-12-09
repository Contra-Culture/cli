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

See github.com/ContraCulture/cli/test/main.go

## Conventions

1. **default command**
2. **flag and environment variable naming**
3. **version command**
4. **help command**
