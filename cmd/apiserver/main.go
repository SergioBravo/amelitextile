package main

import (
	"flag"

	"github.com/SergioBravo/amelitextile/cmd/apiserver/cmd"
)

func main() {
	flag.Parse()
	cmd.Execute()
}
