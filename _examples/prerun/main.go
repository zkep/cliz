package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zkep/cliz"
)

func main() {
	app := cliz.NewCli("pre_run_validation", "CLI application with PreRun callback and validation chain", "1.0.0")

	var name string
	var age int
	var email string
	var password string
	var port int
	var username string

	app.PreRun(preRunCallback)

	app.String("name", "Your name", &name,
		cliz.Required(),
		cliz.Pattern(`^.{3,20}$`).WithMessage("Name must be between 3-20 characters"),
	)

	app.Int("age", "Your age", &age,
		cliz.Required(),
		cliz.Range(18, 99).WithMessage("Age must be between 18-99"),
	)

	app.String("email", "Your email address", &email,
		cliz.Required(),
		cliz.Email(),
	)

	app.String("password", "Your password", &password,
		cliz.Required(),
		cliz.Pattern(`^.{8,}$`).WithMessage("Password must be at least 8 characters"),
		cliz.Pattern(`[A-Za-z0-9@#$%^&+=]{8,}`).WithMessage("Password must contain uppercase, lowercase, numbers and special characters"),
	)

	app.Int("port", "Server port", &port,
		cliz.Required(),
		cliz.Range(1024, 65535).WithMessage("Port must be between 1024-65535"),
	)

	app.String("username", "Your username", &username,
		cliz.Required(),
		cliz.Alphanum(),
		cliz.Len(8),
		cliz.Contains("user").WithMessage("Username must contain 'user'"),
	)

	app.Action(func() error {
		fmt.Println("=== Action Callback ===")
		fmt.Println()
		fmt.Println("All validations passed successfully!")
		fmt.Println()
		fmt.Println("User information:")
		fmt.Printf("Name: %s\n", name)
		fmt.Printf("Age: %d\n", age)
		fmt.Printf("Email: %s\n", email)
		fmt.Printf("Username: %s\n", username)
		fmt.Printf("Server port: %d\n", port)
		fmt.Println()
		fmt.Println("Validation chain examples shown:")
		fmt.Println("- Required validation (must provide a value)")
		fmt.Println("- Range validation (value must be between min and max)")
		fmt.Println("- Alphanum validation (letters and numbers only)")
		fmt.Println("- Email validation (valid email format)")
		fmt.Println("- Pattern validation (regular expression for length and content)")
		fmt.Println("- Contains validation (must contain specific substring)")
		return nil
	})

	if err := app.Run(os.Args...); err != nil {
		log.Fatal(err)
	}
}

func preRunCallback(c *cliz.Cli) error {
	fmt.Println("=== PreRun Callback ===")
	fmt.Printf("Application name: %s\n", c.Name())
	fmt.Printf("Version: %s\n", c.Version())
	fmt.Printf("Description: %s\n", c.ShortDescription())
	fmt.Println("Initialization complete. Starting command execution...")
	fmt.Println("========================================")
	fmt.Println()
	return nil
}
