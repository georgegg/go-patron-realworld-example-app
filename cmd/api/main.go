package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/beatlabs/patron"
)

var (
	theVersion = "dev"
)

func main() {
	if err := run(theVersion); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}

func run(version string) error {
	var (
		showVersion = flag.Bool("v", false, "print version number")
	)
	flag.Parse()

	if *showVersion {
		fmt.Printf("%s %s (runtime: %s)\n", os.Args[0], version, runtime.Version())
		os.Exit(0)
	}

	const serviceName = "go-patron-realworld-example-app"

	err := patron.Setup(serviceName, version)
	if err != nil {
		return fmt.Errorf("failed to set up logging: %v", err)
	}

	srv, err := patron.New(serviceName, version)
	if err != nil {
		return fmt.Errorf("failed to create service %v", err)
	}

	err = srv.Run()
	if err != nil {
		return fmt.Errorf("failed to run service %v", err)
	}
	return nil
}
