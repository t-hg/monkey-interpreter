package main

import (
	"flag"
	"fmt"
	"os"

	. "github.com/t-hg/monkey-interpreter/repl"
)

func printUsage() {
  fmt.Printf(`Usage: monkey-interpreter COMMAND

Commands:
  repl

`)
  os.Exit(1)
}

func checkForNoMoreArgs(index int) {
  if len(flag.Args()) > index + 1 {
    fmt.Printf("Unknown command '%s'.\n", flag.Arg(index + 1))
    printUsage()
  }
}

func main() {
  flag.Parse()
  if len(flag.Args()) == 0 {
    printUsage()
  }
  switch flag.Arg(0) {
  case "repl":
    checkForNoMoreArgs(0)
    StartRepl()
  default:
    printUsage()
  }
}
