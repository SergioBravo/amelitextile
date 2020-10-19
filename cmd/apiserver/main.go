package main

import (
	"amelitextile/cmd"
	"flag"
)

func main() {
	flag.Parse()
	cmd.Execute()
}
