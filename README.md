gocli
=====

`gocli` extends [github.com/jawher/mow.cli](https://github.com/jawher/mow.cli) command line application
framework with common utility functions to help developers creates full featured command line application.

Basic Usage
-----------

1. Create a [CommandHandler](https://godoc.org/github.com/ncrypthic/gocli#CommandHandler)

2. [Register](https://godoc.org/github.com/ncrypthic/gocli#Register) the command to gocli

3. [Start your cli](https://godoc.org/github.com/ncrypthic/gocli#Start)

For real example, see [example](example/)

gocli packages
--------------

| Package |               Desc               |
----------------------------------------------
| exec    | Sub processes management package |
| form    | Cli form helper                  |
| utils   | Interaction package              |
