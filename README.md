# cliz

[English](README.md) | [简体中文](README_CN.md)

A modern, lightweight, and intuitive CLI framework for Go. cliz provides a simple API for creating command-line interfaces with subcommands, flags, validation, and advanced features like inheritance, positional arguments, and custom error handling.

## Feature Highlights

### Core Features
- **Subcommand Support**: Easily create hierarchical subcommand systems
- **Flexible Flag System**: Support for various flag types and validators
- **Struct Tag Integration**: Quick flag definition via struct tags
- **Positional Arguments**: Support for positional parameter passing
- **Flag Validation**: Built-in rich validators with custom validation rules support
- **Inheritance Mechanism**: Subcommands can inherit parent command flags

### Type Support
- Basic types: `string`, `int`, `uint`, `bool`, `float32`, `float64`
- Slice types: `[]string`, `[]int`, `[]uint`, `[]bool`, `[]float32`, `[]float64`
- Integer types: `int8`, `int16`, `int32`, `int64`, `uint8`, `uint16`, `uint32`, `uint64` and their slice forms

### Validators
- `Required`: Required flags
- `Range`: Numeric range validation
- `Length`: String length validation
- `Pattern`: Regular expression validation
- `In`: Enum value validation
- `Contains`: Substring validation
- `Email`: Email format validation
- `URL`: URL format validation
- `Alpha`: Contains only letters
- `AlphaNum`: Contains only letters and numbers
- Custom validators

## Installation

```bash
go get github.com/zkep/cliz
```

## Quick Start

### Basic Usage

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zkep/cliz"
)

func main() {
	app := cliz.NewCli("myapp", "A sample CLI application", "1.0.0")

	var name string
	var age int
	var verbose bool

	app.String("name", "Your name", &name, cliz.Required())
	app.Int("age", "Your age", &age, cliz.Range(18, 99))
	app.Bool("verbose", "Enable verbose output", &verbose)

	app.Action(func() error {
		fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
		if verbose {
			fmt.Println("Verbose mode enabled")
		}
		return nil
	})

	if err := app.Run(os.Args...); err != nil {
		log.Fatal(err)
	}
}
```

### Subcommands

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zkep/cliz"
)

func main() {
	app := cliz.NewCli("myapp", "CLI application with subcommands", "1.0.0")

	var name string
	app.String("name", "Your name", &name)

	greet := app.NewSubCommand("greet", "Greet someone")
	greet.Action(func() error {
		fmt.Printf("Hello, %s!\n", name)
		return nil
	})

	farewell := app.NewSubCommand("farewell", "Say goodbye")
	farewell.Action(func() error {
		fmt.Printf("Goodbye, %s!\n", name)
		return nil
	})

	app.AddCommand(greet)
	app.AddCommand(farewell)

	if err := app.Run(os.Args...); err != nil {
		log.Fatal(err)
	}
}
```

### Struct Tags

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zkep/cliz"
)

func main() {
	app := cliz.NewCli("struct_example", "CLI using struct flags", "1.0.0")

	type Config struct {
		Name     string `name:"name" description:"Your name" validate:"required"`
		Age      int    `name:"age" description:"Your age" validate:"range=18-99"`
		Email    string `name:"email" description:"Your email" validate:"email"`
		Verbose  bool   `name:"verbose" description:"Enable verbose output"`
	}

	var config Config
	app.AddFlags(&config)

	app.Action(func() error {
		fmt.Printf("Name: %s\n", config.Name)
		fmt.Printf("Age: %d\n", config.Age)
		fmt.Printf("Email: %s\n", config.Email)
		fmt.Printf("Verbose: %t\n", config.Verbose)
		return nil
	})

	if err := app.Run(os.Args...); err != nil {
		log.Fatal(err)
	}
}
```

### Positional Arguments

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zkep/cliz"
)

func main() {
	app := cliz.NewCli("positional_example", "CLI with positional args", "1.0.0")

	type Args struct {
		Source string `position:"0"`
		Dest   string `position:"1"`
	}

	var args Args
	app.AddPositionalArgs(&args)

	app.Action(func() error {
		fmt.Printf("Copy from %s to %s\n", args.Source, args.Dest)
		return nil
	})

	if err := app.Run(os.Args...); err != nil {
		log.Fatal(err)
	}
}
```

## Advanced Features

### Flag Inheritance

Subcommands can inherit flags from their parent command:

```go
parent := cliz.NewCommand("parent", "Parent command")
parent.String("name", "Your name", &name)

child := parent.NewSubCommand("child", "Child command")
child.InheritFlags(parent)
```

### PreRun Callback

Callback function executed before command execution:

```go
app.PreRun(func(c *cliz.Cli) error {
	fmt.Println("Before running command")
	return nil
})
```

### Custom Banner

```go
app.Banner(func(c *cliz.Cli) string {
	return fmt.Sprintf("Custom Banner - %s v%s", c.Name(), c.Version())
})
```

### Default Command

Default command executed when no command is specified:

```go
defaultCmd := cliz.NewCommand("default", "Default command")
defaultCmd.Action(func() error {
	fmt.Println("This is the default command")
	return nil
})

app.DefaultCommand(defaultCmd)
```

## API Documentation

### Main Types

#### `Cli`
- `NewCli(name, description, version string) *Cli`: Create a new CLI application
- `Run(args ...string) error`: Execute the CLI application
- `PrintHelp()`: Print help information
- `NewSubCommand(name, description string) *Command`: Create a subcommand
- `AddFlags(flags any) *Cli`: Add flags from struct
- `Action(callback Action) *Cli`: Set command execution callback
- `PreRun(callback func(*Cli) error)`: Set pre-run callback
- `DefaultCommand(defaultCommand *Command) *Cli`: Set default command

#### `Command`
- `NewCommand(name, description string) *Command`: Create a new command
- `AddSubCommand(cmd *Command) *Command`: Add a subcommand
- `String(name, description string, variable *string, validators ...Validator) *Command`: Add string flag
- `Int(name, description string, variable *int, validators ...Validator) *Command`: Add integer flag
- `Bool(name, description string, variable *bool, validators ...Validator) *Command`: Add boolean flag
- `Float64(name, description string, variable *float64, validators ...Validator) *Command`: Add float flag
- `StringSlice(name, description string, variable *[]string, validators ...Validator) *Command`: Add string slice flag
- `IntSlice(name, description string, variable *[]int, validators ...Validator) *Command`: Add integer slice flag
- `BoolSlice(name, description string, variable *[]bool, validators ...Validator) *Command`: Add boolean slice flag
- `AddPositionalArgs(args any) *Command`: Add positional arguments
- `InheritFlags(parent *Command) *Command`: Inherit flags from parent command

### Validators

- `Required() Validator`: Required validation
- `Range(min, max any) Validator`: Range validation
- `Length(min, max int) Validator`: Length validation
- `Pattern(regex string) Validator`: Regular expression validation
- `In(values ...string) Validator`: Enum validation
- `Contains(substr string) Validator`: Substring validation
- `Email() Validator`: Email validation
- `URL() Validator`: URL validation
- `Alpha() Validator`: Letter-only validation
- `AlphaNum() Validator`: Alphanumeric validation

Each validator supports `WithMessage(msg string)` method for custom error messages.

## Examples

For more examples, check the [_examples](_examples) directory:

- [basic](_examples/basic/main.go): Basic usage
- [subcommands](_examples/subcommands/main.go): Subcommands example
- [struct](_examples/struct/main.go): Struct flags example
- [positional](_examples/positional/main.go): Positional arguments example
- [advanced](_examples/advanced/main.go): Advanced validators example
- [inherit](_examples/inherit/main.go): Inheritance example
- [banner](_examples/banner/main.go): Banner example
- [prerun](_examples/prerun/main.go): PreRun callback example

## License

MIT License

## Contributing

Welcome to submit Issues and Pull Requests!
