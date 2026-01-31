package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zkep/cliz"
)

func main() {
	app := cliz.NewCli("subcommands_example", "CLI application with subcommands", "1.0.0")

	var name string
	var age int

	app.String("name", "Your name", &name)
	app.Int("age", "Your age", &age)

	app.Action(func() error {
		fmt.Printf("Hello, %s!\n", name)
		fmt.Printf("You are %d years old.\n", age)
		return nil
	})

	greet := app.NewSubCommand("greet", "Greet someone")

	greet.Action(func() error {
		fmt.Printf("Hello, %s! Welcome to the subcommand!\n", name)
		return nil
	})

	farewell := app.NewSubCommand("farewell", "Say goodbye")

	farewell.Action(func() error {
		fmt.Printf("Goodbye, %s! See you later!\n", name)
		return nil
	})

	app.AddCommand(greet)
	app.AddCommand(farewell)

	if err := app.Run(os.Args...); err != nil {
		log.Fatal(err)
	}
}
