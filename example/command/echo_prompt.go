package command

import (
	"fmt"
	"os"
	"regexp"

	"github.com/jawher/mow.cli"
	"github.com/ncrypthic/gocli"
	"github.com/ncrypthic/gocli/form"
	"github.com/ncrypthic/gocli/form/field"
)

type echoMessage struct {
	name  string
	email string
}

type emailValidator struct{}

func (v *emailValidator) Validate(d []byte) (valid bool, msg string) {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if valid = re.Match(d); !valid {
		msg = fmt.Sprintf("`%s` is not a valid email\n", d)
	}
	return valid, msg
}

func promptHandler(cmd *cli.Cmd) {
	cmd.Action = func() {
		f := form.NewForm()
		m := echoMessage{}
		fields := []field.Field{
			field.NewString(&m.name, "name", "Enter your name: "),
			field.WithValidator(field.NewString(&m.email, "email", "Enter your email: "), &emailValidator{}),
		}
		f.Prompt(os.Stdout, os.Stdin, fields)
		fmt.Printf("%v\n", m)
	}
}

func init() {
	gocli.Register("echoPrompt", "echoed message", promptHandler)
}
