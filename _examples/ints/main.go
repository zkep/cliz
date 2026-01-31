package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zkep/cliz"
)

func main() {
	app := cliz.NewCli("integer_flags", "CLI application demonstrating integer flag types", "1.0.0")

	var int8Flag int8
	var int16Flag int16
	var int32Flag int32
	var uintFlag uint
	var uint8Flag uint8
	var uint16Flag uint16
	var uint32Flag uint32

	app.Int8("int8", "int8 flag value", &int8Flag, cliz.Range(-128, 127))
	app.Int16("int16", "int16 flag value", &int16Flag, cliz.Range(-32768, 32767))
	app.Int32("int32", "int32 flag value", &int32Flag, cliz.Range(-2147483648, 2147483647))
	app.Uint("uint", "uint flag value", &uintFlag, cliz.Range(0, 4294967295))
	app.Uint8("uint8", "uint8 flag value", &uint8Flag, cliz.Range(0, 255))
	app.Uint16("uint16", "uint16 flag value", &uint16Flag, cliz.Range(0, 65535))
	app.Uint32("uint32", "uint32 flag value", &uint32Flag, cliz.Range(0, 4294967295))

	app.Action(func() error {
		fmt.Println("=== Integer Flags Demonstration ===")
		fmt.Printf("int8: %-15d\t(%T)\n", int8Flag, int8Flag)
		fmt.Printf("int16: %-14d\t(%T)\n", int16Flag, int16Flag)
		fmt.Printf("int32: %-14c\t(%T)\n", int32Flag, int32Flag)
		fmt.Printf("uint: %-15d\t(%T)\n", uintFlag, uintFlag)
		fmt.Printf("uint8: %-14d\t(%T)\n", uint8Flag, uint8Flag)
		fmt.Printf("uint16: %-13d\t(%T)\n", uint16Flag, uint16Flag)
		fmt.Printf("uint32: %-13c\t(%T)\n", uint32Flag, uint32Flag)
		fmt.Println()
		fmt.Println("Use cases:")
		fmt.Println("- int8: small signed integers (-128 to 127)")
		fmt.Println("- uint8: small unsigned integers (0 to 255)")
		fmt.Println("- int16/uint16: medium integers for limited ranges")
		fmt.Println("- int32/uint32: standard 32-bit integer values")
		return nil
	})

	if err := app.Run(os.Args...); err != nil {
		log.Fatal(err)
	}
}
