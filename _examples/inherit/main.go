package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zkep/cliz"
)

func main() {
	app := cliz.NewCli("inherit_flags_example", "CLI application with flag inheritance", "1.0.0")

	var name string
	var age int
	var verbose bool

	app.String("name", "Your name", &name)
	app.Int("age", "Your age", &age)
	app.Bool("verbose", "Enable verbose output", &verbose)

	app.Action(func() error {
		if verbose {
			fmt.Println("=== Main Command ===")
		}
		fmt.Printf("Hello, %s!\n", name)
		fmt.Printf("You are %d years old.\n", age)
		return nil
	})

	greet := app.NewSubCommandInheritFlags("greet", "Greet someone with inherited flags")

	greet.Action(func() error {
		if verbose {
			fmt.Println("=== Greet Subcommand ===")
		}
		fmt.Printf("Hello, %s! How are you doing?\n", name)
		return nil
	})

	farewell := app.NewSubCommandInheritFlags("farewell", "Say goodbye with inherited flags")

	farewell.Action(func() error {
		if verbose {
			fmt.Println("=== Farewell Subcommand ===")
		}
		fmt.Printf("Goodbye, %s! Take care!\n", name)
		return nil
	})

	app.AddCommand(greet)
	app.AddCommand(farewell)

	if err := app.Run(os.Args...); err != nil {
		log.Fatal(err)
	}
}
