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
		_, err := exec.Execute(ctx, "pwd", "-L")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		err = <-ctx.Done()
		switch {
		case err == exec.ErrContextTimeout:
			fmt.Println("timeout!")
		case err != nil:
			fmt.Println(err.Error())
		default:
			fmt.Println("pwd ok")
		}
	}
}

func pwdPipeHandler(cmd *cli.Cmd) {
	cmd.Action = func() {
		ctx := exec.NewExecutionContext()
		ctx = exec.WithIOPipe(ctx)
		_, err := exec.Execute(ctx, "pwd", "-L")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		err = <-ctx.Done()
		switch {
		case err == exec.ErrContextTimeout:
			fmt.Println("timeout!")
		case err != nil:
			fmt.Println(err.Error())
		default:
			fmt.Println("pwd ok!")
		}
	}
}

func pwdIOHandler(cmd *cli.Cmd) {
	cmd.Action = func() {
		ctx := exec.NewExecutionContext()
		ctx = exec.WithTimeout(ctx, 1*time.Second)
		_, err := exec.Execute(ctx, "ping", "8.8.8.8", "-t", "100")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		err = <-ctx.Done()
		switch {
		case err == exec.ErrContextTimeout:
			fmt.Println("timeout!")
		case err != nil:
			fmt.Println(err.Error())
		default:
			fmt.Println("pwd ok!")
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
	Register("pwdPipe", "print working directory", pwdHandler)
	Start([]string{"testapp", "pwdPipe"}, "testapp", "testing cli app")
	//Output:
	//pwd ok
}

func ExampleRegisterOsCommandWithTimeout() {
	Register("pwdTimeout", "print working directory with timeout", pwdIOHandler)
	Start([]string{"testapp", "pwdTimeout"}, "testapp", "testing cli app")
	//Output:
	//timeout!
}
