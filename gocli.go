package gocli

import (
	"github.com/jawher/mow.cli"
)

type CommandHandler func(*cli.Cmd)

type Command struct {
	Name string
	Desc string
	cli.CmdInitializer
}

var commands []Command = make([]Command, 0)

//Register will register a command into app command list
func Register(name, desc string, init cli.CmdInitializer) {
	commands = append(commands, Command{name, desc, init})
}

//Start will run specific command handler
func Start(args []string, cliAppName, cliAppDesc string) (app *cli.Cli, err error) {
	instance := cli.App(cliAppName, cliAppDesc)
	for _, cmd := range commands {
		instance.Command(cmd.Name, cmd.Desc, cmd.CmdInitializer)
	}
	err = instance.Run(args)
	return
}
