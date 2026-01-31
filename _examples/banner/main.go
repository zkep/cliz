package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zkep/cliz"
)

func main() {
	app := cliz.NewCli("custom_banner", "CLI application with custom banner and error handler", "1.0.0")

	app.SetBannerFunction(customBanner)

	app.SetErrorFunction(customErrorHandler)

	var name string
	var age int
	var verbose bool

	app.String("name", "Your name", &name, cliz.Required())
	app.Int("age", "Your age", &age, cliz.Range(0, 120))
	app.Bool("verbose", "Enable verbose output", &verbose)

	app.Action(func() error {
		if verbose {
			fmt.Println("\n==================================================")
			fmt.Println("Verbose Output Enabled")
			fmt.Println()
		}

		fmt.Printf("Hello, %s!\n", name)
		fmt.Printf("You are %d years old.\n", age)

		if verbose {
			fmt.Println()
			fmt.Println("Additional verbose information:")
			fmt.Printf("- Name length: %d characters\n", len(name))
			fmt.Printf("- Birth year: %d\n", 2024-age)
			fmt.Println()
			fmt.Println("==================================================")
		}

		return nil
	})

	if err := app.Run(os.Args...); err != nil {
		log.Fatal(err)
	}
}

func customBanner(c *cliz.Cli) string {
	return fmt.Sprintf(`
╔══════════════════════════════════════════════════════════╗
║                                                          ║
║   %-50s ║
║                                                          ║
║   Version: %-46v ║
║                                                          ║
║   Description: %-41v ║
║                                                          ║
╚══════════════════════════════════════════════════════════╝`,
		c.Name(),
		c.Version(),
		c.ShortDescription())
}

func customErrorHandler(commandPath string, err error) error {
	return fmt.Errorf("\n❌  Error in command '%s':\n%s\n\n💡  Please check your input and try again.\n", commandPath, err)
}
