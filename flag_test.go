package cliz

import (
	"fmt"
	"testing"
)

func TestBool(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue bool
	cli.Bool("verbose", "enable verbose mode", &flagValue)
	err := cli.Run("--verbose")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if flagValue != true {
		t.Fatalf("Expected flagValue to be true, got false")
	}
}

func TestString(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue string
	cli.String("name", "set name", &flagValue)
	err := cli.Run("--name=test")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if flagValue != "test" {
		t.Fatalf("Expected flagValue 'test', got '%s'", flagValue)
	}
}

func TestInt(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue int
	cli.Int("count", "set count", &flagValue)
	err := cli.Run("--count=10")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if flagValue != 10 {
		t.Fatalf("Expected flagValue 10, got %d", flagValue)
	}
}

func TestMulti(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flag1 bool
	var flag2 string
	cli.Bool("verbose", "enable verbose mode", &flag1)
	cli.String("name", "set name", &flag2)
	err := cli.Run("--verbose", "--name=test")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if flag1 != true {
		t.Fatalf("Expected flag1 to be true, got false")
	}
	if flag2 != "test" {
		t.Fatalf("Expected flag2 'test', got '%s'", flag2)
	}
}

func TestInvalid(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	err := cli.Run("--invalid-flag")
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestHelp(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	err := cli.Run("--help")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestStringSlice(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue []string
	cli.StringSlice("names", "set names", &flagValue)
	err := cli.Run("--names=test1", "--names=test2")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(flagValue) != 2 {
		t.Fatalf("Expected flagValue length 2, got %d", len(flagValue))
	}
	if flagValue[0] != "test1" {
		t.Fatalf("Expected flagValue[0] 'test1', got '%s'", flagValue[0])
	}
	if flagValue[1] != "test2" {
		t.Fatalf("Expected flagValue[1] 'test2', got '%s'", flagValue[1])
	}
}

func TestIntSlice(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue []int
	cli.IntSlice("values", "set values", &flagValue)
	err := cli.Run("--values=1", "--values=2", "--values=3")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(flagValue) != 3 {
		t.Fatalf("Expected flagValue length 3, got %d", len(flagValue))
	}
	if flagValue[0] != 1 {
		t.Fatalf("Expected flagValue[0] 1, got %d", flagValue[0])
	}
	if flagValue[1] != 2 {
		t.Fatalf("Expected flagValue[1] 2, got %d", flagValue[1])
	}
	if flagValue[2] != 3 {
		t.Fatalf("Expected flagValue[2] 3, got %d", flagValue[2])
	}
}

func TestInt8(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue int8
	cli.Int8("verbose", "enable verbose mode", &flagValue)
	err := cli.Run("--verbose=10")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestUint64(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue uint64
	cli.Uint64("verbose", "enable verbose mode", &flagValue)
	err := cli.Run("--verbose=10")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestFloat64(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue float64
	cli.Float64("verbose", "enable verbose mode", &flagValue)
	err := cli.Run("--verbose=1.5")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestNoDesc(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue bool
	cli.Bool("verbose", "", &flagValue)
	err := cli.Run("--verbose")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if flagValue != true {
		t.Fatalf("Expected flagValue to be true, got false")
	}
}

func TestFlagDesc(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue bool
	cli.Bool("verbose", "This is a very long description for a flag that should wrap properly and provide all the necessary information about what the flag does and how it affects the command execution", &flagValue)
	err := cli.Run("--help")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestSpecialChars(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue string
	cli.String("name", "description with special characters: !@#$%^&*()_+", &flagValue)
	err := cli.Run("--name=test-value")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if flagValue != "test-value" {
		t.Fatalf("Expected flagValue 'test-value', got '%s'", flagValue)
	}
}

func TestInvalidValue(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue int
	cli.Int("count", "set count", &flagValue)
	err := cli.Run("--count=not-a-number")
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestEmptyValue(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue int
	cli.Int("count", "set count", &flagValue)
	err := cli.Run("--count=")
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestHyphen(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue bool
	cli.Bool("enable-feature", "enable feature", &flagValue)
	err := cli.Run("--enable-feature")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if flagValue != true {
		t.Fatalf("Expected flagValue to be true, got false")
	}
}

func TestEqualInValue(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue string
	cli.String("url", "set url", &flagValue)
	err := cli.Run("--url=http://example.com/path?param=value")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if flagValue != "http://example.com/path?param=value" {
		t.Fatalf("Expected flagValue 'http://example.com/path?param=value', got '%s'", flagValue)
	}
}

func TestSpaceInValue(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue string
	cli.String("message", "set message", &flagValue)
	err := cli.Run("--message=hello world")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if flagValue != "hello world" {
		t.Fatalf("Expected flagValue 'hello world', got '%s'", flagValue)
	}
}

func TestDashInValue(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue string
	cli.String("message", "set message", &flagValue)
	err := cli.Run("--message=hello-world")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if flagValue != "hello-world" {
		t.Fatalf("Expected flagValue 'hello-world', got '%s'", flagValue)
	}
}

func TestNumInValue(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue string
	cli.String("id", "set id", &flagValue)
	err := cli.Run("--id=12345")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if flagValue != "12345" {
		t.Fatalf("Expected flagValue '12345', got '%s'", flagValue)
	}
}

func TestMixedChars(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue string
	cli.String("token", "set token", &flagValue)
	err := cli.Run("--token=abc123-def456")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if flagValue != "abc123-def456" {
		t.Fatalf("Expected flagValue 'abc123-def456', got '%s'", flagValue)
	}
}

func TestSpecialValue(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue string
	cli.String("name", "set name", &flagValue)
	err := cli.Run("--name=test!@#$%^&*()_+")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if flagValue != "test!@#$%^&*()_+" {
		t.Fatalf(`Expected flagValue 'test!@#$%%^&*()_+', got '%s'`, flagValue)
	}
}

func TestMultiLine(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flag1 bool
	var flag2 string
	cli.Bool("verbose", "enable verbose mode", &flag1)
	cli.String("name", "set name", &flag2)
	err := cli.Run("--verbose", "--name=test")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if flag1 != true {
		t.Fatalf("Expected flag1 to be true, got false")
	}
	if flag2 != "test" {
		t.Fatalf("Expected flag2 'test', got '%s'", flag2)
	}
}

func TestShort(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue bool
	cli.Bool("v", "enable verbose mode (short)", &flagValue)
	cli.Bool("verbose", "enable verbose mode (long)", &flagValue)
	err := cli.Run("-v")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if flagValue != true {
		t.Fatalf("Expected flagValue to be true, got false")
	}
}

func TestShortLong(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flag1 bool
	var flag2 bool
	cli.Bool("v", "enable verbose mode (short)", &flag1)
	cli.Bool("verbose", "enable verbose mode (long)", &flag1)
	cli.Bool("d", "enable debug mode (short)", &flag2)
	cli.Bool("debug", "enable debug mode (long)", &flag2)
	err := cli.Run("-v", "-d")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if flag1 != true {
		t.Fatalf("Expected flag1 to be true, got false")
	}
	if flag2 != true {
		t.Fatalf("Expected flag2 to be true, got false")
	}
}

func TestRepeat(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue []string
	cli.StringSlice("name", "set name", &flagValue)
	err := cli.Run("--name=test1", "--name=test2")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(flagValue) != 2 {
		t.Fatalf("Expected flagValue length 2, got %d", len(flagValue))
	}
	if flagValue[0] != "test1" {
		t.Fatalf("Expected flagValue[0] 'test1', got '%s'", flagValue[0])
	}
	if flagValue[1] != "test2" {
		t.Fatalf("Expected flagValue[1] 'test2', got '%s'", flagValue[1])
	}
}

func TestRepeatOrder(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flag1 bool
	var flag2 []string
	cli.Bool("verbose", "enable verbose mode", &flag1)
	cli.StringSlice("name", "set name", &flag2)
	err := cli.Run("--name=test1", "--verbose", "--name=test2")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if flag1 != true {
		t.Fatalf("Expected flag1 to be true, got false")
	}
	if len(flag2) != 2 {
		t.Fatalf("Expected flag2 length 2, got %d", len(flag2))
	}
	if flag2[0] != "test1" {
		t.Fatalf("Expected flag2[0] 'test1', got '%s'", flag2[0])
	}
	if flag2[1] != "test2" {
		t.Fatalf("Expected flag2[1] 'test2', got '%s'", flag2[1])
	}
}

func TestRepeatPos(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue []int
	cli.IntSlice("count", "set count", &flagValue)
	err := cli.Run("--count=1", "--count=2")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(flagValue) != 2 {
		t.Fatalf("Expected flagValue length 2, got %d", len(flagValue))
	}
	if flagValue[0] != 1 {
		t.Fatalf("Expected flagValue[0] 1, got %d", flagValue[0])
	}
	if flagValue[1] != 2 {
		t.Fatalf("Expected flagValue[1] 2, got %d", flagValue[1])
	}
}

func TestRepeatPosFlags(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flag1 bool
	var flag2 []int
	cli.Bool("verbose", "enable verbose mode", &flag1)
	cli.IntSlice("count", "set count", &flag2)
	err := cli.Run("--count=1", "--verbose", "--count=2")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if flag1 != true {
		t.Fatalf("Expected flag1 to be true, got false")
	}
	if len(flag2) != 2 {
		t.Fatalf("Expected flag2 length 2, got %d", len(flag2))
	}
	if flag2[0] != 1 {
		t.Fatalf("Expected flag2[0] 1, got %d", flag2[0])
	}
	if flag2[1] != 2 {
		t.Fatalf("Expected flag2[1] 2, got %d", flag2[1])
	}
}

func TestRepeatPosArgs(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue []int
	cli.IntSlice("count", "set count", &flagValue)
	err := cli.Run("--count=1", "arg1", "--count=2", "arg2")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(flagValue) != 2 {
		t.Fatalf("Expected flagValue length 2, got %d", len(flagValue))
	}
	if flagValue[0] != 1 {
		t.Fatalf("Expected flagValue[0] 1, got %d", flagValue[0])
	}
	if flagValue[1] != 2 {
		t.Fatalf("Expected flagValue[1] 2, got %d", flagValue[1])
	}
}

func TestRepeatPosArgsFlags(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flag1 bool
	var flag2 []int
	cli.Bool("verbose", "enable verbose mode", &flag1)
	cli.IntSlice("count", "set count", &flag2)
	err := cli.Run("--count=1", "arg1", "--verbose", "--count=2", "arg2")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if flag1 != true {
		t.Fatalf("Expected flag1 to be true, got false")
	}
	if len(flag2) != 2 {
		t.Fatalf("Expected flag2 length 2, got %d", len(flag2))
	}
	if flag2[0] != 1 {
		t.Fatalf("Expected flag2[0] 1, got %d", flag2[0])
	}
	if flag2[1] != 2 {
		t.Fatalf("Expected flag2[1] 2, got %d", flag2[1])
	}
}

func TestRepeatPosArgsFlagsOrder(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flag1 bool
	var flag2 []int
	var flag3 string
	cli.Bool("verbose", "enable verbose mode", &flag1)
	cli.IntSlice("count", "set count", &flag2)
	cli.String("name", "set name", &flag3)
	err := cli.Run("--count=1", "arg1", "--verbose", "--name=test", "--count=2", "arg2")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if flag1 != true {
		t.Fatalf("Expected flag1 to be true, got false")
	}
	if len(flag2) != 2 {
		t.Fatalf("Expected flag2 length 2, got %d", len(flag2))
	}
	if flag2[0] != 1 {
		t.Fatalf("Expected flag2[0] 1, got %d", flag2[0])
	}
	if flag2[1] != 2 {
		t.Fatalf("Expected flag2[1] 2, got %d", flag2[1])
	}
	if flag3 != "test" {
		t.Fatalf("Expected flag3 'test', got '%s'", flag3)
	}
}

func TestRepeatPosArgsFlagsOrderCmd(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	subCmd := cli.NewSubCommand("cmd", "command")
	var flag1 bool
	var flag2 []int
	var flag3 string
	subCmd.Bool("verbose", "enable verbose mode", &flag1)
	subCmd.IntSlice("count", "set count", &flag2)
	subCmd.String("name", "set name", &flag3)
	var executed bool
	subCmd.Action(func() error {
		executed = true
		return nil
	})
	err := cli.Run("cmd", "--count=1", "arg1", "--verbose", "--name=test", "--count=2", "arg2")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if executed != true {
		t.Fatalf("Expected executed to be true, got false")
	}
	if flag1 != true {
		t.Fatalf("Expected flag1 to be true, got false")
	}
	if len(flag2) != 2 {
		t.Fatalf("Expected flag2 length 2, got %d", len(flag2))
	}
	if flag2[0] != 1 {
		t.Fatalf("Expected flag2[0] 1, got %d", flag2[0])
	}
	if flag2[1] != 2 {
		t.Fatalf("Expected flag2[1] 2, got %d", flag2[1])
	}
	if flag3 != "test" {
		t.Fatalf("Expected flag3 'test', got '%s'", flag3)
	}
}

func TestEmptyStr(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue string
	cli.String("name", "set name", &flagValue)
	err := cli.Run("--name=")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if flagValue != "" {
		t.Fatalf("Expected flagValue '', got '%s'", flagValue)
	}
}

func TestQuoted(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue string
	cli.String("name", "set name", &flagValue)
	err := cli.Run("--name='quoted value'")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if flagValue != "'quoted value'" {
		t.Fatalf("Expected flagValue '\\'quoted value\\'', got '%s'", flagValue)
	}
}

func TestSpaces(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue string
	cli.String("message", "set message", &flagValue)
	err := cli.Run("--message='hello world'")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if flagValue != "'hello world'" {
		t.Fatalf("Expected flagValue '\\'hello world\\'', got '%s'", flagValue)
	}
}

func TestInt16(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue int16
	cli.Int16("count", "set count", &flagValue)
	err := cli.Run("--count=10")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if flagValue != 10 {
		t.Fatalf("Expected flagValue 10, got %d", flagValue)
	}
}

func TestInt32(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue int32
	cli.Int32("count", "set count", &flagValue)
	err := cli.Run("--count=10")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if flagValue != 10 {
		t.Fatalf("Expected flagValue 10, got %d", flagValue)
	}
}

func TestInt64(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue int64
	cli.Int64("count", "set count", &flagValue)
	err := cli.Run("--count=10")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if flagValue != 10 {
		t.Fatalf("Expected flagValue 10, got %d", flagValue)
	}
}

func TestUint(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue uint
	cli.Uint("count", "set count", &flagValue)
	err := cli.Run("--count=10")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if flagValue != 10 {
		t.Fatalf("Expected flagValue 10, got %d", flagValue)
	}
}

func TestUint8(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue uint8
	cli.Uint8("count", "set count", &flagValue)
	err := cli.Run("--count=10")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if flagValue != 10 {
		t.Fatalf("Expected flagValue 10, got %d", flagValue)
	}
}

func TestUint16(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue uint16
	cli.Uint16("count", "set count", &flagValue)
	err := cli.Run("--count=10")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if flagValue != 10 {
		t.Fatalf("Expected flagValue 10, got %d", flagValue)
	}
}

func TestUint32(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue uint32
	cli.Uint32("count", "set count", &flagValue)
	err := cli.Run("--count=10")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if flagValue != 10 {
		t.Fatalf("Expected flagValue 10, got %d", flagValue)
	}
}

func TestFloat32(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue float32
	cli.Float32("count", "set count", &flagValue)
	err := cli.Run("--count=10.5")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if flagValue != 10.5 {
		t.Fatalf("Expected flagValue 10.5, got %f", flagValue)
	}
}

func TestBoolSlice(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue []bool
	cli.BoolSlice("flags", "set flags", &flagValue)
	err := cli.Run("--flags=true", "--flags=false", "--flags=true")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(flagValue) != 3 {
		t.Fatalf("Expected flagValue length 3, got %d", len(flagValue))
	}
	if flagValue[0] != true {
		t.Fatalf("Expected flagValue[0] true, got false")
	}
	if flagValue[1] != false {
		t.Fatalf("Expected flagValue[1] false, got true")
	}
	if flagValue[2] != true {
		t.Fatalf("Expected flagValue[2] true, got false")
	}
}

func TestInt8Slice(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue []int8
	cli.Int8Slice("counts", "set counts", &flagValue)
	err := cli.Run("--counts=1", "--counts=2", "--counts=3")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(flagValue) != 3 {
		t.Fatalf("Expected flagValue length 3, got %d", len(flagValue))
	}
	if flagValue[0] != 1 {
		t.Fatalf("Expected flagValue[0] 1, got %d", flagValue[0])
	}
	if flagValue[1] != 2 {
		t.Fatalf("Expected flagValue[1] 2, got %d", flagValue[1])
	}
	if flagValue[2] != 3 {
		t.Fatalf("Expected flagValue[2] 3, got %d", flagValue[2])
	}
}

func TestInt16Slice(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue []int16
	cli.Int16Slice("counts", "set counts", &flagValue)
	err := cli.Run("--counts=1", "--counts=2", "--counts=3")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(flagValue) != 3 {
		t.Fatalf("Expected flagValue length 3, got %d", len(flagValue))
	}
	if flagValue[0] != 1 {
		t.Fatalf("Expected flagValue[0] 1, got %d", flagValue[0])
	}
	if flagValue[1] != 2 {
		t.Fatalf("Expected flagValue[1] 2, got %d", flagValue[1])
	}
	if flagValue[2] != 3 {
		t.Fatalf("Expected flagValue[2] 3, got %d", flagValue[2])
	}
}

func TestInt32Slice(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue []int32
	cli.Int32Slice("counts", "set counts", &flagValue)
	err := cli.Run("--counts=1", "--counts=2", "--counts=3")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(flagValue) != 3 {
		t.Fatalf("Expected flagValue length 3, got %d", len(flagValue))
	}
	if flagValue[0] != 1 {
		t.Fatalf("Expected flagValue[0] 1, got %d", flagValue[0])
	}
	if flagValue[1] != 2 {
		t.Fatalf("Expected flagValue[1] 2, got %d", flagValue[1])
	}
	if flagValue[2] != 3 {
		t.Fatalf("Expected flagValue[2] 3, got %d", flagValue[2])
	}
}

func TestInt64Slice(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue []int64
	cli.Int64Slice("counts", "set counts", &flagValue)
	err := cli.Run("--counts=1", "--counts=2", "--counts=3")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(flagValue) != 3 {
		t.Fatalf("Expected flagValue length 3, got %d", len(flagValue))
	}
	if flagValue[0] != 1 {
		t.Fatalf("Expected flagValue[0] 1, got %d", flagValue[0])
	}
	if flagValue[1] != 2 {
		t.Fatalf("Expected flagValue[1] 2, got %d", flagValue[1])
	}
	if flagValue[2] != 3 {
		t.Fatalf("Expected flagValue[2] 3, got %d", flagValue[2])
	}
}

func TestUintSlice(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue []uint
	cli.UintSlice("counts", "set counts", &flagValue)
	err := cli.Run("--counts=1", "--counts=2", "--counts=3")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(flagValue) != 3 {
		t.Fatalf("Expected flagValue length 3, got %d", len(flagValue))
	}
	if flagValue[0] != 1 {
		t.Fatalf("Expected flagValue[0] 1, got %d", flagValue[0])
	}
	if flagValue[1] != 2 {
		t.Fatalf("Expected flagValue[1] 2, got %d", flagValue[1])
	}
	if flagValue[2] != 3 {
		t.Fatalf("Expected flagValue[2] 3, got %d", flagValue[2])
	}
}

func TestUint8Slice(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue []uint8
	cli.Uint8Slice("counts", "set counts", &flagValue)
	err := cli.Run("--counts=1", "--counts=2", "--counts=3")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(flagValue) != 3 {
		t.Fatalf("Expected flagValue length 3, got %d", len(flagValue))
	}
	if flagValue[0] != 1 {
		t.Fatalf("Expected flagValue[0] 1, got %d", flagValue[0])
	}
	if flagValue[1] != 2 {
		t.Fatalf("Expected flagValue[1] 2, got %d", flagValue[1])
	}
	if flagValue[2] != 3 {
		t.Fatalf("Expected flagValue[2] 3, got %d", flagValue[2])
	}
}

func TestUint16Slice(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue []uint16
	cli.Uint16Slice("counts", "set counts", &flagValue)
	err := cli.Run("--counts=1", "--counts=2", "--counts=3")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(flagValue) != 3 {
		t.Fatalf("Expected flagValue length 3, got %d", len(flagValue))
	}
	if flagValue[0] != 1 {
		t.Fatalf("Expected flagValue[0] 1, got %d", flagValue[0])
	}
	if flagValue[1] != 2 {
		t.Fatalf("Expected flagValue[1] 2, got %d", flagValue[1])
	}
	if flagValue[2] != 3 {
		t.Fatalf("Expected flagValue[2] 3, got %d", flagValue[2])
	}
}

func TestUint32Slice(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue []uint32
	cli.Uint32Slice("counts", "set counts", &flagValue)
	err := cli.Run("--counts=1", "--counts=2", "--counts=3")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(flagValue) != 3 {
		t.Fatalf("Expected flagValue length 3, got %d", len(flagValue))
	}
	if flagValue[0] != 1 {
		t.Fatalf("Expected flagValue[0] 1, got %d", flagValue[0])
	}
	if flagValue[1] != 2 {
		t.Fatalf("Expected flagValue[1] 2, got %d", flagValue[1])
	}
	if flagValue[2] != 3 {
		t.Fatalf("Expected flagValue[2] 3, got %d", flagValue[2])
	}
}

func TestUint64Slice(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue []uint64
	cli.Uint64Slice("counts", "set counts", &flagValue)
	err := cli.Run("--counts=1", "--counts=2", "--counts=3")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(flagValue) != 3 {
		t.Fatalf("Expected flagValue length 3, got %d", len(flagValue))
	}
	if flagValue[0] != 1 {
		t.Fatalf("Expected flagValue[0] 1, got %d", flagValue[0])
	}
	if flagValue[1] != 2 {
		t.Fatalf("Expected flagValue[1] 2, got %d", flagValue[1])
	}
	if flagValue[2] != 3 {
		t.Fatalf("Expected flagValue[2] 3, got %d", flagValue[2])
	}
}

func TestFloat32Slice(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue []float32
	cli.Float32Slice("counts", "set counts", &flagValue)
	err := cli.Run("--counts=1.1", "--counts=2.2", "--counts=3.3")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(flagValue) != 3 {
		t.Fatalf("Expected flagValue length 3, got %d", len(flagValue))
	}
	if flagValue[0] != 1.1 {
		t.Fatalf("Expected flagValue[0] 1.1, got %f", flagValue[0])
	}
	if flagValue[1] != 2.2 {
		t.Fatalf("Expected flagValue[1] 2.2, got %f", flagValue[1])
	}
	if flagValue[2] != 3.3 {
		t.Fatalf("Expected flagValue[2] 3.3, got %f", flagValue[2])
	}
}

func TestFloat64Slice(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue []float64
	cli.Float64Slice("counts", "set counts", &flagValue)
	err := cli.Run("--counts=1.1", "--counts=2.2", "--counts=3.3")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(flagValue) != 3 {
		t.Fatalf("Expected flagValue length 3, got %d", len(flagValue))
	}
	if flagValue[0] != 1.1 {
		t.Fatalf("Expected flagValue[0] 1.1, got %f", flagValue[0])
	}
	if flagValue[1] != 2.2 {
		t.Fatalf("Expected flagValue[1] 2.2, got %f", flagValue[1])
	}
	if flagValue[2] != 3.3 {
		t.Fatalf("Expected flagValue[2] 3.3, got %f", flagValue[2])
	}
}

func TestIntSliceInvalidValue(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue []int
	cli.IntSlice("counts", "set counts", &flagValue)
	err := cli.Run("--counts=not-a-number")
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestBoolSliceInvalidValue(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue []bool
	cli.BoolSlice("flags", "set flags", &flagValue)
	err := cli.Run("--flags=invalid-boolean")
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestFloat32SliceInvalidValue(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var flagValue []float32
	cli.Float32Slice("counts", "set counts", &flagValue)
	err := cli.Run("--counts=invalid-number")
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestAddFlagsSliceTypes(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	type config struct {
		BoolSlice    []bool    `name:"bool-slice" description:"bool slice"`
		StringSlice  []string  `name:"string-slice" description:"string slice"`
		IntSlice     []int     `name:"int-slice" description:"int slice"`
		Int8Slice    []int8    `name:"int8-slice" description:"int8 slice"`
		Int16Slice   []int16   `name:"int16-slice" description:"int16 slice"`
		Int32Slice   []int32   `name:"int32-slice" description:"int32 slice"`
		Int64Slice   []int64   `name:"int64-slice" description:"int64 slice"`
		UintSlice    []uint    `name:"uint-slice" description:"uint slice"`
		Uint8Slice   []uint8   `name:"uint8-slice" description:"uint8 slice"`
		Uint16Slice  []uint16  `name:"uint16-slice" description:"uint16 slice"`
		Uint32Slice  []uint32  `name:"uint32-slice" description:"uint32 slice"`
		Uint64Slice  []uint64  `name:"uint64-slice" description:"uint64 slice"`
		Float32Slice []float32 `name:"float32-slice" description:"float32 slice"`
		Float64Slice []float64 `name:"float64-slice" description:"float64 slice"`
	}
	var cfg config
	cli.AddFlags(&cfg)

	err := cli.Run("--bool-slice=true", "--bool-slice=false", "--string-slice=hello", "--string-slice=world", "--int-slice=10", "--int-slice=20", "--int8-slice=1", "--int8-slice=2", "--int16-slice=100", "--int16-slice=200", "--int32-slice=1000", "--int32-slice=2000", "--int64-slice=10000", "--int64-slice=20000", "--uint-slice=5", "--uint-slice=10", "--uint8-slice=3", "--uint8-slice=4", "--uint16-slice=50", "--uint16-slice=100", "--uint32-slice=10000", "--uint32-slice=20000", "--uint64-slice=100000", "--uint64-slice=200000", "--float32-slice=3.14", "--float32-slice=2.718", "--float64-slice=1.618", "--float64-slice=0.618")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(cfg.BoolSlice) != 2 || cfg.BoolSlice[0] != true || cfg.BoolSlice[1] != false {
		t.Fatalf("Expected BoolSlice [true, false], got %v", cfg.BoolSlice)
	}
	if len(cfg.StringSlice) != 2 || cfg.StringSlice[0] != "hello" || cfg.StringSlice[1] != "world" {
		t.Fatalf("Expected StringSlice [hello, world], got %v", cfg.StringSlice)
	}
	if len(cfg.IntSlice) != 2 || cfg.IntSlice[0] != 10 || cfg.IntSlice[1] != 20 {
		t.Fatalf("Expected IntSlice [10, 20], got %v", cfg.IntSlice)
	}
	if len(cfg.Int8Slice) != 2 || cfg.Int8Slice[0] != 1 || cfg.Int8Slice[1] != 2 {
		t.Fatalf("Expected Int8Slice [1, 2], got %v", cfg.Int8Slice)
	}
	if len(cfg.Int16Slice) != 2 || cfg.Int16Slice[0] != 100 || cfg.Int16Slice[1] != 200 {
		t.Fatalf("Expected Int16Slice [100, 200], got %v", cfg.Int16Slice)
	}
	if len(cfg.Int32Slice) != 2 || cfg.Int32Slice[0] != 1000 || cfg.Int32Slice[1] != 2000 {
		t.Fatalf("Expected Int32Slice [1000, 2000], got %v", cfg.Int32Slice)
	}
	if len(cfg.Int64Slice) != 2 || cfg.Int64Slice[0] != 10000 || cfg.Int64Slice[1] != 20000 {
		t.Fatalf("Expected Int64Slice [10000, 20000], got %v", cfg.Int64Slice)
	}
	if len(cfg.UintSlice) != 2 || cfg.UintSlice[0] != 5 || cfg.UintSlice[1] != 10 {
		t.Fatalf("Expected UintSlice [5, 10], got %v", cfg.UintSlice)
	}
	if len(cfg.Uint8Slice) != 2 || cfg.Uint8Slice[0] != 3 || cfg.Uint8Slice[1] != 4 {
		t.Fatalf("Expected Uint8Slice [3, 4], got %v", cfg.Uint8Slice)
	}
	if len(cfg.Uint16Slice) != 2 || cfg.Uint16Slice[0] != 50 || cfg.Uint16Slice[1] != 100 {
		t.Fatalf("Expected Uint16Slice [50, 100], got %v", cfg.Uint16Slice)
	}
	if len(cfg.Uint32Slice) != 2 || cfg.Uint32Slice[0] != 10000 || cfg.Uint32Slice[1] != 20000 {
		t.Fatalf("Expected Uint32Slice [10000, 20000], got %v", cfg.Uint32Slice)
	}
	if len(cfg.Uint64Slice) != 2 || cfg.Uint64Slice[0] != 100000 || cfg.Uint64Slice[1] != 200000 {
		t.Fatalf("Expected Uint64Slice [100000, 200000], got %v", cfg.Uint64Slice)
	}
	if len(cfg.Float32Slice) != 2 || cfg.Float32Slice[0] != 3.14 || cfg.Float32Slice[1] != 2.718 {
		t.Fatalf("Expected Float32Slice [3.14, 2.718], got %v", cfg.Float32Slice)
	}
	if len(cfg.Float64Slice) != 2 || cfg.Float64Slice[0] != 1.618 || cfg.Float64Slice[1] != 0.618 {
		t.Fatalf("Expected Float64Slice [1.618, 0.618], got %v", cfg.Float64Slice)
	}
}

func TestAddFlagsDefaultValues(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	type config struct {
		BoolFlag    bool    `cliz:"--bool-flag,description=bool flag"`
		IntFlag     int     `cliz:"--int-flag,description=int flag"`
		UintFlag    uint    `cliz:"--uint-flag,description=uint flag"`
		Float64Flag float64 `cliz:"--float64-flag,description=float64 flag"`
	}
	var cfg = config{
		BoolFlag:    true,
		IntFlag:     100,
		UintFlag:    200,
		Float64Flag: 3.14,
	}
	cli.AddFlags(&cfg)

	err := cli.Run([]string{}...)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !cfg.BoolFlag {
		t.Fatalf("Expected BoolFlag to be true")
	}
	if cfg.IntFlag != 100 {
		t.Fatalf("Expected IntFlag to be 100, got %d", cfg.IntFlag)
	}
	if cfg.UintFlag != 200 {
		t.Fatalf("Expected UintFlag to be 200, got %d", cfg.UintFlag)
	}
	if cfg.Float64Flag != 3.14 {
		t.Fatalf("Expected Float64Flag to be 3.14, got %f", cfg.Float64Flag)
	}
}

func TestSetFieldDefaultValueAllTypes(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	type config struct {
		BoolField    bool    `name:"bool-field" description:"bool field" default:"true"`
		StringField  string  `name:"string-field" description:"string field" default:"test"`
		IntField     int     `name:"int-field" description:"int field" default:"100"`
		Int8Field    int8    `name:"int8-field" description:"int8 field" default:"10"`
		Int16Field   int16   `name:"int16-field" description:"int16 field" default:"100"`
		Int32Field   int32   `name:"int32-field" description:"int32 field" default:"1000"`
		Int64Field   int64   `name:"int64-field" description:"int64 field" default:"10000"`
		UintField    uint    `name:"uint-field" description:"uint field" default:"50"`
		Uint8Field   uint8   `name:"uint8-field" description:"uint8 field" default:"5"`
		Uint16Field  uint16  `name:"uint16-field" description:"uint16 field" default:"500"`
		Uint32Field  uint32  `name:"uint32-field" description:"uint32 field" default:"5000"`
		Uint64Field  uint64  `name:"uint64-field" description:"uint64 field" default:"50000"`
		Float32Field float32 `name:"float32-field" description:"float32 field" default:"3.14"`
		Float64Field float64 `name:"float64-field" description:"float64 field" default:"2.718"`
	}
	var cfg config
	cli.AddFlags(&cfg)

	err := cli.Run([]string{}...)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestAddFlagsWithAllPrimitiveTypes(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	type config struct {
		BoolField    bool    `name:"bool-field" description:"bool field"`
		StringField  string  `name:"string-field" description:"string field"`
		IntField     int     `name:"int-field" description:"int field"`
		Int8Field    int8    `name:"int8-field" description:"int8 field"`
		Int16Field   int16   `name:"int16-field" description:"int16 field"`
		Int32Field   int32   `name:"int32-field" description:"int32 field"`
		Int64Field   int64   `name:"int64-field" description:"int64 field"`
		UintField    uint    `name:"uint-field" description:"uint field"`
		Uint8Field   uint8   `name:"uint8-field" description:"uint8 field"`
		Uint16Field  uint16  `name:"uint16-field" description:"uint16 field"`
		Uint32Field  uint32  `name:"uint32-field" description:"uint32 field"`
		Uint64Field  uint64  `name:"uint64-field" description:"uint64 field"`
		Float32Field float32 `name:"float32-field" description:"float32 field"`
		Float64Field float64 `name:"float64-field" description:"float64 field"`
	}
	var cfg config
	cli.AddFlags(&cfg)

	err := cli.Run("--bool-field", "--string-field=test", "--int-field=10", "--int8-field=5", "--int16-field=100", "--int32-field=1000", "--int64-field=10000", "--uint-field=10", "--uint8-field=5", "--uint16-field=100", "--uint32-field=1000", "--uint64-field=10000", "--float32-field=3.14", "--float64-field=2.718")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestValidateTagsWithRange(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	type config struct {
		Age int `name:"age" description:"age" validate:"range=18-100"`
	}
	var cfg config
	cli.AddFlags(&cfg)
}

func TestValidateTagsRequiredWithRange(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	type config struct {
		Age int `name:"age" description:"age" validate:"required,range=18-100"`
	}
	var cfg config
	cli.AddFlags(&cfg)
}

func TestValidateTagsRequiredOnly(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	type config struct {
		Name string `name:"name" description:"name" validate:"required"`
	}
	var cfg config
	cli.AddFlags(&cfg)
}

func TestValidateTagsWithLen(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	type config struct {
		Name string `name:"name" description:"name" validate:"len=5"`
	}
	var cfg config
	cli.AddFlags(&cfg)
}

func TestValidateTagsWithPattern(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	type config struct {
		Name string `name:"name" description:"name" validate:"pattern=^[a-z]+$"`
	}
	var cfg config
	cli.AddFlags(&cfg)
}

func TestValidateTagsWithEmail(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	type config struct {
		Email string `name:"email" description:"email" validate:"email"`
	}
	var cfg config
	cli.AddFlags(&cfg)
}

func TestValidateTagsWithURL(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	type config struct {
		URL string `name:"url" description:"url" validate:"url"`
	}
	var cfg config
	cli.AddFlags(&cfg)
}

func TestValidateTagsWithAlpha(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	type config struct {
		Name string `name:"name" description:"name" validate:"alpha"`
	}
	var cfg config
	cli.AddFlags(&cfg)
}

func TestValidateTagsWithAlphanum(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	type config struct {
		Name string `name:"name" description:"name" validate:"alphanum"`
	}
	var cfg config
	cli.AddFlags(&cfg)
}

func TestValidateTagsWithContains(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	type config struct {
		Name string `name:"name" description:"name" validate:"contains=test"`
	}
	var cfg config
	cli.AddFlags(&cfg)
}

func TestValidateTagsWithIn(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	type config struct {
		Option string `name:"option" description:"option" validate:"in=option1,option2"`
	}
	var cfg config
	cli.AddFlags(&cfg)
}

func TestValidateTagsWithEq(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	type config struct {
		Value int `name:"value" description:"value" validate:"eq=10"`
	}
	var cfg config
	cli.AddFlags(&cfg)
}

func TestValidateTagsWithGt(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	type config struct {
		Value int `name:"value" description:"value" validate:"gt=5"`
	}
	var cfg config
	cli.AddFlags(&cfg)
}

func TestValidateTagsWithLt(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	type config struct {
		Value int `name:"value" description:"value" validate:"lt=10"`
	}
	var cfg config
	cli.AddFlags(&cfg)
}

func TestValidatorErrorWithMessageTemplate(t *testing.T) {
	err := &ValidatorError{
		Field:           "testField",
		MessageTemplate: "value '%s' is invalid with min '%d'",
		Args:            []any{"invalid", 10},
	}

	expected := "testField: value 'invalid' is invalid with min '10'"
	if err.Error() != expected {
		t.Fatalf("Expected error '%s', got '%s'", expected, err.Error())
	}
}

func TestValidatorErrorWithEmptyTemplate(t *testing.T) {
	err := &ValidatorError{
		Field:           "testField",
		MessageTemplate: "",
		Message:         "fallback message",
	}

	expected := "testField: fallback message"
	if err.Error() != expected {
		t.Fatalf("Expected error '%s', got '%s'", expected, err.Error())
	}
}

func TestSetFieldDefaultValueWithErrors(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	type config struct {
		InvalidBool bool `name:"invalid-bool" description:"invalid bool" default:"not-a-bool"`
		InvalidInt  int  `name:"invalid-int" description:"invalid int" default:"not-a-number"`
	}
	var cfg config
	cli.AddFlags(&cfg)

	err := cli.Run([]string{}...)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestWithMessageForAllValidators(t *testing.T) {
	validators := []struct {
		name      string
		validator Validator
	}{
		{"Range", Range(0, 100)},
		{"Required", Required()},
		{"Len", Len(5)},
		{"Pattern", Pattern("^[a-z]+$")},
		{"In", In("option1", "option2")},
		{"Eq", Eq(10)},
		{"Gt", Gt(10)},
		{"Lt", Lt(100)},
		{"Contains", Contains("test")},
		{"Email", Email()},
		{"URL", URL()},
		{"Alpha", Alpha()},
		{"Alphanum", Alphanum()},
	}

	for _, v := range validators {
		t.Run(v.name, func(t *testing.T) {
			withMsg := v.validator.WithMessage("custom message")
			if withMsg == nil {
				t.Fatalf("WithMessage returned nil for %s", v.name)
			}
		})
	}
}

func TestAddFlagsWithValidationTags(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	type config struct {
		Email string `name:"email" description:"email address" validate:"email"`
		Age   int    `name:"age" description:"age" validate:"min=18,max=100"`
	}
	var cfg config
	cli.AddFlags(&cfg)

	err := cli.Run("--email=test@example.com", "--age=25")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestParseFlagsWithValidationErrors(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var age int
	cli.Int("age", "age", &age, Range(18, 100))

	err := cli.Run("--age=15")
	if err == nil {
		t.Fatalf("Expected validation error for age < 18")
	}
}

func TestParseFlagsWithMultipleValidationErrors(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var name string
	var age int
	cli.String("name", "name", &name, Required())
	cli.Int("age", "age", &age, Range(18, 100))

	err := cli.Run("--age=15")
	if err == nil {
		t.Fatalf("Expected validation error for missing name and invalid age")
	}
}

func TestParseFlagsWithPositionalArgs(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	type Args struct {
		File string `position:"0"`
	}
	var args Args
	cli.AddPositionalArgs(&args)

	err := cli.Run("file.txt")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if args.File != "file.txt" {
		t.Fatalf("Expected file.txt, got %s", args.File)
	}
}

func TestParseFlagsWithFlagParsingError(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var age int
	cli.Int("age", "age", &age)

	err := cli.Run("--age=not-a-number")
	if err == nil {
		t.Fatalf("Expected flag parsing error")
	}
}

func TestParseFlagsWithCustomValidator(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var name string
	cli.String("name", "name", &name, Custom(func(value any) error {
		s, _ := value.(string)
		if len(s) < 3 {
			return fmt.Errorf("name must be at least 3 characters")
		}
		return nil
	}))

	err := cli.Run("--name=ab")
	if err == nil {
		t.Fatalf("Expected validation error for short name")
	}
}

func TestSetFieldDefaultValueComplexScenario(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	type config struct {
		BoolWithDefault   bool   `name:"bool-default" description:"bool with default" default:"false"`
		StringWithDefault string `name:"string-default" description:"string with default" default:"default"`
		IntWithDefault    int    `name:"int-default" description:"int with default" default:"0"`
	}
	var cfg config
	cli.AddFlags(&cfg)

	err := cli.Run([]string{}...)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}
