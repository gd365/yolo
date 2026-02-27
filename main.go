package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yolooks/yolo/pkg/service"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	initCommand := flag.NewFlagSet("init", flag.ExitOnError)
	projectName := initCommand.String("name", "", "Name of the project (required)")
	projectPort := initCommand.Int("port", 8080, "Port of the project")
	verbose := initCommand.Bool("v", false, "Verbose output")

	versionCommand := flag.NewFlagSet("version", flag.ExitOnError)

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		initCommand.Parse(os.Args[2:])
		if *projectName == "" {
			fmt.Println("error: -name is required")
			initCommand.Usage()
			os.Exit(1)
		}
		if *projectPort <= 0 || *projectPort > 65535 {
			fmt.Println("error: -port must be between 1 and 65535")
			os.Exit(1)
		}
		runInit(*projectName, *projectPort, *verbose)
	case "version":
		versionCommand.Parse(os.Args[2:])
		printVersion()
	default:
		printUsage()
		os.Exit(1)
	}
}

func runInit(name string, port int, verbose bool) {
	rander := service.NewRander(name, port, service.WithVerbose(verbose))

	if err := rander.InitDir(); err != nil {
		printError("init directory", err)
		os.Exit(1)
	}

	if err := rander.InitPkg(); err != nil {
		printError("init package", err)
		os.Exit(1)
	}

	if err := rander.RunGoMod(); err != nil {
		printError("init go module", err)
		os.Exit(1)
	}

	fmt.Printf("\n✓ Project '%s' created successfully!\n", name)
	fmt.Printf("  cd %s && go run cmd/server.go\n", name)
}

func printError(context string, err error) {
	fmt.Fprintf(os.Stderr, "error: %s failed\n  %v\n", context, err)
}

func printVersion() {
	fmt.Printf("yolo %s (commit: %s, built: %s)\n", version, commit, date)
}

func printUsage() {
	fmt.Println("yolo - Golang project scaffolding")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  yolo init -name <project> [-port <port>] [-v]")
	fmt.Println("  yolo version")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  init     Create a new project")
	fmt.Println("  version  Show version info")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -name    Project name (required)")
	fmt.Println("  -port    Project port (default: 8080)")
	fmt.Println("  -v       Verbose output")
}
