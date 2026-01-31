package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zkep/cliz"
)

func main() {
	app := cliz.NewCli("advanced_validators", "CLI application demonstrating advanced validators", "1.0.0")

	var phoneNumber string
	var zipCode string
	var username string
	var gender string
	var email string
	var url string
	var password string
	var message string
	var code string
	var usernamePattern string
	var color string
	var country string

	app.String("phone", "Phone number (e.g., +1-555-123-4567)", &phoneNumber,
		cliz.Pattern(`^\+?[1-9]\d{1,14}$`).WithMessage("Phone number format is invalid, should be E.164 format"),
	)

	app.String("zip", "ZIP code (5 digits)", &zipCode,
		cliz.Pattern(`^\d{5}(-\d{4})?$`).WithMessage("ZIP code format is invalid"),
	)

	app.String("gender", "Gender", &gender,
		cliz.In("male", "female", "other").WithMessage("Gender must be male, female or other"),
	)

	app.String("color", "Favorite color", &color,
		cliz.In("red", "blue", "green", "yellow", "purple", "orange").WithMessage("Color must be one of the specified options"),
	)

	app.String("country", "Country", &country,
		cliz.In("USA", "Canada", "UK", "Australia", "Japan", "China", "Germany", "France").WithMessage("Country must be one of the supported options"),
	)

	app.String("message", "Message (must contain 'hello')", &message,
		cliz.Contains("hello").WithMessage("Message must contain 'hello' text"),
	)

	app.String("code", "Code (must contain 'APP')", &code,
		cliz.Contains("APP").WithMessage("Code must contain 'APP'"),
	)

	app.String("email", "Email address", &email,
		cliz.Email().WithMessage("Email format is invalid"),
	)

	app.String("url", "Website URL", &url,
		cliz.URL().WithMessage("URL format is invalid"),
	)

	app.String("password", "Password (complex)", &password,
		cliz.Pattern(`^.{8,}$`).WithMessage("Password must be at least 8 characters"),
		cliz.Pattern(`.*[A-Z].*`).WithMessage("Password must contain uppercase letters"),
		cliz.Pattern(`.*[a-z].*`).WithMessage("Password must contain lowercase letters"),
		cliz.Pattern(`.*\d.*`).WithMessage("Password must contain numbers"),
		cliz.Pattern(`.*[@#$%^&+=].*`).WithMessage("Password must contain special characters"),
	)

	app.String("username", "Username (letters, numbers, underscores, 3-20 characters)", &username,
		cliz.Pattern(`^[a-zA-Z0-9_]{3,20}$`).WithMessage("Username must be 3-20 characters, containing letters, numbers, and underscores"),
	)

	app.String("username-pattern", "Username with pattern", &usernamePattern,
		cliz.Alphanum(),
		cliz.Len(5),
	)

	app.Action(func() error {
		return actionCallback(&phoneNumber, &zipCode, &username, &gender, &email, &url, &password, &message, &code, &usernamePattern, &color, &country)
	})

	if err := app.Run(os.Args...); err != nil {
		log.Fatal(err)
	}
}

func actionCallback(phoneNumber, zipCode, username, gender, email, url, password, message, code, usernamePattern, color, country *string) error {
	fmt.Println("=== Advanced Validators Demonstration ===")
	fmt.Println()
	fmt.Println("All validations passed successfully!")
	fmt.Println()
	fmt.Println("Input values:")
	if *phoneNumber != "" {
		fmt.Printf("Phone: %s\n", *phoneNumber)
	}
	if *zipCode != "" {
		fmt.Printf("ZIP: %s\n", *zipCode)
	}
	if *username != "" {
		fmt.Printf("Username: %s\n", *username)
	}
	if *gender != "" {
		fmt.Printf("Gender: %s\n", *gender)
	}
	if *email != "" {
		fmt.Printf("Email: %s\n", *email)
	}
	if *url != "" {
		fmt.Printf("URL: %s\n", *url)
	}
	if *password != "" {
		fmt.Printf("Password: %s characters\n", len(*password))
	}
	if *message != "" {
		fmt.Printf("Message: %s\n", *message)
	}
	if *code != "" {
		fmt.Printf("Code: %s\n", *code)
	}
	if *usernamePattern != "" {
		fmt.Printf("Username (pattern): %s\n", *usernamePattern)
	}
	if *color != "" {
		fmt.Printf("Color: %s\n", *color)
	}
	if *country != "" {
		fmt.Printf("Country: %s\n", *country)
	}
	fmt.Println()
	fmt.Println("Advanced validators demonstrated:")
	fmt.Println("- Pattern: Regular expression validation")
	fmt.Println("- In: Check if value is in predefined list")
	fmt.Println("- Contains: Check if string contains substring")
	fmt.Println("- Email: Validate email format")
	fmt.Println("- URL: Validate URL format")
	fmt.Println("- Combined validators: Multiple constraints")
	fmt.Println()
	fmt.Println("Example usage:")
	fmt.Println("  ./advanced_validators --phone +1-555-123-4567 --zip 12345 --gender male")
	fmt.Println("  ./advanced_validators --email test@example.com --url https://example.com")
	fmt.Println("  ./advanced_validators --password 'Str0ngP@ssw0rd' --message 'hello world'")
	return nil
}
