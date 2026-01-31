package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zkep/cliz"
)

type Config struct {
	Name    string  `name:"name" description:"Your name" default:"Guest"`
	Age     int     `name:"age" description:"Your age" default:"25"`
	Email   string  `name:"email" description:"Your email" default:"user@example.com"`
	Port    int     `name:"port" description:"Server port" default:"8080"`
	Debug   bool    `name:"debug" description:"Enable debug mode" default:"false"`
	Verbose bool    `name:"verbose" description:"Enable verbose output" default:"false"`
	Limit   int     `name:"limit" description:"Limit value" default:"100"`
	Timeout float64 `name:"timeout" description:"Timeout in seconds" default:"30.0"`
}

func main() {
	app := cliz.NewCli("flag_struct_tags", "CLI application demonstrating flag struct tags", "1.0.0")

	var config Config
	var customFlag string

	app.AddFlags(&config)
	app.String("custom", "Custom flag not in struct", &customFlag)
	app.Action(actionCallback(&config, &customFlag))

	if err := app.Run(os.Args...); err != nil {
		log.Fatal(err)
	}
}

func actionCallback(config *Config, customFlag *string) func() error {
	return func() error {
		fmt.Println("=== Flag Struct Tags Demonstration ===")
		fmt.Println()
		fmt.Println("Flags parsed from struct with tags:")
		fmt.Printf("Name: %s\n", config.Name)
		fmt.Printf("Age: %d\n", config.Age)
		fmt.Printf("Email: %s\n", config.Email)
		fmt.Printf("Port: %d\n", config.Port)
		fmt.Printf("Debug: %t\n", config.Debug)
		fmt.Printf("Verbose: %t\n", config.Verbose)
		fmt.Printf("Limit: %d\n", config.Limit)
		fmt.Printf("Timeout: %f\n", config.Timeout)
		fmt.Println()
		fmt.Printf("Custom flag: %s\n", *customFlag)
		fmt.Println()
		fmt.Println("This demonstrates how to use struct fields with tags to define flags.")
		fmt.Println("Tags supported: name, description, default")
		fmt.Println()
		fmt.Println("Example usage:")
		fmt.Println("  ./flag_struct_tags --name John --age 30 --debug")
		fmt.Println("  ./flag_struct_tags --port 9090 --timeout 60.0")
		fmt.Println("  ./flag_struct_tags --custom value --verbose")
		return nil
	}
}
