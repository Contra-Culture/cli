package cli

type (
	App struct {
		name        string
		description string
		commands    map[string]*Command
	}
	AppCfgr struct {
		app *App
	}
	Command struct {
		name        string
		description string
		cmd         string
		args        []*Argument
		params      []*Param
		Quest       *Quest
	}
	CommandCfgr struct {
		command *Command
	}
	Argument struct {
		name        string
		description string
	}
	Param struct {
		name        string
		description string
		flag        string
	}
	Quest struct {
		question           string
		param              *Param
		quests             map[string]*Quest
		answerQuestNameMap func(string) (string, error)
	}
	QuestCfgr struct {
		quest *Quest
	}
)

func New(cfg func(*AppCfgr)) *App {
	app := &App{}
	cfg(
		&AppCfgr{
			app: app,
		})
	return app
}
func (c *AppCfgr) Name(n string) {

}
func (c *AppCfgr) Description(d string) {

}
func (c *AppCfgr) Command() {

}
func (c *AppCfgr) DefaultCommand() {

}
func (c *CommandCfgr) Name(n string) {

}
func (c *CommandCfgr) Description(d string) {

}
func (c *CommandCfgr) Cmd(cmd string) {

}
func (c *CommandCfgr) Map(p func(string) bool, qn string) {

}
