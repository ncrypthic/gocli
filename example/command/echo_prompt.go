package command

import (
	"fmt"
	"os"

	"github.com/jawher/mow.cli"
	"github.com/ncrypthic/gocli"
	"github.com/ncrypthic/gocli/form"
	"github.com/ncrypthic/gocli/form/field"
)

type echoMessage struct {
	name  string
	email string
}

func promptHandler(cmd *cli.Cmd) {
	cmd.Action = func() {
		f := form.NewForm()
		m := echoMessage{}
		fields := []field.Field{
			field.NewString(&m.name, "name", "Enter your name: "),
			field.NewString(&m.email, "email", "Enter your email: "),
		}
		f.Prompt(os.Stdout, os.Stdin, fields)
		fmt.Printf("%v\n", m)
	}
}

func init() {
	gocli.Register("echoPrompt", "echoed message", promptHandler)
}
