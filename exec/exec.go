package exec

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

//ErrContextTimeout indicated context cancelled due to timed out
type ErrContextTimeout error

//ErrContextCancelled indicated context cancelled
type ErrContextCancelled error

//EnvironmentVariables is a key-value string pairs for environment variables
type EnvironmentVariables map[string]string

//ToSliceString join each key-value string pair in EnvironmentVariables and returned it as string slice
func (envar EnvironmentVariables) ToSliceString() []string {
	variables := os.Environ()
	for k, v := range envar {
		if k[0] == '+' {
			actualVariableName := k[1:]
			existingIndex := -1
			for i, osVar := range variables {
				if osVar[0:len(actualVariableName)] == actualVariableName {
					existingIndex = i
				}
			}
			if existingIndex > -1 {
				variables[existingIndex] = variables[existingIndex] + ";" + v
			} else {
				variables = append(variables, fmt.Sprintf(`%s=%s`, actualVariableName, v))
			}
		} else {
			variables = append(variables, fmt.Sprintf(`%s=%s`, k, v))
		}
	}
	return variables
}

//ExecutionContext holds context for os command execution
type ExecutionContext struct {
	context.Context
	vars    EnvironmentVariables
	wd      string
	pipedIO bool
	chDone  chan error
	timeout time.Duration
}

//NewExecutionContext creates default execution context
func NewExecutionContext() ExecutionContext {
	chDone := make(chan error)
	return ExecutionContext{chDone: chDone}
}

//Done returns an error channel when a context is done
func (ctx ExecutionContext) Done() <-chan error {
	return ctx.chDone
}

//WithCancel returns copy ExecutionContext with a cancel function to cancel the execution context
func WithCancel(ctx ExecutionContext) (ExecutionContext, func()) {
	cancelFunc := func() {
		ctx.chDone <- ErrContextCancelled(fmt.Errorf("Context cancelled"))
	}
	return ExecutionContext{
		vars:    ctx.vars,
		wd:      ctx.wd,
		chDone:  ctx.chDone,
		timeout: ctx.timeout,
	}, cancelFunc
}

//WithTimeout adds a timeout duration for an execution context. When the context runs longer than timeout
// duration, it will be cancelled with ErrContextTimeout passed to the context `Done` method
func WithTimeout(ec ExecutionContext, d time.Duration) ExecutionContext {
	return ExecutionContext{
		vars:    ec.vars,
		wd:      ec.wd,
		chDone:  ec.chDone,
		timeout: d,
	}
}

//WithEnvironmentVariables append specified environment variables
func WithEnvironmentVariables(ec ExecutionContext, vars EnvironmentVariables) ExecutionContext {
	return ExecutionContext{
		vars:    vars,
		wd:      ec.wd,
		chDone:  ec.chDone,
		timeout: ec.timeout,
	}
}

//WithWorkingDirectory sets working directory for process that will be executed
func WithWorkingDirectory(ec ExecutionContext, workdir string) ExecutionContext {
	return ExecutionContext{
		vars:    ec.vars,
		wd:      workdir,
		chDone:  ec.chDone,
		timeout: ec.timeout,
	}
}

//WithIOPipe will execute context with io piped
func WithIOPipe(ec ExecutionContext) ExecutionContext {
	return ExecutionContext{
		vars:    ec.vars,
		wd:      ec.wd,
		chDone:  ec.chDone,
		timeout: ec.timeout,
		pipedIO: true,
	}
}

//Execute will run command within specified context
func Execute(ctx ExecutionContext, cmd string, args ...string) (proc *exec.Cmd, err error) {
	defer close(ctx.chDone)
	proc = exec.Command(cmd, args...)
	if len(ctx.vars) > 0 {
		proc.Env = ctx.vars.ToSliceString()
	}
	if ctx.wd != "" {
		proc.Dir = ctx.wd
	}
	if ctx.timeout > 0 {
		go func() {
			for {
				select {
				case <-time.After(ctx.timeout):
					ctx.chDone <- ErrContextTimeout(fmt.Errorf("Context timed out"))
					if !proc.ProcessState.Exited() {
						proc.Process.Kill()
					}
					return
				case err := <-ctx.chDone:
					switch err.(type) {
					case ErrContextCancelled:
						if !proc.ProcessState.Exited() {
							proc.Process.Kill()
						}
					}
					return
				}
			}
		}()
	}
	if ctx.pipedIO {
		proc.Stdout = os.Stdout
		proc.Stdin = os.Stdin
		proc.Stderr = os.Stderr
	}
	// fmt.Println("Process envar:")
	// fmt.Printf("%s\n", strings.Join(proc.Env, "\n"))
	if err = proc.Start(); err != nil {
		fmt.Printf("The command `%s %s` failed to run, please make sure you are running command with correct arguments\n", cmd, strings.Join(args, " "))
	}
	err = proc.Wait()
	return
}
