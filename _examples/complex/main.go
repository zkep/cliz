package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zkep/cliz"
)

func main() {
	app := cliz.NewCli("complex_subcommands", "CLI application with complex subcommand hierarchy", "1.0.0")

	var globalName string
	var globalAge int
	var verbose bool

	// Global flags
	app.String("name", "Your name", &globalName, cliz.Required())
	app.Int("age", "Your age", &globalAge, cliz.Range(18, 99))
	app.Bool("verbose", "Enable verbose output", &verbose)

	app.Action func() error {
		fmt.Println("=== Root Command ===")
		fmt.Printf("Hello, %s!\n", globalName)
		fmt.Printf("You are %d years old.\n", globalAge)
		fmt.Println()
		fmt.Println("Available commands: user, file, server, database")
		if verbose {
			fmt.Println("\nVerbose mode enabled.")
			fmt.Println("Root command executed successfully.")
		}
		return nil
	}

	// First level subcommands
	userCmd := app.NewSubCommand("user", "User management commands")
	fileCmd := app.NewSubCommand("file", "File system commands")
	serverCmd := app.NewSubCommand("server", "Server management commands")
	databaseCmd := app.NewSubCommand("database", "Database commands")

	// User command subcommands
	var username string
	var email string
	userCmd.String("username", "Username", &username, cliz.Required())
	userCmd.String("email", "Email address", &email, cliz.Email())

	createUserCmd := userCmd.NewSubCommand("create", "Create a new user")
	createUserCmd.Action func() error {
		fmt.Println("=== Create User ===")
		fmt.Printf("Creating user: %s\n", username)
		fmt.Printf("Email: %s\n", email)
		fmt.Println("User created successfully!")
		if verbose {
			fmt.Println("\nVerbose: User creation process completed in 1.2s")
		}
		return nil
	}

	deleteUserCmd := userCmd.NewSubCommand("delete", "Delete an existing user")
	deleteUserCmd.Action func() error {
		fmt.Println("=== Delete User ===")
		fmt.Printf("Deleting user: %s\n", username)
		fmt.Println("User deleted successfully!")
		if verbose {
			fmt.Println("\nVerbose: User deletion process completed in 0.8s")
		}
		return nil
	}

	// File command subcommands
	var filePath string
	fileCmd.String("path", "File path", &filePath, cliz.Required())

	readFileCmd := fileCmd.NewSubCommand("read", "Read file content")
	readFileCmd.Action func() error {
		fmt.Println("=== Read File ===")
		fmt.Printf("Reading file: %s\n", filePath)
		fmt.Println("File content (simulated):")
		fmt.Println("Lorem ipsum dolor sit amet, consectetur adipiscing elit.")
		if verbose {
			fmt.Println("\nVerbose: File read successfully, 1024 bytes loaded")
		}
		return nil
	}

	writeFileCmd := fileCmd.NewSubCommand("write", "Write content to file")
	var content string
	writeFileCmd.String("content", "Content to write", &content, cliz.Required())
	writeFileCmd.Action func() error {
		fmt.Println("=== Write File ===")
		fmt.Printf("Writing to file: %s\n", filePath)
		fmt.Printf("Content: %s\n", content)
		fmt.Println("File written successfully!")
		if verbose {
			fmt.Println("\nVerbose: File write completed, 156 bytes written")
		}
		return nil
	}

	// Second level subcommands
	copyFileCmd := fileCmd.NewSubCommand("copy", "Copy files")
	copyFileCmd.String("source", "Source path", &filePath, cliz.Required())
	var destPath string
	copyFileCmd.String("destination", "Destination path", &destPath, cliz.Required())

	copyLocalCmd := copyFileCmd.NewSubCommand("local", "Copy local files")
	copyLocalCmd.Action func() error {
		fmt.Println("=== Copy Local File ===")
		fmt.Printf("From: %s\n", filePath)
		fmt.Printf("To: %s\n", destPath)
		fmt.Println("File copied locally successfully!")
		if verbose {
			fmt.Println("\nVerbose: Local file copy completed in 2.5s")
		}
		return nil
	}

	copyRemoteCmd := copyFileCmd.NewSubCommand("remote", "Copy to/from remote server")
	var host string
	copyRemoteCmd.String("host", "Remote host", &host, cliz.Required())
	copyRemoteCmd.Action func() error {
		fmt.Println("=== Copy Remote File ===")
		fmt.Printf("From: %s\n", filePath)
		fmt.Printf("To: %s@%s\n", destPath, host)
		fmt.Println("File copied to remote server successfully!")
		if verbose {
			fmt.Println("\nVerbose: Remote file copy completed in 5.8s")
		}
		return nil
	}

	// Server command subcommands
	serverStartCmd := serverCmd.NewSubCommand("start", "Start server")
	var serverPort int
	serverStartCmd.Int("port", "Server port", &serverPort, cliz.Range(1024, 65535))
	serverStartCmd.Action func() error {
		fmt.Println("=== Start Server ===")
		fmt.Printf("Starting server on port %d...\n", serverPort)
		fmt.Println("Server started successfully!")
		if verbose {
			fmt.Println("\nVerbose: Server process PID: 12345")
			fmt.Println("Verbose: Listening on all interfaces")
		}
		return nil
	}

	serverStopCmd := serverCmd.NewSubCommand("stop", "Stop server")
	serverStopCmd.Action func() error {
		fmt.Println("=== Stop Server ===")
		fmt.Println("Stopping server...")
		fmt.Println("Server stopped successfully!")
		if verbose {
			fmt.Println("\nVerbose: Server process terminated")
		}
		return nil
	}

	serverStatusCmd := serverCmd.NewSubCommand("status", "Check server status")
	serverStatusCmd.Action func() error {
		fmt.Println("=== Server Status ===")
		fmt.Println("Status: Running")
		fmt.Println("Uptime: 5 days, 3 hours, 15 minutes")
		fmt.Println("CPU Usage: 12.5%")
		fmt.Println("Memory Usage: 512 MB")
		if verbose {
			fmt.Println("\nVerbose: Status check completed")
		}
		return nil
	}

	// Database command subcommands
	dbConnectCmd := databaseCmd.NewSubCommand("connect", "Connect to database")
	var dbHost string
	var dbUser string
	dbConnectCmd.String("host", "Database host", &dbHost)
	dbConnectCmd.String("user", "Database user", &dbUser)
	dbConnectCmd.Action func() error {
		fmt.Println("=== Connect to Database ===")
		fmt.Printf("Connecting to %s as %s...\n", dbHost, dbUser)
		fmt.Println("Connected successfully!")
		if verbose {
			fmt.Println("\nVerbose: Connection established, session ID: abc123")
		}
		return nil
	}

	dbDisconnectCmd := databaseCmd.NewSubCommand("disconnect", "Disconnect from database")
	dbDisconnectCmd.Action func() error {
		fmt.Println("=== Disconnect from Database ===")
		fmt.Println("Disconnecting from database...")
		fmt.Println("Disconnected successfully!")
		if verbose {
			fmt.Println("\nVerbose: Connection closed")
		}
		return nil
	}

	if err := app.Run(os.Args...); err != nil {
		log.Fatal(err)
	}
}
