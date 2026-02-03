package cliz

import (
	"fmt"
	"os"
)

// Cli is the main CLI application object.
// It manages the command hierarchy, flags, and execution flow.
type Cli struct {
	version        string                    // Application version string
	rootCommand    *Command                  // Root command of the CLI hierarchy
	defaultCommand *Command                  // Command to execute when no command is specified
	preRunCommand  func(*Cli) error          // Callback executed before running any command
	bannerFunction func(*Cli) string         // Callback to generate banner output
	errorHandler   func(string, error) error // Custom error handler
}

// defaultBannerFunction generates the default application banner.
// The banner shows the application name, version (if provided), and short description.
func defaultBannerFunction(c *Cli) string {
	version := ""
	if len(c.Version()) > 0 {
		version = " " + c.Version()
	}
	return fmt.Sprintf("%s%s - %s", c.Name(), version, c.ShortDescription())
}

// NewCli creates a new CLI application with the given name, description, and version.
// The name is used as the root command name.
// The description appears in help output.
// The version is optional and appears in the banner if provided.
func NewCli(name, description, version string) *Cli {
	cli := &Cli{
		version:        version,
		bannerFunction: defaultBannerFunction,
	}
	cli.rootCommand = NewCommand(name, description)
	cli.rootCommand.setApp(cli)
	cli.rootCommand.setParentCommandPath("")
	return cli
}

// Version returns the application version string.
func (c *Cli) Version() string {
	return c.version
}

// Name returns the application name (root command name).
func (c *Cli) Name() string {
	return c.rootCommand.name
}

// ShortDescription returns the application short description (root command description).
func (c *Cli) ShortDescription() string {
	return c.rootCommand.shortdescription
}

// SetBannerFunction sets the function that generates the banner string.
// This allows customization of the banner output.
func (c *Cli) SetBannerFunction(fn func(*Cli) string) {
	c.bannerFunction = fn
}

// SetErrorFunction sets a custom error handler for undefined flag errors.
// The first argument receives the command path that was being executed.
// The second argument receives the undefined flag error.
// Custom handlers can provide more user-friendly error messages.
func (c *Cli) SetErrorFunction(fn func(string, error) error) {
	c.errorHandler = fn
}

// AddCommand adds a command to the application's root command.
// Commands added through this method become subcommands of the root.
func (c *Cli) AddCommand(command *Command) {
	c.rootCommand.AddCommand(command)
}

// PrintBanner prints the application banner to standard output.
// The banner format is determined by the bannerFunction.
func (c *Cli) PrintBanner() {
	fmt.Println(c.bannerFunction(c))
	fmt.Println("")
}

// PrintHelp prints the application's help information.
// This includes the root command help and all available subcommands.
func (c *Cli) PrintHelp() {
	c.rootCommand.PrintHelp()
}

// Run executes the CLI application with the given arguments.
// If no arguments are provided, it reads from os.Args[1:].
// This method initiates the command parsing and execution pipeline.
func (c *Cli) Run(args ...string) error {
	if c.preRunCommand != nil {
		err := c.preRunCommand(c)
		if err != nil {
			return err
		}
	}
	if args == nil {
		args = os.Args[1:]
	}
	return c.rootCommand.run(args)
}

// DefaultCommand sets the command to execute when no command is specified.
// The default command runs when the user doesn't provide any command arguments.
func (c *Cli) DefaultCommand(defaultCommand *Command) *Cli {
	c.defaultCommand = defaultCommand
	return c
}

// NewSubCommand creates a new subcommand under the root command.
// This is a convenience method that delegates to rootCommand.NewSubCommand.
func (c *Cli) NewSubCommand(name, description string) *Command {
	return c.rootCommand.NewSubCommand(name, description)
}

// NewSubCommandInheritFlags creates a new subcommand under the root command that inherits flags.
// The new subcommand inherits all flags from the root command (except help flag).
// This is a convenience method that delegates to rootCommand.NewSubCommandInheritFlags.
func (c *Cli) NewSubCommandInheritFlags(name, description string) *Command {
	return c.rootCommand.NewSubCommandInheritFlags(name, description)
}

// PreRun sets a callback function to execute before running any command.
// This can be used for pre-command initialization or validation.
func (c *Cli) PreRun(callback func(*Cli) error) {
	c.preRunCommand = callback
}

// AddFlags adds multiple flags to the root command by reflecting on a struct.
// The struct should be passed as a pointer.
// This method uses struct tags to configure flags automatically.
// Supported tags: `name:`, `description:`, `default:`, `pos:`, `sep:`.
func (c *Cli) AddFlags(flags any) *Cli {
	c.rootCommand.AddFlags(flags)
	return c
}

// Action sets the action callback for the root command.
// The callback is executed when the root command is run without subcommands.
func (c *Cli) Action(callback Action) *Cli {
	c.rootCommand.Action(callback)
	return c
}

// LongDescription sets the long description for the root command.
// The long description appears in detailed help output.
func (c *Cli) LongDescription(longdescription string) *Cli {
	c.rootCommand.LongDescription(longdescription)
	return c
}

// OtherArgs returns the non-flag arguments passed to the CLI.
// This should only be called within the context of an action callback.
// The returned slice contains all arguments that were not parsed as flags.
func (c *Cli) OtherArgs() []string {
	return c.rootCommand.flags.Args()
}

// Bool adds a boolean flag to the root command.
// The flag is added with the given name, description, and variable pointer.
// This is a convenience method that delegates to rootCommand.Bool.
func (c *Cli) Bool(name, description string, variable *bool, validators ...Validator) *Cli {
	c.rootCommand.Bool(name, description, variable, validators...)
	return c
}

// BoolSlice adds a slice of bool flags to the root command.
// The flag is added with the given name, description, and variable pointer.
// This is a convenience method that delegates to rootCommand.BoolSlice.
func (c *Cli) BoolSlice(name, description string, variable *[]bool, validators ...Validator) *Cli {
	c.rootCommand.BoolSlice(name, description, variable, validators...)
	return c
}

// String adds a string flag to the root command.
// The flag is added with the given name, description, and variable pointer.
// This is a convenience method that delegates to rootCommand.String.
func (c *Cli) String(name, description string, variable *string, validators ...Validator) *Cli {
	c.rootCommand.String(name, description, variable, validators...)
	return c
}

func (c *Cli) StringSlice(name, description string, variable *[]string, validators ...Validator) *Cli {
	c.rootCommand.StringSlice(name, description, variable, validators...)
	return c
}

// Int adds an integer flag to the root command.
// The flag is added with the given name, description, and variable pointer.
// This is a convenience method that delegates to rootCommand.Int.
func (c *Cli) Int(name, description string, variable *int, validators ...Validator) *Cli {
	c.rootCommand.Int(name, description, variable, validators...)
	return c
}

// IntSlice adds a slice of integer flags to the root command.
// The flag is added with the given name, description, and variable pointer.
// This is a convenience method that delegates to rootCommand.IntSlice.
func (c *Cli) IntSlice(name, description string, variable *[]int, validators ...Validator) *Cli {
	c.rootCommand.IntSlice(name, description, variable, validators...)
	return c
}

// Int8 adds an int8 flag to the root command.
// The flag is added with the given name, description, and variable pointer.
// This is a convenience method that delegates to rootCommand.Int8.
func (c *Cli) Int8(name, description string, variable *int8, validators ...Validator) *Cli {
	c.rootCommand.Int8(name, description, variable, validators...)
	return c
}

// Int16 adds an int16 flag to the root command.
// The flag is added with the given name, description, and variable pointer.
// This is a convenience method that delegates to rootCommand.Int16.
func (c *Cli) Int16(name, description string, variable *int16, validators ...Validator) *Cli {
	c.rootCommand.Int16(name, description, variable, validators...)
	return c
}

// Int32 adds an int32 flag to the root command.
// The flag is added with the given name, description, and variable pointer.
// This is a convenience method that delegates to rootCommand.Int32.
func (c *Cli) Int32(name, description string, variable *int32, validators ...Validator) *Cli {
	c.rootCommand.Int32(name, description, variable, validators...)
	return c
}

// Int64 adds an int64 flag to the root command.
// The flag is added with the given name, description, and variable pointer.
func (c *Cli) Int64(name, description string, variable *int64, validators ...Validator) *Cli {
	c.rootCommand.Int64(name, description, variable, validators...)
	return c
}

// Int8Slice adds a slice of int8 flags to the root command.
// The flag is added with the given name, description, and variable pointer.
// This is a convenience method that delegates to rootCommand.Int8Slice.
func (c *Cli) Int8Slice(name, description string, variable *[]int8, validators ...Validator) *Cli {
	c.rootCommand.Int8Slice(name, description, variable, validators...)
	return c
}

// Int16Slice adds a slice of int16 flags to the root command.
// The flag is added with the given name, description, and variable pointer.
// This is a convenience method that delegates to rootCommand.Int16Slice.
func (c *Cli) Int16Slice(name, description string, variable *[]int16, validators ...Validator) *Cli {
	c.rootCommand.Int16Slice(name, description, variable, validators...)
	return c
}

// Int32Slice adds a slice of int32 flags to the root command.
// The flag is added with the given name, description, and variable pointer.
// This is a convenience method that delegates to rootCommand.Int32Slice.
func (c *Cli) Int32Slice(name, description string, variable *[]int32, validators ...Validator) *Cli {
	c.rootCommand.Int32Slice(name, description, variable, validators...)
	return c
}

// Int64Slice adds a slice of int64 flags to the root command.
// The flag is added with the given name, description, and variable pointer.
// This is a convenience method that delegates to rootCommand.Int64Slice.
func (c *Cli) Int64Slice(name, description string, variable *[]int64, validators ...Validator) *Cli {
	c.rootCommand.Int64Slice(name, description, variable, validators...)
	return c
}

// Uint adds a uint flag to the root command.
// The flag is added with the given name, description, and variable pointer.
// This is a convenience method that delegates to rootCommand.Uint.
func (c *Cli) Uint(name, description string, variable *uint, validators ...Validator) *Cli {
	c.rootCommand.Uint(name, description, variable, validators...)
	return c
}

// UintSlice adds a slice of uint flags to the root command.
// The flag is added with the given name, description, and variable pointer.
// This is a convenience method that delegates to rootCommand.UintSlice.
func (c *Cli) UintSlice(name, description string, variable *[]uint, validators ...Validator) *Cli {
	c.rootCommand.UintSlice(name, description, variable, validators...)
	return c
}

// Uint8 adds a uint8 flag to the root command.
// The flag is added with the given name, description, and variable pointer.
// This is a convenience method that delegates to rootCommand.Uint8.
func (c *Cli) Uint8(name, description string, variable *uint8, validators ...Validator) *Cli {
	c.rootCommand.Uint8(name, description, variable, validators...)
	return c
}

// Uint8Slice adds a slice of uint8 flags to the root command.
// The flag is added with the given name, description, and variable pointer.
// This is a convenience method that delegates to rootCommand.Uint8Slice.
func (c *Cli) Uint8Slice(name, description string, variable *[]uint8, validators ...Validator) *Cli {
	c.rootCommand.Uint8Slice(name, description, variable, validators...)
	return c
}

// Uint16 adds a uint16 flag to the root command.
// The flag is added with the given name, description, and variable pointer.
// This is a convenience method that delegates to rootCommand.Uint16.
func (c *Cli) Uint16(name, description string, variable *uint16, validators ...Validator) *Cli {
	c.rootCommand.Uint16(name, description, variable, validators...)
	return c
}

// Uint16Slice adds a slice of uint16 flags to the root command.
// The flag is added with the given name, description, and variable pointer.
// This is a convenience method that delegates to rootCommand.Uint16Slice.
func (c *Cli) Uint16Slice(name, description string, variable *[]uint16, validators ...Validator) *Cli {
	c.rootCommand.Uint16Slice(name, description, variable, validators...)
	return c
}

// Uint32 adds a uint32 flag to the root command.
// The flag is added with the given name, description, and variable pointer.
// This is a convenience method that delegates to rootCommand.Uint32.
func (c *Cli) Uint32(name, description string, variable *uint32, validators ...Validator) *Cli {
	c.rootCommand.Uint32(name, description, variable, validators...)
	return c
}

// Uint32Slice adds a slice of uint32 flags to the root command.
// The flag is added with the given name, description, and variable pointer.
// This is a convenience method that delegates to rootCommand.Uint32Slice.
func (c *Cli) Uint32Slice(name, description string, variable *[]uint32, validators ...Validator) *Cli {
	c.rootCommand.Uint32Slice(name, description, variable, validators...)
	return c
}

// UInt64 adds a uint64 flag to the root command.
// The flag is added with the given name, description, and variable pointer.
// This is a convenience method that delegates to rootCommand.UInt64.
func (c *Cli) Uint64(name, description string, variable *uint64, validators ...Validator) *Cli {
	c.rootCommand.Uint64(name, description, variable, validators...)
	return c
}

// UInt64Slice adds a slice of uint64 flags to the root command.
// The flag is added with the given name, description, and variable pointer.
// This is a convenience method that delegates to rootCommand.UInt64Slice.
func (c *Cli) Uint64Slice(name, description string, variable *[]uint64, validators ...Validator) *Cli {
	c.rootCommand.Uint64Slice(name, description, variable, validators...)
	return c
}

// Float32 adds a float32 flag to the root command.
// The flag is added with the given name, description, and variable pointer.
func (c *Cli) Float32(name, description string, variable *float32, validators ...Validator) *Cli {
	c.rootCommand.Float32(name, description, variable, validators...)
	return c
}

// Float32Slice adds a slice of float32 flags to the root command.
// The flag is added with the given name, description, and variable pointer.
func (c *Cli) Float32Slice(name, description string, variable *[]float32, validators ...Validator) *Cli {
	c.rootCommand.Float32Slice(name, description, variable, validators...)
	return c
}

// Float64 adds a float64 flag to the root command.
// The flag is added with the given name, description, and variable pointer.
// This is a convenience method that delegates to rootCommand.Float64.
func (c *Cli) Float64(name, description string, variable *float64, validators ...Validator) *Cli {
	c.rootCommand.Float64(name, description, variable, validators...)
	return c
}

// Float64Slice adds a slice of float64 flags to the root command.
// The flag is added with the given name, description, and variable pointer.
// This is a convenience method that delegates to rootCommand.Float64Slice.
func (c *Cli) Float64Slice(name, description string, variable *[]float64, validators ...Validator) *Cli {
	c.rootCommand.Float64Slice(name, description, variable, validators...)
	return c
}

// AddPositionalArgs adds positional arguments to the root command by reflecting on a struct.
// The struct should be passed as a pointer.
// This method uses struct tags to configure positional arguments automatically.
// Supported tags: `position:`.
func (c *Cli) AddPositionalArgs(argsStruct any) *Cli {
	c.rootCommand.AddPositionalArgs(argsStruct)
	return c
}
