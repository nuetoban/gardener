package main

import (
	"fmt"
	"os"

	"github.com/nuetoban/gardener/fs/local"
	"github.com/nuetoban/gardener/hotbed"

	"github.com/akamensky/argparse"
)

func main() {
	var err error

	// Create new parser object
	parser := argparse.NewParser("gardener", "Creates directories tree from provided yaml")
	// Create string flag
	s := parser.String("y", "yaml", &argparse.Options{Required: true, Help: "Path to .yaml file"})
	// Parse input
	err = parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
	}

	gardener := hotbed.New(local.FS{})

	file, err := os.Open(*s)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = gardener.ParseYAML(file)
	if err != nil {
		panic(err)
	}

	err = gardener.Create()
	if err != nil {
		panic(err)
	}
}
