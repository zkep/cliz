package cliz

import (
	"testing"
)

func TestValidatorError(t *testing.T) {
	err := &ValidatorError{
		Field:   "test-field",
		Message: "test error",
	}
	if err.Error() != "test-field: test error" {
		t.Fatalf("Expected error message 'test-field: test error', got '%s'", err.Error())
	}
}

func TestLongDescription(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	cli.LongDescription("This is a long description that provides more details about the command.")
	err := cli.Run("--help")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestDefaultCommand(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	defaultCmd := cli.NewSubCommand("default", "default command")
	var defaultExecuted bool
	defaultCmd.Action(func() error {
		defaultExecuted = true
		return nil
	})
	cli.DefaultCommand(defaultCmd)
	err := cli.Run([]string{}...)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if defaultExecuted != true {
		t.Fatalf("Expected default command to be executed, got false")
	}
}

func TestOtherArgs(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var otherArgs []string
	cli.Action(func() error {
		otherArgs = cli.OtherArgs()
		return nil
	})
	err := cli.Run("arg1", "arg2", "arg3")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(otherArgs) != 3 {
		t.Fatalf("Expected 3 other args, got %d", len(otherArgs))
	}
	if otherArgs[0] != "arg1" || otherArgs[1] != "arg2" || otherArgs[2] != "arg3" {
		t.Fatalf("Unexpected other args: %v", otherArgs)
	}
}

func TestAddPositionalArgsOnCli(t *testing.T) {
	type Args struct {
		Name string `position:"0"`
		Age  int    `position:"1"`
	}
	var args Args
	cli := NewCli("test-app", "test description", "1.0.0")
	cli.AddPositionalArgs(&args)
	cli.Action(func() error {
		return nil
	})
	err := cli.Run("bob", "30")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if args.Name != "bob" {
		t.Fatalf("Expected name 'bob', got '%s'", args.Name)
	}
	if args.Age != 30 {
		t.Fatalf("Expected age 30, got %d", args.Age)
	}
}

func TestNoVersion(t *testing.T) {
	cli := NewCli("test-app", "test description", "")
	if cli.Version() != "" {
		t.Fatalf("Expected empty version, got '%s'", cli.Version())
	}
	err := cli.Run("--help")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestAddCommand(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	subCmd := NewCommand("subcmd", "subcommand description")
	var subExecuted bool
	subCmd.Action(func() error {
		subExecuted = true
		return nil
	})
	cli.AddCommand(subCmd)
	err := cli.Run("subcmd")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if subExecuted != true {
		t.Fatalf("Expected subcommand to be executed, got false")
	}
}

func TestNewSubCommand(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	subCmd := cli.NewSubCommand("subcmd", "subcommand description")
	var subExecuted bool
	subCmd.Action(func() error {
		subExecuted = true
		return nil
	})
	err := cli.Run("subcmd")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if subExecuted != true {
		t.Fatalf("Expected subcommand to be executed, got false")
	}
}

func TestNewSubCommandInheritFlags(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var rootFlag bool
	cli.Bool("root-flag", "root flag", &rootFlag)
	subCmd := cli.NewSubCommandInheritFlags("subcmd", "subcommand description")
	var subFlag bool
	subCmd.Bool("sub-flag", "sub flag", &subFlag)
	var subExecuted bool
	subCmd.Action(func() error {
		subExecuted = true
		return nil
	})
	err := cli.Run("subcmd", "--root-flag", "--sub-flag")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if subExecuted != true {
		t.Fatalf("Expected subcommand to be executed, got false")
	}
	if rootFlag != true {
		t.Fatalf("Expected root flag to be set, got false")
	}
	if subFlag != true {
		t.Fatalf("Expected sub flag to be set, got false")
	}
}

func TestAllFlagTypes(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")

	var boolField bool
	cli.Bool("bool-field", "bool field", &boolField)

	var stringField string
	cli.String("string-field", "string field", &stringField)

	var intField int
	cli.Int("int-field", "int field", &intField)

	var int8Field int8
	cli.Int8("int8-field", "int8 field", &int8Field)

	var int16Field int16
	cli.Int16("int16-field", "int16 field", &int16Field)

	var int32Field int32
	cli.Int32("int32-field", "int32 field", &int32Field)

	var int64Field int64
	cli.Int64("int64-field", "int64 field", &int64Field)

	var uintField uint
	cli.Uint("uint-field", "uint field", &uintField)

	var uint8Field uint8
	cli.Uint8("uint8-field", "uint8 field", &uint8Field)

	var uint16Field uint16
	cli.Uint16("uint16-field", "uint16 field", &uint16Field)

	var uint32Field uint32
	cli.Uint32("uint32-field", "uint32 field", &uint32Field)

	var uint64Field uint64
	cli.Uint64("uint64-field", "uint64 field", &uint64Field)

	var float32Field float32
	cli.Float32("float32-field", "float32 field", &float32Field)

	var float64Field float64
	cli.Float64("float64-field", "float64 field", &float64Field)

	var executed bool
	cli.Action(func() error {
		executed = true
		return nil
	})

	err := cli.Run("--bool-field", "--string-field=test", "--int-field=10", "--int8-field=10", "--int16-field=10", "--int32-field=10", "--int64-field=10", "--uint-field=10", "--uint8-field=10", "--uint16-field=10", "--uint32-field=10", "--uint64-field=10", "--float32-field=10.5", "--float64-field=10.5")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if executed != true {
		t.Fatalf("Expected executed to be true, got false")
	}

	if boolField != true {
		t.Fatalf("Expected boolField true, got %v", boolField)
	}
	if stringField != "test" {
		t.Fatalf("Expected stringField 'test', got '%s'", stringField)
	}
	if intField != 10 {
		t.Fatalf("Expected intField 10, got %d", intField)
	}
}

func TestAllSliceFlagTypes(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")

	var boolSlice []bool
	cli.BoolSlice("bool-slice", "bool slice", &boolSlice)

	var stringSlice []string
	cli.StringSlice("string-slice", "string slice", &stringSlice)

	var intSlice []int
	cli.IntSlice("int-slice", "int slice", &intSlice)

	var int8Slice []int8
	cli.Int8Slice("int8-slice", "int8 slice", &int8Slice)

	var int16Slice []int16
	cli.Int16Slice("int16-slice", "int16 slice", &int16Slice)

	var int32Slice []int32
	cli.Int32Slice("int32-slice", "int32 slice", &int32Slice)

	var int64Slice []int64
	cli.Int64Slice("int64-slice", "int64 slice", &int64Slice)

	var uintSlice []uint
	cli.UintSlice("uint-slice", "uint slice", &uintSlice)

	var uint8Slice []uint8
	cli.Uint8Slice("uint8-slice", "uint8 slice", &uint8Slice)

	var uint16Slice []uint16
	cli.Uint16Slice("uint16-slice", "uint16 slice", &uint16Slice)

	var uint32Slice []uint32
	cli.Uint32Slice("uint32-slice", "uint32 slice", &uint32Slice)

	var uint64Slice []uint64
	cli.Uint64Slice("uint64-slice", "uint64 slice", &uint64Slice)

	var float32Slice []float32
	cli.Float32Slice("float32-slice", "float32 slice", &float32Slice)

	var float64Slice []float64
	cli.Float64Slice("float64-slice", "float64 slice", &float64Slice)

	var executed bool
	cli.Action(func() error {
		executed = true
		return nil
	})

	err := cli.Run("--bool-slice=true", "--bool-slice=false", "--string-slice=a", "--string-slice=b", "--int-slice=1", "--int-slice=2")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if executed != true {
		t.Fatalf("Expected executed to be true, got false")
	}

	if len(boolSlice) != 2 || boolSlice[0] != true || boolSlice[1] != false {
		t.Fatalf("Expected boolSlice [true, false], got %v", boolSlice)
	}
	if len(stringSlice) != 2 || stringSlice[0] != "a" || stringSlice[1] != "b" {
		t.Fatalf("Expected stringSlice [a, b], got %v", stringSlice)
	}
	if len(intSlice) != 2 || intSlice[0] != 1 || intSlice[1] != 2 {
		t.Fatalf("Expected intSlice [1, 2], got %v", intSlice)
	}
}
