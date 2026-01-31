package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zkep/cliz"
)

func main() {
	app := cliz.NewCli("struct_flags_example", "CLI application using struct flags", "1.0.0")

	type Config struct {
		Name     string  `name:"name" description:"Your name" validate:"required,error_required=Name cannot be empty"`
		Age      int     `name:"age" description:"Your age" validate:"required,range=18-99,error_range=Age must be between 18 and 99"`
		Email    string  `name:"email" description:"Your email" validate:"required,error_required=Email cannot be empty,email,error_email=Please enter a valid email address"`
		Website  string  `name:"website" description:"Your website" validate:"url,error_url=Please enter a valid website URL"`
		Verbose  bool    `name:"verbose" description:"Enable verbose output"`
		Score    float64 `name:"score" description:"Your score" validate:"range=0-100,error_range=Score must be between 0 and 100"`
		Status   string  `name:"status" description:"Your status (active/inactive)" validate:"in=active|inactive,error_in=Status must be active or inactive"`
		Password string  `name:"password" description:"Your password" validate:"required,pattern=^(?=.*[a-z])(?=.*[A-Z])(?=.*\\d)[a-zA-Z\\d]{8,}$,error_pattern=Password must contain at least one lowercase letter, one uppercase letter, and one number, and be at least 8 characters long"`
	}

	var config Config

	app.AddFlags(&config)

	app.Action(func() error {
		if config.Verbose {
			fmt.Println("=== Configuration ===")
		}

		fmt.Printf("Name: %s\n", config.Name)
		fmt.Printf("Age: %d\n", config.Age)
		fmt.Printf("Email: %s\n", config.Email)
		fmt.Printf("Website: %s\n", config.Website)
		fmt.Printf("Verbose: %t\n", config.Verbose)
		fmt.Printf("Score: %.2f\n", config.Score)
		fmt.Printf("Status: %s\n", config.Status)
		fmt.Printf("Password: %s\n", config.Password)

		if config.Verbose {
			fmt.Println("=== End of Output ===")
		}

		return nil
	})

	if err := app.Run(os.Args...); err != nil {
		log.Fatal(err)
	}
}
