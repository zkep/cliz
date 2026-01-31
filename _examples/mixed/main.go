package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zkep/cliz"
)

func main() {
	app := cliz.NewCli("mixed_flags", "CLI application demonstrating mixed flag types", "1.0.0")

	var boolFlag bool
	var stringFlag string
	var intFlag int
	var floatFlag float64
	var stringSliceFlag []string
	var intSliceFlag []int

	app.Bool("boolean", "Boolean flag", &boolFlag)
	app.String("string", "String flag", &stringFlag, cliz.Required(), cliz.Len(5))
	app.Int("int", "Integer flag", &intFlag, cliz.Range(-1000, 1000))
	app.Float64("float", "Floating point flag", &floatFlag, cliz.Range(-10000.0, 10000.0))
	app.StringSlice("strings", "Multiple string values", &stringSliceFlag)
	app.IntSlice("ints", "Multiple integer values", &intSliceFlag)

	app.Action(func() error {
		fmt.Println("=== Mixed Flag Types Demonstration ===")
		fmt.Println()
		fmt.Println("Basic flags:")
		fmt.Printf("Boolean: %t\n", boolFlag)
		fmt.Printf("String: %s\n", stringFlag)
		fmt.Printf("Integer: %d\n", intFlag)
		fmt.Printf("Float: %f\n", floatFlag)
		fmt.Println()
		fmt.Println("Slice flags:")
		fmt.Printf("Strings slice: %v\n", stringSliceFlag)
		fmt.Printf("Ints slice: %v\n", intSliceFlag)
		fmt.Println()
		fmt.Println("This demonstrates how to mix different flag types in a single CLI application.")
		fmt.Println()
		fmt.Println("Example usage:")
		fmt.Println("  ./mixed_flags --boolean --string \"Hello World\" --int 42 --float 3.14")
		fmt.Println("  ./mixed_flags --string \"test\" --strings apple --strings banana --strings cherry")
		fmt.Println("  ./mixed_flags --int 100 --ints 1 --ints 2 --ints 3 --ints 4 --ints 5")
		return nil
	})

	if err := app.Run(os.Args...); err != nil {
		log.Fatal(err)
	}
}
