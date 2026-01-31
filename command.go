package cliz

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"
)

// Command represents a command that may be run by the user.
// It contains all the information needed to define and execute a command,
// including flags, subcommands, and action callbacks.
type Command struct {
	name              string                   // Name of the command
	commandPath       string                   // Full path to the command (including parent commands)
	shortdescription  string                   // Short description shown in help output
	longdescription   string                   // Long description shown in detailed help
	subCommands       []*Command               // List of subcommands
	subCommandsMap    map[string]*Command      // Map for fast subcommand lookup
	longestSubcommand int                      // Length of the longest subcommand name for formatting
	actionCallback    Action                   // Action to execute when the command runs
	app               *Cli                     // Reference to the parent Cli application
	flags             *flag.FlagSet            // Flag set for command-specific flags
	flagCount         int                      // Number of flags defined
	helpFlag          bool                     // Whether the help flag was requested
	hidden            bool                     // Whether the command is hidden from help
	positionalArgsMap map[string]reflect.Value // Map for positional arguments by index
	flagValidations   map[string][]Validator   // Map of flag names to validators
	flagVariables     map[string]reflect.Value // Map of flag names to their variable addresses for validation
}

// Action defines the callback function that executes when the command runs.
// Commands with an action callback will execute this function when invoked.
// The function should return an error if the command execution fails.
type Action func() error

// NewCommand creates a new Command with the given name and description.
// The command name should be unique within its parent command.
// The description should be a concise summary of what the command does.
func NewCommand(name string, description string) *Command {
	result := &Command{
		name:              name,
		shortdescription:  description,
		subCommandsMap:    make(map[string]*Command),
		flags:             flag.NewFlagSet(name, flag.ExitOnError),
		flagValidations:   make(map[string][]Validator),
		flagVariables:     make(map[string]reflect.Value),
		hidden:            false,
		positionalArgsMap: make(map[string]reflect.Value),
	}

	return result
}

func (c *Command) setParentCommandPath(parentCommandPath string) {
	// Set up command path
	if parentCommandPath != "" {
		c.commandPath += parentCommandPath + " "
	}
	c.commandPath += c.name

	// Set up flag set
	c.flags = flag.NewFlagSet(c.commandPath, flag.ContinueOnError)
	c.Bool("help", "Get help on the '"+strings.ToLower(c.commandPath)+"' command.", &c.helpFlag)
}

func (c *Command) inheritFlags(inheritFlags *flag.FlagSet) {
	// inherit flags
	inheritFlags.VisitAll(func(f *flag.Flag) {
		if f.Name != "help" {
			c.flags.Var(f.Value, f.Name, f.Usage)
		}
	})
}

func (c *Command) setApp(app *Cli) {
	c.app = app
}

// Action sets the action callback for the command.
// The action will be executed when the command runs if no subcommand is specified.
// Only one action can be set per command.
func (c *Command) Action(action Action) *Command {
	c.actionCallback = action
	return c
}

// NewSubCommand creates a new subcommand under the current command.
// The subcommand name should be unique within the current command.
// The description should be a concise summary of what the subcommand does.
func (c *Command) NewSubCommand(name string, description string) *Command {
	result := NewCommand(name, description)
	result.setApp(c.app)
	result.setParentCommandPath(c.commandPath)
	c.addSubCommand(result)
	return result
}

// addSubCommand adds a subcommand to the current command.
// This method is used internally to manage subcommands.
func (c *Command) addSubCommand(cmd *Command) {
	c.subCommands = append(c.subCommands, cmd)
	c.subCommandsMap[cmd.name] = cmd
	if len(cmd.name) > c.longestSubcommand {
		c.longestSubcommand = len(cmd.name)
	}
}

// Hidden marks the command as hidden.
// Hidden commands will not appear in the help output unless specifically requested.
func (c *Command) Hidden(hidden bool) *Command {
	c.hidden = hidden
	return c
}

// PrintHelp displays the help text for the command.
// The help text includes the command description, flags, and subcommands.
// Subcommands are only shown if the command has any subcommands.
func (c *Command) PrintHelp() {
	fmt.Printf("%s\n\n", c.commandPath)
	fmt.Printf("%s\n\n", c.shortdescription)

	if len(c.subCommands) > 0 {
		fmt.Printf("Commands:\n\n")
		for _, cmd := range c.subCommands {
			if cmd.hidden {
				continue
			}
			if cmd.shortdescription != "" {
				fmt.Printf("  %-*s %s\n", c.longestSubcommand, cmd.name, cmd.shortdescription)
			} else {
				fmt.Printf("  %-*s\n", c.longestSubcommand, cmd.name)
			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("Flags:\n\n")
	c.flags.VisitAll(func(f *flag.Flag) {
		fmt.Printf("  -%s", f.Name)
		usage := f.Usage
		if usage != "" {
			fmt.Printf(" %s", usage)
		}
		fmt.Printf("\n")
	})
	fmt.Printf("\n")
}

// SetName sets the name of the command.
// This method should be used carefully, as changing the command name
// after it has been added as a subcommand may break the subcommand map.
func (c *Command) SetName(name string) *Command {
	c.name = name
	return c
}

// SetShortDescription sets the short description of the command.
// The short description is shown in the help output alongside the command name.
// It should be a concise summary of what the command does.
func (c *Command) SetShortDescription(description string) *Command {
	c.shortdescription = description
	return c
}

// SetLongDescription sets the long description of the command.
// The long description is shown in the detailed help output.
// It can contain more comprehensive information about the command.
func (c *Command) SetLongDescription(description string) *Command {
	c.longdescription = description
	return c
}

// Name returns the name of the command.
// The name is the identifier used to invoke the command.
func (c *Command) Name() string {
	return c.name
}

// CommandPath returns the full path to the command, including parent commands.
// The command path is used in help output and error messages.
// It is constructed from the names of the parent commands and the current command.
func (c *Command) CommandPath() string {
	return c.commandPath
}

// ShortDescription returns the short description of the command.
// The short description is shown in the help output alongside the command name.
func (c *Command) ShortDescription() string {
	return c.shortdescription
}

// LongDescription sets the long description of the command.
// The long description is shown in the detailed help output.
// It can contain more comprehensive information about the command.
func (c *Command) LongDescription(description string) {
	c.longdescription = description
}

// AddCommand adds a command to the list of subcommands.
// This method is used internally to manage subcommands.
func (c *Command) AddCommand(cmd *Command) {
	c.addSubCommand(cmd)
}

// NewSubCommandInheritFlags creates a new subcommand that inherits all flags from the current command.
// The new subcommand will have all the same flags as the parent command.
// The subcommand name should be unique within the current command.
// The description should be a concise summary of what the subcommand does.
func (c *Command) NewSubCommandInheritFlags(name string, description string) *Command {
	result := NewCommand(name, description)
	result.setApp(c.app)
	result.setParentCommandPath(c.commandPath)
	result.inheritFlags(c.flags)
	c.addSubCommand(result)
	return result
}

// SubCommands returns the list of subcommands.
// The list includes all subcommands, including hidden ones.
func (c *Command) SubCommands() []*Command {
	return c.subCommands
}

// IsHidden returns whether the command is hidden from help output.
// Hidden commands are not shown in the help output unless specifically requested.
func (c *Command) IsHidden() bool {
	return c.hidden
}

// IsHelpRequested returns whether the help flag was requested.
// This method can be used to determine if the command should display help.
func (c *Command) IsHelpRequested() bool {
	return c.helpFlag
}

// App returns the parent Cli application.
// This method can be used to access application-wide configuration.
func (c *Command) App() *Cli {
	return c.app
}

// ExitWithError prints the error message and exits with a non-zero status code.
// This method should be used to display error messages to the user.
// The error message is printed to stderr.
func (c *Command) ExitWithError(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
}

// Exit prints the message and exits with a zero status code.
// This method should be used to display messages to the user without an error.
// The message is printed to stdout.
func (c *Command) Exit(message string) {
	fmt.Println(message)
	os.Exit(0)
}

// run executes the Command with the given arguments.
// This method handles flag parsing, subcommand execution, and action invocation.
// If a subcommand is specified, it delegates execution to that subcommand.
// If the help flag is requested, it displays the help text and exits.
// If an action callback is defined, it executes that callback.
// Returns an error if any step of the execution fails.
func (c *Command) run(args []string) error {
	// Check for help flag before parsing flags
	command := c
	command_args := args
	for i, arg := range args {
		subcommand := command.subCommandsMap[arg]
		if subcommand != nil {
			command = subcommand
			command_args = args[i+1:]
		}
	}
	if command.flagCount > 1 && (len(command_args) == 0 || (command_args[0] == "--help" || command_args[0] == "-h")) {
		command.helpFlag = true
	}
	// Help takes precedence
	if command.helpFlag {
		command.PrintHelp()
		return nil
	}
	// Parse flags
	err := command.parseFlags(command_args)
	if err != nil {
		if command.app != nil && command.app.errorHandler != nil {
			return command.app.errorHandler(command.commandPath, err)
		}
		return fmt.Errorf("%w\nSee '%s --help' for usage", err, command.commandPath)
	}

	// If we have a subcommand, run it
	if command.actionCallback != nil {
		return command.actionCallback()
	}

	// If we haven't specified a subcommand
	// check for an app level default command
	if command.app != nil && command.app.defaultCommand != nil {
		// Prevent recursion!
		if command.app.defaultCommand != command {
			// only run default command if no args passed
			if len(command_args) == 0 {
				return command.app.defaultCommand.run(command_args)
			}
		}
	}

	return nil
}

// Run executes the Command with the given arguments.
// This is a public wrapper around the private run method.
func (c *Command) Run(args ...string) error {
	return c.run(args)
}
