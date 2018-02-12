package command

import (
	"fmt"

	"github.com/jawher/mow.cli"
	"github.com/ncrypthic/gocli"
)

func handler(cmd *cli.Cmd) {
	cmd.Spec = "MESSAGE"
	message := cmd.String(cli.StringArg{Name: "MESSAGE", Value: "", Desc: "Echo string"})
	cmd.Action = func() {
		fmt.Println(*message)
	}
}

func init() {
	gocli.Register("echo", "echoed message", handler)
}
