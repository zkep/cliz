package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zkep/cliz"
)

func main() {
	app := cliz.NewCli("basic_example", "A basic CLI application", "1.0.0")

	var name string
	var age int
	var verbose bool

	app.String("name", "Your name", &name, cliz.Required().WithMessage("Name cannot be empty"))
	app.Int("age", "Your age", &age, cliz.Range(18, 99).WithMessage("Age must be between %v and %v"))
	app.Bool("verbose", "Enable verbose output", &verbose)

	app.Action(func() error {
		if verbose {
			fmt.Println("=== Verbose Output ===")
		}

		fmt.Printf("Hello, %s!\n", name)
		fmt.Printf("You are %d years old.\n", age)

		if verbose {
			fmt.Println("=== End of Output ===")
		}

		return nil
	})

	if err := app.Run(os.Args...); err != nil {
		log.Fatal(err)
	}
}
