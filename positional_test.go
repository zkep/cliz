package cliz

import (
	"testing"
)

func TestPositionalArgsInvalidBool(t *testing.T) {
	type Args struct {
		Arg1 bool `position:"0"`
	}

	args := Args{}
	cmd := NewCommand("test", "test command")
	cmd.AddPositionalArgs(&args)

	cmd.Action(func() error {
		return nil
	})

	err := cmd.run([]string{"invalid"})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if args.Arg1 != false {
		t.Fatalf("Expected arg1 false (default), got %v", args.Arg1)
	}
}

func TestPositionalArgsNoPositionalArgsMap(t *testing.T) {
	cmd := NewCommand("test", "test command")

	cmd.Action(func() error {
		return nil
	})

	err := cmd.run([]string{"arg1", "arg2"})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestPositionalArgsInvalidInt64(t *testing.T) {
	type Args struct {
		Arg1 int64 `position:"0"`
	}

	args := Args{}
	cmd := NewCommand("test", "test command")
	cmd.AddPositionalArgs(&args)

	cmd.Action(func() error {
		return nil
	})

	err := cmd.run([]string{"invalid"})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if args.Arg1 != 0 {
		t.Fatalf("Expected arg1 0 (default), got %d", args.Arg1)
	}
}

func TestPositionalArgsInvalidUint(t *testing.T) {
	type Args struct {
		Arg1 uint `position:"0"`
	}

	args := Args{}
	cmd := NewCommand("test", "test command")
	cmd.AddPositionalArgs(&args)

	cmd.Action(func() error {
		return nil
	})

	err := cmd.run([]string{"invalid"})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if args.Arg1 != 0 {
		t.Fatalf("Expected arg1 0 (default), got %d", args.Arg1)
	}
}

func TestPositionalArgsInvalidUint64(t *testing.T) {
	type Args struct {
		Arg1 uint64 `position:"0"`
	}

	args := Args{}
	cmd := NewCommand("test", "test command")
	cmd.AddPositionalArgs(&args)

	cmd.Action(func() error {
		return nil
	})

	err := cmd.run([]string{"invalid"})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if args.Arg1 != 0 {
		t.Fatalf("Expected arg1 0 (default), got %d", args.Arg1)
	}
}

func TestPositionalArgsInvalidFloat32(t *testing.T) {
	type Args struct {
		Arg1 float32 `position:"0"`
	}

	args := Args{}
	cmd := NewCommand("test", "test command")
	cmd.AddPositionalArgs(&args)

	cmd.Action(func() error {
		return nil
	})

	err := cmd.run([]string{"invalid"})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if args.Arg1 != 0 {
		t.Fatalf("Expected arg1 0 (default), got %.2f", args.Arg1)
	}
}

func TestPositionalArgsInvalidFloat64(t *testing.T) {
	type Args struct {
		Arg1 float64 `position:"0"`
	}

	args := Args{}
	cmd := NewCommand("test", "test command")
	cmd.AddPositionalArgs(&args)

	cmd.Action(func() error {
		return nil
	})

	err := cmd.run([]string{"invalid"})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if args.Arg1 != 0 {
		t.Fatalf("Expected arg1 0 (default), got %.2f", args.Arg1)
	}
}

func TestPositionalArgsNegativeToUint(t *testing.T) {
	type Args struct {
		Arg1 uint `position:"0"`
	}

	args := Args{}
	cmd := NewCommand("test", "test command")
	cmd.AddPositionalArgs(&args)

	cmd.Action(func() error {
		t.Logf("Action called: Arg1 = %d", args.Arg1)
		return nil
	})
	err := cmd.parsePositionalArgs([]string{"-42"})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if args.Arg1 != 0 {
		t.Fatalf("Expected arg1 0 (default for invalid uint conversion), got %d", args.Arg1)
	}
}

func TestPositionalArgsNoPositionTag(t *testing.T) {
	type Args struct {
		Arg1 string `:""`
		Arg2 string
	}

	args := Args{}
	cmd := NewCommand("test", "test command")
	cmd.AddPositionalArgs(&args)

	cmd.Action(func() error {
		return nil
	})

	err := cmd.run([]string{"hello", "world"})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if args.Arg1 != "" {
		t.Fatalf("Expected arg1 empty (no position tag), got '%s'", args.Arg1)
	}
	if args.Arg2 != "" {
		t.Fatalf("Expected arg2 empty (no position tag), got '%s'", args.Arg2)
	}
}

func TestPositionalArgsEmptyPositionTag(t *testing.T) {
	type Args struct {
		Arg1 string `position:""`
		Arg2 string `position:"1"`
	}

	args := Args{}
	cmd := NewCommand("test", "test command")
	cmd.AddPositionalArgs(&args)

	cmd.Action(func() error {
		return nil
	})

	err := cmd.run([]string{"arg1", "arg2"})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if args.Arg1 != "" {
		t.Fatalf("Expected arg1 empty (empty position tag), got '%s'", args.Arg1)
	}
	if args.Arg2 != "arg2" {
		t.Fatalf("Expected arg2 'arg2', got '%s'", args.Arg2)
	}
}
