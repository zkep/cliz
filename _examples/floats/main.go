package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zkep/cliz"
)

func main() {
	app := cliz.NewCli("floating_point_flags", "CLI application demonstrating floating point flag types", "1.0.0")

	var float64Flag float64
	var float64SliceFlag []float64

	app.Float64("float64", "float64 flag value", &float64Flag, cliz.Range(-1000000.0, 1000000.0))
	app.Float64Slice("float64s", "slice of float64 values", &float64SliceFlag)

	app.Action(func() error {
		fmt.Println("=== Floating Point Flags Demonstration ===")
		fmt.Println()
		fmt.Println("Single floating point values:")
		fmt.Printf("float64: %-15f\t(%T)\n", float64Flag, float64Flag)
		fmt.Println()
		fmt.Println("Floating point slice values:")
		fmt.Printf("float64s: %-13v\t(%T)\n", float64SliceFlag, float64SliceFlag)
		fmt.Println()
		fmt.Println("Use cases:")
		fmt.Println("- float64: double precision floating point numbers (15 decimal digits)")
		fmt.Println("- float64s: multiple floating point values with repeated flags")
		fmt.Println()
		fmt.Println("Example usage:")
		fmt.Println("  ./floating_point_flags --float64 2.71828")
		fmt.Println("  ./floating_point_flags --float64s 10.1 --float64s 20.2 --float64s 30.3")
		return nil
	})

	if err := app.Run(os.Args...); err != nil {
		log.Fatal(err)
	}
}
