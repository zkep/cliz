package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zkep/cliz"
)

func main() {
	app := cliz.NewCli("hidden_default_commands", "CLI application with hidden and default commands", "1.0.0")

	var name string
	var verbose bool

	app.String("name", "Your name", &name, cliz.Required())
	app.Bool("verbose", "Enable verbose output", &verbose)

	defaultCmd := cliz.NewCommand("default", "Default command when no command is specified")
	defaultCmd.Action(defaultCommandAction(name))

	hiddenCmd := cliz.NewCommand("secret", "Hidden command (not shown in help)")
	hiddenCmd.Hidden(true)
	hiddenCmd.Action(hiddenCommandAction(name))

	greetCmd := app.NewSubCommand("greet", "Greet someone")
	greetCmd.Action(greetCommandAction(name, verbose))

	farewellCmd := app.NewSubCommand("farewell", "Say goodbye")
	farewellCmd.Action(farewellCommandAction(name, verbose))

	app.DefaultCommand(defaultCmd)

	app.AddCommand(defaultCmd)
	app.AddCommand(hiddenCmd)

	if err := app.Run(os.Args...); err != nil {
		log.Fatal(err)
	}
}

func defaultCommandAction(name string) func() error {
	return func() error {
		fmt.Println("=== Default Command ===")
		fmt.Printf("Hello, %s!\n", name)
		fmt.Println()
		fmt.Println("Note: This is the default command that runs when no command is specified.")
		fmt.Println("Use --help to see available commands.")
		return nil
	}
}

func hiddenCommandAction(name string) func() error {
	return func() error {
		fmt.Println("=== Hidden Secret Command ===")
		fmt.Printf("Welcome to the secret command, %s!\n", name)
		fmt.Println()
		fmt.Println("This command is hidden from help output.")
		fmt.Println("You can only use it if you know it exists.")
		return nil
	}
}

func greetCommandAction(name string, verbose bool) func() error {
	return func() error {
		fmt.Println("=== Greet Command ===")
		fmt.Printf("Hello, %s!\n", name)
		if verbose {
			fmt.Println("Welcome! Hope you're having a great day.")
		}
		return nil
	}
}

func farewellCommandAction(name string, verbose bool) func() error {
	return func() error {
		fmt.Println("=== Farewell Command ===")
		fmt.Printf("Goodbye, %s!\n", name)
		if verbose {
			fmt.Println("See you later! Have a wonderful day.")
		}
		return nil
	}
}
