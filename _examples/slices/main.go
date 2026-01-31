package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zkep/cliz"
)

func main() {
	app := cliz.NewCli("slice_flags", "CLI application demonstrating slice flag types", "1.0.0")

	var boolSliceFlag []bool
	var stringSliceFlag []string
	var intSliceFlag []int
	var int8SliceFlag []int8
	var int16SliceFlag []int16
	var int32SliceFlag []int32
	var int64SliceFlag []int64
	var uintSliceFlag []uint
	var uint8SliceFlag []uint8
	var uint16SliceFlag []uint16
	var uint32SliceFlag []uint32
	var uint64SliceFlag []uint64

	app.BoolSlice("bools", "slice of bool values", &boolSliceFlag)
	app.StringSlice("strings", "slice of string values", &stringSliceFlag)
	app.IntSlice("ints", "slice of int values", &intSliceFlag)
	app.Int8Slice("int8s", "slice of int8 values", &int8SliceFlag, cliz.Range(-128, 127))
	app.Int16Slice("int16s", "slice of int16 values", &int16SliceFlag, cliz.Range(-32768, 32767))
	app.Int32Slice("int32s", "slice of int32 values", &int32SliceFlag, cliz.Range(-2147483648, 2147483647))
	app.Int64Slice("int64s", "slice of int64 values", &int64SliceFlag, cliz.Range(-9223372036854775808, 9223372036854775807))
	app.UintSlice("uints", "slice of uint values", &uintSliceFlag, cliz.Range(0, 4294967295))
	app.Uint8Slice("uint8s", "slice of uint8 values", &uint8SliceFlag, cliz.Range(0, 255))
	app.Uint16Slice("uint16s", "slice of uint16 values", &uint16SliceFlag, cliz.Range(0, 65535))
	app.Uint32Slice("uint32s", "slice of uint32 values", &uint32SliceFlag, cliz.Range(0, 4294967295))
	app.Uint64Slice("uint64s", "slice of uint64 values", &uint64SliceFlag, cliz.Range(0, 18446744073709551615))

	app.Action(func() error {
		fmt.Println("=== Slice Flags Demonstration ===")
		fmt.Println()
		fmt.Printf("Bool slice flag: %-20v (%d items)\n", boolSliceFlag, len(boolSliceFlag))
		fmt.Printf("String slice flag: %-17v (%d items)\n", stringSliceFlag, len(stringSliceFlag))
		fmt.Printf("Int slice flag: %-20v (%d items)\n", intSliceFlag, len(intSliceFlag))
		fmt.Printf("Int8 slice flag: %-19v (%d items)\n", int8SliceFlag, len(int8SliceFlag))
		fmt.Printf("Int16 slice flag: %-18v (%d items)\n", int16SliceFlag, len(int16SliceFlag))
		fmt.Printf("Int32 slice flag: %-18v (%d items)\n", int32SliceFlag, len(int32SliceFlag))
		fmt.Printf("Int64 slice flag: %-18v (%d items)\n", int64SliceFlag, len(int64SliceFlag))
		fmt.Printf("Uint slice flag: %-19v (%d items)\n", uintSliceFlag, len(uintSliceFlag))
		fmt.Printf("Uint8 slice flag: %-18v (%d items)\n", uint8SliceFlag, len(uint8SliceFlag))
		fmt.Printf("Uint16 slice flag: %-17v (%d items)\n", uint16SliceFlag, len(uint16SliceFlag))
		fmt.Printf("Uint32 slice flag: %-17v (%d items)\n", uint32SliceFlag, len(uint32SliceFlag))
		fmt.Printf("Uint64 slice flag: %-17v (%d items)\n", uint64SliceFlag, len(uint64SliceFlag))
		fmt.Println()
		fmt.Println("Use cases:")
		fmt.Println("- Multiple values for the same flag (e.g., --file file1 --file file2 --file file3)")
		fmt.Println("- Collection of related items")
		fmt.Println("- Configuration options that can be specified multiple times")
		fmt.Println()
		fmt.Println("Example usage:")
		fmt.Println("  ./slice_flags --bools true --bools false --bools true")
		fmt.Println("  ./slice_flags --strings apple --strings banana --strings cherry")
		fmt.Println("  ./slice_flags --ints 10 --ints 20 --ints 30 --ints 40")
		return nil
	})

	if err := app.Run(os.Args...); err != nil {
		log.Fatal(err)
	}
}
