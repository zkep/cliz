package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zkep/cliz"
)

type PositionalArgs struct {
	Source      string `position:"0"`
	Destination string `position:"1"`
	Mode        string `position:"2"`
}

func main() {
	app := cliz.NewCli("positional_struct", "CLI application demonstrating positional arguments with struct", "1.0.0")

	var args PositionalArgs
	var force bool
	var verbose bool

	app.AddPositionalArgs(&args)
	app.Bool("force", "Force operation", &force)
	app.Bool("verbose", "Enable verbose output", &verbose)

	app.Action(actionCallback(&args, &force, &verbose))

	if err := app.Run(os.Args...); err != nil {
		log.Fatal(err)
	}
}

func actionCallback(args *PositionalArgs, force, verbose *bool) func() error {
	return func() error {
		if *verbose {
			fmt.Println("=== Verbose Output ===")
			fmt.Printf("Source: %s\n", args.Source)
			fmt.Printf("Destination: %s\n", args.Destination)
			fmt.Printf("Mode: %s\n", args.Mode)
			fmt.Printf("Force: %t\n", *force)
			fmt.Println()
		}

		fmt.Printf("Operation: %s -> %s\n", args.Source, args.Destination)
		fmt.Printf("Mode: %s\n", args.Mode)

		if *force {
			fmt.Println("Forcing operation...")
		}

		fmt.Println()
		fmt.Println("Positional arguments were parsed from the command line.")
		fmt.Println("This demonstrates how to use struct tags to capture positional arguments.")
		fmt.Println()
		fmt.Println("Example usage:")
		fmt.Println("  ./positional_struct source.txt destination.txt copy --force --verbose")
		fmt.Println("  ./positional_struct ./files ./backup archive --force")
		fmt.Println("  ./positional_struct file1.jpg new_name.jpg rename")
		return nil
	}
}
