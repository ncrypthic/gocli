package main

import (
	"os"

	"github.com/ncrypthic/gocli"
	_ "github.com/ncrypthic/gocli/example/command"
)

func main() {
	gocli.Start(os.Args, "cliapp", "example")
}
