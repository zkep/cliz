package cliz

import (
	"testing"
)

func TestCreate(t *testing.T) {
	cmd := NewCommand("test", "test command")
	if cmd == nil {
		t.Fatalf("NewCommand returned nil")
	}
	if cmd.name != "test" {
		t.Fatalf("Expected name 'test', got '%s'", cmd.name)
	}
	if cmd.shortdescription != "test command" {
		t.Fatalf("Expected description 'test command', got '%s'", cmd.shortdescription)
	}
}

func TestAction(t *testing.T) {
	cmd := NewCommand("test", "test command")
	cmd.Action(func() error {
		return nil
	})
	if cmd.actionCallback == nil {
		t.Fatalf("Expected actionCallback to be set, but it was nil")
	}
}

func TestHidden(t *testing.T) {
	cmd := NewCommand("test", "test command")
	if cmd.hidden != false {
		t.Fatalf("Expected hidden to be false, got true")
	}
	cmd.Hidden(true)
	if cmd.hidden != true {
		t.Fatalf("Expected hidden to be true, got false")
	}
}

func TestSub(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	subCmd := cli.NewSubCommand("sub", "subcommand")
	var executed bool
	subCmd.Action(func() error {
		executed = true
		return nil
	})
	err := cli.Run("sub")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if executed != true {
		t.Fatalf("Expected executed to be true, got false")
	}
}

func TestSubWithFlags(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	subCmd := cli.NewSubCommand("sub", "subcommand")
	var flagValue int
	subCmd.Int("count", "set count", &flagValue)
	var executed bool
	subCmd.Action(func() error {
		executed = true
		return nil
	})
	err := cli.Run("sub", "--count=10")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if executed != true {
		t.Fatalf("Expected executed to be true, got false")
	}
	if flagValue != 10 {
		t.Fatalf("Expected flagValue 10, got %d", flagValue)
	}
}

func TestNested(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	subCmd := cli.NewSubCommand("sub", "subcommand")
	subSubCmd := subCmd.NewSubCommand("nested", "nested subcommand")
	var executed bool
	subSubCmd.Action(func() error {
		executed = true
		return nil
	})
	err := cli.Run("sub", "nested")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if executed != true {
		t.Fatalf("Expected executed to be true, got false")
	}
}

func TestInherit(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	var parentFlag int
	cli.Int("parent-flag", "parent flag", &parentFlag)
	subCmd := cli.NewSubCommandInheritFlags("sub", "subcommand with inherited flags")
	var executed bool
	subCmd.Action(func() error {
		executed = true
		return nil
	})
	err := cli.Run("sub", "--parent-flag=10")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if executed != true {
		t.Fatalf("Expected executed to be true, got false")
	}
	if parentFlag != 10 {
		t.Fatalf("Expected parentFlag 10, got %d", parentFlag)
	}
}

func TestAction2(t *testing.T) {
	cmd := NewCommand("test", "test command")
	cmd.Action(func() error {
		return nil
	})
	if cmd.actionCallback == nil {
		t.Fatalf("Action callback was not set")
	}
}

func TestDefault(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	defaultCmd := NewCommand("default", "default command")
	var executed bool
	defaultCmd.Action(func() error {
		executed = true
		return nil
	})
	cli.DefaultCommand(defaultCmd)
	err := cli.Run()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if executed != true {
		t.Fatalf("Expected executed to be true, got false")
	}
}

func TestDefaultArgs(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	defaultCmd := NewCommand("default", "default command")
	var executed bool
	defaultCmd.Action(func() error {
		executed = true
		return nil
	})
	cli.DefaultCommand(defaultCmd)
	err := cli.Run("--help")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if executed != false {
		t.Fatalf("Expected executed to be false, got true")
	}
}

func TestMultiple(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	cmd1 := cli.NewSubCommand("cmd1", "command 1")
	cmd2 := cli.NewSubCommand("cmd2", "command 2")
	var cmd1Executed bool
	var cmd2Executed bool
	cmd1.Action(func() error {
		cmd1Executed = true
		return nil
	})
	cmd2.Action(func() error {
		cmd2Executed = true
		return nil
	})
	err := cli.Run("cmd1")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if cmd1Executed != true {
		t.Fatalf("Expected cmd1Executed to be true, got false")
	}
	if cmd2Executed != false {
		t.Fatalf("Expected cmd2Executed to be false, got true")
	}
	cmd1Executed = false
	err = cli.Run("cmd2")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if cmd1Executed != false {
		t.Fatalf("Expected cmd1Executed to be false, got true")
	}
	if cmd2Executed != true {
		t.Fatalf("Expected cmd2Executed to be true, got false")
	}
}

func TestLongDesc(t *testing.T) {
	cli := NewCli("test-app", "test description", "1.0.0")
	cmd := cli.NewSubCommand("cmd", "command")
	cmd.LongDescription("This is a long description\nwith multiple lines\nand detailed information")
	err := cli.Run("cmd", "--help")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestEmptyDesc(t *testing.T) {
	cli := NewCli("test-app", "", "1.0.0")
	err := cli.Run("--help")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}
