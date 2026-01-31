package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zkep/cliz"
)

func main() {
	app := cliz.NewCli("command_inheritance", "CLI application demonstrating command inheritance", "1.0.0")

	var globalName string
	var globalAge int
	var serverName string
	var serverPort int
	var userName string

	app.String("name", "Global name flag", &globalName, cliz.Required())
	app.Int("age", "Global age flag", &globalAge, cliz.Range(18, 99))
	app.Action(rootAction(&globalName, &globalAge))

	serverCmd := app.NewSubCommandInheritFlags("server", "Server management (inherits flags)")
	serverCmd.String("server-name", "Server name", &serverName, cliz.Required())
	serverCmd.Int("port", "Server port", &serverPort, cliz.Range(1024, 65535))

	startCmd := serverCmd.NewSubCommand("start", "Start server")
	startCmd.Action(serverStartAction(&globalName, &globalAge, &serverName, &serverPort))

	stopCmd := serverCmd.NewSubCommand("stop", "Stop server")
	stopCmd.Action(serverStopAction(&serverName, &globalName))

	userCmd := app.NewSubCommand("user", "User management (no inheritance)")
	userCmd.String("name", "User name", &userName, cliz.Required())

	createUserCmd := userCmd.NewSubCommand("create", "Create user")
	createUserCmd.Action(createUserAction(&userName))

	if err := app.Run(os.Args...); err != nil {
		log.Fatal(err)
	}
}

func rootAction(globalName *string, globalAge *int) func() error {
	return func() error {
		fmt.Println("=== Root Command ===")
		fmt.Printf("Name: %s\n", *globalName)
		fmt.Printf("Age: %d\n", *globalAge)
		fmt.Println()
		fmt.Println("This is the root command with global flags.")
		fmt.Println("Subcommands can inherit these flags.")
		return nil
	}
}

func serverStartAction(globalName *string, globalAge *int, serverName *string, serverPort *int) func() error {
	return func() error {
		fmt.Println("=== Start Server ===")
		fmt.Printf("Global name: %s\n", *globalName)
		fmt.Printf("Global age: %d\n", *globalAge)
		fmt.Printf("Server name: %s\n", *serverName)
		fmt.Printf("Port: %d\n", *serverPort)
		fmt.Println()
		fmt.Println("This subcommand inherited flags from the parent command.")
		fmt.Println("It adds its own flags while still having access to inherited flags.")
		return nil
	}
}

func serverStopAction(serverName *string, globalName *string) func() error {
	return func() error {
		fmt.Println("=== Stop Server ===")
		fmt.Printf("Server name: %s\n", *serverName)
		fmt.Printf("Global name: %s\n", *globalName)
		fmt.Println()
		fmt.Println("Stopping server...")
		fmt.Println("Server stopped successfully!")
		return nil
	}
}

func createUserAction(userName *string) func() error {
	return func() error {
		fmt.Println("=== Create User ===")
		fmt.Printf("User name: %s\n", *userName)
		fmt.Println()
		fmt.Println("This subcommand does not inherit parent flags.")
		fmt.Println("It defines its own flags independently.")
		return nil
	}
}
