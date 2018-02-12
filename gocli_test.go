package gocli

import (
	"fmt"
	"time"

	"github.com/jawher/mow.cli"
	"github.com/ncrypthic/gocli/exec"
)

func echoHandler(cmd *cli.Cmd) {
	cmd.Spec = "ARG"
	msg := cmd.String(cli.StringArg{Name: "ARG", Value: "", Desc: "Echo string"})
	cmd.Action = func() {
		if msg != nil {
			fmt.Println(*msg)
		}
	}
}

func pwdHandler(cmd *cli.Cmd) {
	cmd.Action = func() {
		ctx := exec.NewExecutionContext()
		ctx = exec.WithTimeout(ctx, 3*time.Second)
		if _, err := exec.Execute(ctx, "pwd", "-L"); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("pwd ok")
		}
	}
}

func pwdIOHandler(cmd *cli.Cmd) {
	cmd.Action = func() {
		ctx := exec.NewExecutionContext()
		ctx = exec.WithTimeout(ctx, 3*time.Second)
		ctx = exec.WithIOPipe(ctx)
		if _, err := exec.Execute(ctx, "pwd", "-L"); err != nil {
			fmt.Println(err.Error())
		}
	}
}

func ExampleRegisterCustomCommand() {
	Register("echo", "echo command", echoHandler)
	Start([]string{"testapp", "echo", "test echo"}, "testapp", "testing cli app")
	//Output:
	//test echo
}

func ExampleRegisterOsCommand() {
	Register("pwd", "print working directory", pwdHandler)
	Start([]string{"testapp", "pwd"}, "testapp", "testing cli app")
	//Output:
	//pwd ok
}

func ExampleRegisterOsCommandWithIOPipe() {
	Register("pwd", "print working directory", pwdHandler)
	Start([]string{"testapp", "pwd"}, "testapp", "testing cli app")
	//Output:
	//pwd ok
}
