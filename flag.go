package cliz

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/zkep/cliz/validator"
)

// Bool adds a boolean flag to the command.
// The flag will be parsed from the command line arguments and stored in the provided variable.
func (c *Command) Bool(name, description string, variable *bool, validators ...Validator) *Command {
	c.flags.BoolVar(variable, name, *variable, description)
	c.flagVariables[name] = reflect.ValueOf(variable).Elem()
	if len(validators) > 0 {
		c.flagValidations[name] = validators
	}
	c.flagCount++
	return c
}

// String adds a string flag to the command.
// The flag will be parsed from the command line arguments and stored in the provided variable.
func (c *Command) String(name, description string, variable *string, validators ...Validator) *Command {
	c.flags.StringVar(variable, name, *variable, description)
	c.flagVariables[name] = reflect.ValueOf(variable).Elem()
	if len(validators) > 0 {
		c.flagValidations[name] = validators
	}
	c.flagCount++
	return c
}

// Int adds an integer flag to the command.
// The flag will be parsed from the command line arguments and stored in the provided variable.
func (c *Command) Int(name, description string, variable *int, validators ...Validator) *Command {
	c.flags.IntVar(variable, name, *variable, description)
	c.flagVariables[name] = reflect.ValueOf(variable).Elem()
	if len(validators) > 0 {
		c.flagValidations[name] = validators
	}
	c.flagCount++
	return c
}

// Int8 adds an int8 flag to the command.
// The flag will be parsed from the command line arguments and stored in the provided variable.
func (c *Command) Int8(name, description string, variable *int8, validators ...Validator) *Command {
	c.flags.Var(newInt8Value(*variable, variable), name, description)
	c.flagVariables[name] = reflect.ValueOf(variable).Elem()
	if len(validators) > 0 {
		c.flagValidations[name] = validators
	}
	c.flagCount++
	return c
}

// Int8Slice adds a repeated int8 flag to the command.
// The flag can be specified multiple times, and all values will be collected into the slice.
func (c *Command) Int8Slice(name, description string, variable *[]int8, validators ...Validator) *Command {
	c.flags.Var(newInt8SliceValue(*variable, variable), name, description)
	c.flagVariables[name] = reflect.ValueOf(variable).Elem()
	if len(validators) > 0 {
		c.flagValidations[name] = validators
	}
	c.flagCount++
	return c
}

// Int16 adds an int16 flag to the command.
// The flag will be parsed from the command line arguments and stored in the provided variable.
func (c *Command) Int16(name, description string, variable *int16, validators ...Validator) *Command {
	c.flags.Var(newInt16Value(*variable, variable), name, description)
	c.flagVariables[name] = reflect.ValueOf(variable).Elem()
	if len(validators) > 0 {
		c.flagValidations[name] = validators
	}
	c.flagCount++
	return c
}

// Int16Slice adds a repeated int16 flag to the command.
// The flag can be specified multiple times, and all values will be collected into the slice.
func (c *Command) Int16Slice(name, description string, variable *[]int16, validators ...Validator) *Command {
	c.flags.Var(newInt16SliceValue(*variable, variable), name, description)
	c.flagVariables[name] = reflect.ValueOf(variable).Elem()
	if len(validators) > 0 {
		c.flagValidations[name] = validators
	}
	c.flagCount++
	return c
}

// Int32 adds an int32 flag to the command.
// The flag will be parsed from the command line arguments and stored in the provided variable.
func (c *Command) Int32(name, description string, variable *int32, validators ...Validator) *Command {
	c.flags.Var(newInt32Value(*variable, variable), name, description)
	c.flagVariables[name] = reflect.ValueOf(variable).Elem()
	if len(validators) > 0 {
		c.flagValidations[name] = validators
	}
	c.flagCount++
	return c
}

// Int32Slice adds a repeated int32 flag to the command.
// The flag can be specified multiple times, and all values will be collected into the slice.
func (c *Command) Int32Slice(name, description string, variable *[]int32, validators ...Validator) *Command {
	c.flags.Var(newInt32SliceValue(*variable, variable), name, description)
	c.flagVariables[name] = reflect.ValueOf(variable).Elem()
	if len(validators) > 0 {
		c.flagValidations[name] = validators
	}
	c.flagCount++
	return c
}

// Int64 adds an int64 flag to the command.
// The flag will be parsed from the command line arguments and stored in the provided variable.
func (c *Command) Int64(name, description string, variable *int64, validators ...Validator) *Command {
	c.flags.Int64Var(variable, name, *variable, description)
	c.flagVariables[name] = reflect.ValueOf(variable).Elem()
	if len(validators) > 0 {
		c.flagValidations[name] = validators
	}
	c.flagCount++
	return c
}

// Int64Slice adds a repeated int64 flag to the command.
// The flag can be specified multiple times, and all values will be collected into the slice.
func (c *Command) Int64Slice(name, description string, variable *[]int64, validators ...Validator) *Command {
	c.flags.Var(newInt64SliceValue(*variable, variable), name, description)
	c.flagVariables[name] = reflect.ValueOf(variable).Elem()
	if len(validators) > 0 {
		c.flagValidations[name] = validators
	}
	c.flagCount++
	return c
}

// Uint adds a uint flag to the command.
// The flag will be parsed from the command line arguments and stored in the provided variable.
func (c *Command) Uint(name, description string, variable *uint, validators ...Validator) *Command {
	c.flags.UintVar(variable, name, *variable, description)
	c.flagVariables[name] = reflect.ValueOf(variable).Elem()
	if len(validators) > 0 {
		c.flagValidations[name] = validators
	}
	c.flagCount++
	return c
}

// UintSlice adds a repeated uint flag to the command.
// The flag can be specified multiple times, and all values will be collected into the slice.
func (c *Command) UintSlice(name, description string, variable *[]uint, validators ...Validator) *Command {
	c.flags.Var(newUintSliceValue(*variable, variable), name, description)
	c.flagVariables[name] = reflect.ValueOf(variable).Elem()
	if len(validators) > 0 {
		c.flagValidations[name] = validators
	}
	c.flagCount++
	return c
}

// Uint8 adds a uint8 flag to the command.
// The flag will be parsed from the command line arguments and stored in the provided variable.
func (c *Command) Uint8(name, description string, variable *uint8, validators ...Validator) *Command {
	c.flags.Var(newUint8Value(*variable, variable), name, description)
	c.flagVariables[name] = reflect.ValueOf(variable).Elem()
	if len(validators) > 0 {
		c.flagValidations[name] = validators
	}
	c.flagCount++
	return c
}

// Uint8Slice adds a repeated uint8 flag to the command.
// The flag can be specified multiple times, and all values will be collected into the slice.
func (c *Command) Uint8Slice(name, description string, variable *[]uint8, validators ...Validator) *Command {
	c.flags.Var(newUint8SliceValue(*variable, variable), name, description)
	c.flagVariables[name] = reflect.ValueOf(variable).Elem()
	if len(validators) > 0 {
		c.flagValidations[name] = validators
	}
	c.flagCount++
	return c
}

// Uint16 adds a uint16 flag to the command.
// The flag will be parsed from the command line arguments and stored in the provided variable.
func (c *Command) Uint16(name, description string, variable *uint16, validators ...Validator) *Command {
	c.flags.Var(newUint16Value(*variable, variable), name, description)
	c.flagVariables[name] = reflect.ValueOf(variable).Elem()
	if len(validators) > 0 {
		c.flagValidations[name] = validators
	}
	c.flagCount++
	return c
}

// Uint16Slice adds a repeated uint16 flag to the command.
// The flag can be specified multiple times, and all values will be collected into the slice.
func (c *Command) Uint16Slice(name, description string, variable *[]uint16, validators ...Validator) *Command {
	c.flags.Var(newUint16SliceValue(*variable, variable), name, description)
	c.flagVariables[name] = reflect.ValueOf(variable).Elem()
	if len(validators) > 0 {
		c.flagValidations[name] = validators
	}
	c.flagCount++
	return c
}

// Uint32 adds a uint32 flag to the command.
// The flag will be parsed from the command line arguments and stored in the provided variable.
func (c *Command) Uint32(name, description string, variable *uint32, validators ...Validator) *Command {
	c.flags.Var(newUint32Value(*variable, variable), name, description)
	c.flagVariables[name] = reflect.ValueOf(variable).Elem()
	if len(validators) > 0 {
		c.flagValidations[name] = validators
	}
	c.flagCount++
	return c
}

// Uint32Slice adds a repeated uint32 flag to the command.
// The flag can be specified multiple times, and all values will be collected into the slice.
func (c *Command) Uint32Slice(name, description string, variable *[]uint32, validators ...Validator) *Command {
	c.flags.Var(newUint32SliceValue(*variable, variable), name, description)
	c.flagVariables[name] = reflect.ValueOf(variable).Elem()
	if len(validators) > 0 {
		c.flagValidations[name] = validators
	}
	c.flagCount++
	return c
}

// Uint64 adds a uint64 flag to the command.
// The flag will be parsed from the command line arguments and stored in the provided variable.
func (c *Command) Uint64(name, description string, variable *uint64, validators ...Validator) *Command {
	c.flags.Var(newUint64Value(*variable, variable), name, description)
	c.flagVariables[name] = reflect.ValueOf(variable).Elem()
	if len(validators) > 0 {
		c.flagValidations[name] = validators
	}
	c.flagCount++
	return c
}

// Uint64Slice adds a repeated uint64 flag to the command.
// The flag can be specified multiple times, and all values will be collected into the slice.
func (c *Command) Uint64Slice(name, description string, variable *[]uint64, validators ...Validator) *Command {
	c.flags.Var(newUint64SliceValue(*variable, variable), name, description)
	c.flagVariables[name] = reflect.ValueOf(variable).Elem()
	if len(validators) > 0 {
		c.flagValidations[name] = validators
	}
	c.flagCount++
	return c
}

// Float64Flag adds a float64 flag to the command.
// The flag will be parsed from the command line arguments and stored in the provided variable.
func (c *Command) Float64(name, description string, variable *float64, validators ...Validator) *Command {
	c.flags.Float64Var(variable, name, *variable, description)
	c.flagVariables[name] = reflect.ValueOf(variable).Elem()
	if len(validators) > 0 {
		c.flagValidations[name] = validators
	}
	c.flagCount++
	return c
}

// Float32Flag adds a float32 flag to the command.
// The flag will be parsed from the command line arguments and stored in the provided variable.
func (c *Command) Float32(name, description string, variable *float32, validators ...Validator) *Command {
	c.flags.Var(newFloat32Value(*variable, variable), name, description)
	c.flagVariables[name] = reflect.ValueOf(variable).Elem()
	if len(validators) > 0 {
		c.flagValidations[name] = validators
	}
	c.flagCount++
	return c
}

// Float32sFlag adds a repeated float32 flag to the command.
// The flag can be specified multiple times, and all values will be collected into the slice.
func (c *Command) Float32Slice(name, description string, variable *[]float32, validators ...Validator) *Command {
	c.flags.Var(newFloat32SliceValue(*variable, variable), name, description)
	c.flagVariables[name] = reflect.ValueOf(variable).Elem()
	if len(validators) > 0 {
		c.flagValidations[name] = validators
	}
	c.flagCount++
	return c
}

// Float64sFlag adds a repeated float64 flag to the command.
// The flag can be specified multiple times, and all values will be collected into the slice.
func (c *Command) Float64Slice(name, description string, variable *[]float64, validators ...Validator) *Command {
	c.flags.Var(newFloat64SliceValue(*variable, variable), name, description)
	c.flagVariables[name] = reflect.ValueOf(variable).Elem()
	if len(validators) > 0 {
		c.flagValidations[name] = validators
	}
	c.flagCount++
	return c
}

func (c *Command) StringSlice(name, description string, variable *[]string, validators ...Validator) *Command {
	c.flags.Var(newStringSliceValue(*variable, variable), name, description)
	c.flagVariables[name] = reflect.ValueOf(variable).Elem()
	if len(validators) > 0 {
		c.flagValidations[name] = validators
	}
	c.flagCount++
	return c
}

func (c *Command) IntSlice(name, description string, variable *[]int, validators ...Validator) *Command {
	c.flags.Var(newIntSliceValue(*variable, variable), name, description)
	c.flagVariables[name] = reflect.ValueOf(variable).Elem()
	if len(validators) > 0 {
		c.flagValidations[name] = validators
	}
	c.flagCount++
	return c
}

// BoolSliceFlag adds a boolean slice flag to the command.
// The flag will be parsed from the command line arguments and stored in the provided variable.
// The slice will contain all values provided for the flag.
func (c *Command) BoolSlice(name, description string, variable *[]bool, validators ...Validator) *Command {
	c.flags.Var(newBoolSliceValue(*variable, variable), name, description)
	c.flagVariables[name] = reflect.ValueOf(variable).Elem()
	if len(validators) > 0 {
		c.flagValidations[name] = validators
	}
	c.flagCount++
	return c
}

// stringSliceValue is a wrapper around a slice of strings that implements the flag.Value interface.
func newStringSliceValue(val []string, p *[]string) flag.Value {
	return &stringSliceValue{val: val, p: p}
}

// stringSliceValue is a wrapper around a slice of strings that implements the flag.Value interface.
type stringSliceValue struct {
	val []string
	p   *[]string
}

func (s *stringSliceValue) String() string {
	return fmt.Sprintf("%v", s.val)
}

func (s *stringSliceValue) Set(value string) error {
	*s.p = append(*s.p, value)
	return nil
}

// int8Value is a wrapper around an int8 that implements the flag.Value interface.
func newInt8Value(val int8, p *int8) flag.Value {
	return &int8Value{val: val, p: p}
}

// int8Value is a wrapper around an int8 that implements the flag.Value interface.
type int8Value struct {
	val int8
	p   *int8
}

func (i *int8Value) String() string {
	return fmt.Sprintf("%d", i.val)
}

func (i *int8Value) Set(value string) error {
	intValue, err := strconv.Atoi(value)
	if err == nil {
		*i.p = int8(intValue)
	}
	return err
}

func newInt16Value(val int16, p *int16) flag.Value {
	return &int16Value{val: val, p: p}
}

// int16Value is a wrapper around an int16 that implements the flag.Value interface.
type int16Value struct {
	val int16
	p   *int16
}

func (i *int16Value) String() string {
	return fmt.Sprintf("%d", i.val)
}

func (i *int16Value) Set(value string) error {
	intValue, err := strconv.Atoi(value)
	if err == nil {
		*i.p = int16(intValue)
	}
	return err
}

func newInt32Value(val int32, p *int32) flag.Value {
	return &int32Value{val: val, p: p}
}

// int32Value is a wrapper around an int32 that implements the flag.Value interface.
type int32Value struct {
	val int32
	p   *int32
}

func (i *int32Value) String() string {
	return fmt.Sprintf("%d", i.val)
}

func (i *int32Value) Set(value string) error {
	intValue, err := strconv.Atoi(value)
	if err == nil {
		*i.p = int32(intValue)
	}
	return err
}

func newUint8Value(val uint8, p *uint8) flag.Value {
	return &uint8Value{val: val, p: p}
}

// uint8Value is a wrapper around an uint8 that implements the flag.Value interface.
type uint8Value struct {
	val uint8
	p   *uint8
}

func (i *uint8Value) String() string {
	return fmt.Sprintf("%d", i.val)
}

func (i *uint8Value) Set(value string) error {
	intValue, err := strconv.Atoi(value)
	if err == nil {
		*i.p = uint8(intValue)
	}
	return err
}

func newUint16Value(val uint16, p *uint16) flag.Value {
	return &uint16Value{val: val, p: p}
}

// uint16Value is a wrapper around an uint16 that implements the flag.Value interface.
type uint16Value struct {
	val uint16
	p   *uint16
}

func (i *uint16Value) String() string {
	return fmt.Sprintf("%d", i.val)
}

func (i *uint16Value) Set(value string) error {
	intValue, err := strconv.Atoi(value)
	if err == nil {
		*i.p = uint16(intValue)
	}
	return err
}

func newUint32Value(val uint32, p *uint32) flag.Value {
	return &uint32Value{val: val, p: p}
}

// uint32Value is a wrapper around an uint32 that implements the flag.Value interface.
type uint32Value struct {
	val uint32
	p   *uint32
}

func (i *uint32Value) String() string {
	return fmt.Sprintf("%d", i.val)
}

func (i *uint32Value) Set(value string) error {
	intValue, err := strconv.Atoi(value)
	if err == nil {
		*i.p = uint32(intValue)
	}
	return err
}

func newUint64Value(val uint64, p *uint64) flag.Value {
	return &uint64Value{val: val, p: p}
}

// uint64Value is a wrapper around an uint64 that implements the flag.Value interface.
type uint64Value struct {
	val uint64
	p   *uint64
}

func (i *uint64Value) String() string {
	return fmt.Sprintf("%d", i.val)
}

func (i *uint64Value) Set(value string) error {
	intValue, err := strconv.ParseUint(value, 10, 64)
	if err == nil {
		*i.p = uint64(intValue)
	}
	return err
}

func newFloat32Value(val float32, p *float32) flag.Value {
	return &float32Value{val: val, p: p}
}

// float32Value is a wrapper around a float32 that implements the flag.Value interface.
type float32Value struct {
	val float32
	p   *float32
}

func (i *float32Value) String() string {
	return fmt.Sprintf("%f", i.val)
}

func (i *float32Value) Set(value string) error {
	floatValue, err := strconv.ParseFloat(value, 32)
	if err == nil {
		*i.p = float32(floatValue)
	}
	return err
}

func newInt8SliceValue(val []int8, p *[]int8) flag.Value {
	return &int8SliceValue{val: val, p: p}
}

// int8sValue is a wrapper around a slice of int8s that implements the flag.Value interface.
type int8SliceValue struct {
	val []int8
	p   *[]int8
}

func (i *int8SliceValue) String() string {
	return fmt.Sprintf("%v", i.val)
}

func (i *int8SliceValue) Set(value string) error {
	intValue, err := strconv.Atoi(value)
	if err == nil {
		*i.p = append(*i.p, int8(intValue))
	}
	return err
}

func newInt16SliceValue(val []int16, p *[]int16) flag.Value {
	return &int16SliceValue{val: val, p: p}
}

// int16SliceValue is a wrapper around a slice of int16s that implements the flag.Value interface.
type int16SliceValue struct {
	val []int16
	p   *[]int16
}

func (i *int16SliceValue) String() string {
	return fmt.Sprintf("%v", i.val)
}

func (i *int16SliceValue) Set(value string) error {
	intValue, err := strconv.Atoi(value)
	if err == nil {
		*i.p = append(*i.p, int16(intValue))
	}
	return err
}

func newInt32SliceValue(val []int32, p *[]int32) flag.Value {
	return &int32SliceValue{val: val, p: p}
}

// int32SliceValue is a wrapper around a slice of int32s that implements the flag.Value interface.
type int32SliceValue struct {
	val []int32
	p   *[]int32
}

func (i *int32SliceValue) String() string {
	return fmt.Sprintf("%v", i.val)
}

func (i *int32SliceValue) Set(value string) error {
	intValue, err := strconv.Atoi(value)
	if err == nil {
		*i.p = append(*i.p, int32(intValue))
	}
	return err
}

func newInt64SliceValue(val []int64, p *[]int64) flag.Value {
	return &int64SliceValue{val: val, p: p}
}

// int64SliceValue is a wrapper around a slice of int64s that implements the flag.Value interface.
type int64SliceValue struct {
	val []int64
	p   *[]int64
}

func (i *int64SliceValue) String() string {
	return fmt.Sprintf("%v", i.val)
}

func (i *int64SliceValue) Set(value string) error {
	intValue, err := strconv.ParseInt(value, 10, 64)
	if err == nil {
		*i.p = append(*i.p, int64(intValue))
	}
	return err
}

func newUintSliceValue(val []uint, p *[]uint) flag.Value {
	return &uintSliceValue{val: val, p: p}
}

// uintSliceValue is a wrapper around a slice of uints that implements the flag.Value interface.
type uintSliceValue struct {
	val []uint
	p   *[]uint
}

func (i *uintSliceValue) String() string {
	return fmt.Sprintf("%v", i.val)
}

func (i *uintSliceValue) Set(value string) error {
	intValue, err := strconv.ParseUint(value, 10, 64)
	if err == nil {
		*i.p = append(*i.p, uint(intValue))
	}
	return err
}

func newUint8SliceValue(val []uint8, p *[]uint8) flag.Value {
	return &uint8SliceValue{val: val, p: p}
}

// uint8SliceValue is a wrapper around a slice of uint8s that implements the flag.Value interface.
type uint8SliceValue struct {
	val []uint8
	p   *[]uint8
}

func (i *uint8SliceValue) String() string {
	return fmt.Sprintf("%v", i.val)
}

func (i *uint8SliceValue) Set(value string) error {
	intValue, err := strconv.Atoi(value)
	if err == nil {
		*i.p = append(*i.p, uint8(intValue))
	}
	return err
}

func newUint16SliceValue(val []uint16, p *[]uint16) flag.Value {
	return &uint16SliceValue{val: val, p: p}
}

// uint16SliceValue is a wrapper around a slice of uint16s that implements the flag.Value interface.
type uint16SliceValue struct {
	val []uint16
	p   *[]uint16
}

func (i *uint16SliceValue) String() string {
	return fmt.Sprintf("%v", i.val)
}

func (i *uint16SliceValue) Set(value string) error {
	intValue, err := strconv.Atoi(value)
	if err == nil {
		*i.p = append(*i.p, uint16(intValue))
	}
	return err
}

func newUint32SliceValue(val []uint32, p *[]uint32) flag.Value {
	return &uint32SliceValue{val: val, p: p}
}

// uint32SliceValue is a wrapper around a slice of uint32s that implements the flag.Value interface.
type uint32SliceValue struct {
	val []uint32
	p   *[]uint32
}

func (i *uint32SliceValue) String() string {
	return fmt.Sprintf("%v", i.val)
}

func (i *uint32SliceValue) Set(value string) error {
	intValue, err := strconv.Atoi(value)
	if err == nil {
		*i.p = append(*i.p, uint32(intValue))
	}
	return err
}

func newUint64SliceValue(val []uint64, p *[]uint64) flag.Value {
	return &uint64SliceValue{val: val, p: p}
}

// uint64SliceValue is a wrapper around a slice of uint64s that implements the flag.Value interface.
type uint64SliceValue struct {
	val []uint64
	p   *[]uint64
}

func (i *uint64SliceValue) String() string {
	return fmt.Sprintf("%v", i.val)
}

func (i *uint64SliceValue) Set(value string) error {
	intValue, err := strconv.ParseUint(value, 10, 64)
	if err == nil {
		*i.p = append(*i.p, uint64(intValue))
	}
	return err
}

func newFloat32SliceValue(val []float32, p *[]float32) flag.Value {
	return &float32SliceValue{val: val, p: p}
}

// float32SliceValue is a wrapper around a slice of float32s that implements the flag.Value interface.
type float32SliceValue struct {
	val []float32
	p   *[]float32
}

func (f *float32SliceValue) String() string {
	return fmt.Sprintf("%v", f.val)
}

func (f *float32SliceValue) Set(value string) error {
	floatValue, err := strconv.ParseFloat(value, 32)
	if err == nil {
		*f.p = append(*f.p, float32(floatValue))
	}
	return err
}

func newFloat64SliceValue(val []float64, p *[]float64) flag.Value {
	return &float64SliceValue{val: val, p: p}
}

// float64SliceValue is a wrapper around a slice of float64s that implements the flag.Value interface.
type float64SliceValue struct {
	val []float64
	p   *[]float64
}

func (f *float64SliceValue) String() string {
	return fmt.Sprintf("%v", f.val)
}

func (f *float64SliceValue) Set(value string) error {
	floatValue, err := strconv.ParseFloat(value, 64)
	if err == nil {
		*f.p = append(*f.p, float64(floatValue))
	}
	return err
}

func newIntSliceValue(val []int, p *[]int) flag.Value {
	return &intSliceValue{val: val, p: p}
}

// intSliceValue is a wrapper around a slice of ints that implements the flag.Value interface.
type intSliceValue struct {
	val []int
	p   *[]int
}

func (i *intSliceValue) String() string {
	return fmt.Sprintf("%v", i.val)
}

func (i *intSliceValue) Set(value string) error {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	*i.p = append(*i.p, intValue)
	return nil
}

// BoolSliceValue represents a slice of booleans that can be used as a flag value.
// This type implements the flag.Value interface.
func newBoolSliceValue(val []bool, p *[]bool) flag.Value {
	return &boolSliceValue{val: val, p: p}
}

// boolSliceValue is a wrapper around a slice of booleans that implements the flag.Value interface.
type boolSliceValue struct {
	val []bool
	p   *[]bool
}

func (b *boolSliceValue) String() string {
	return fmt.Sprintf("%v", b.val)
}

func (b *boolSliceValue) Set(value string) error {
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return err
	}
	*b.p = append(*b.p, boolValue)
	return nil
}

// AddFlags adds flags to the command based on the provided struct.
// The struct fields are mapped to flags using the 'name' tag for the flag name
// and the 'description' tag for the flag description.
func (c *Command) AddFlags(flags any) *Command {
	// Recursive helper function to process struct fields
	var processStruct func(value reflect.Value)
	processStruct = func(value reflect.Value) {
		typ := value.Type()
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			fieldValue := value.Field(i)

			// Recursively process embedded structs
			if field.Anonymous && fieldValue.Kind() == reflect.Struct {
				processStruct(fieldValue)
				continue
			}

			name := field.Tag.Get("name")
			description := field.Tag.Get("description")
			if name == "" || description == "" {
				continue
			}

			defaultValue := field.Tag.Get("default")
			if defaultValue != "" {
				switch fieldValue.Kind() {
				case reflect.Bool:
					if val, err := strconv.ParseBool(defaultValue); err == nil {
						fieldValue.SetBool(val)
					}
				case reflect.String:
					fieldValue.SetString(defaultValue)
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					if val, err := strconv.ParseInt(defaultValue, 10, 64); err == nil {
						fieldValue.SetInt(val)
					}
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					if val, err := strconv.ParseUint(defaultValue, 10, 64); err == nil {
						fieldValue.SetUint(val)
					}
				case reflect.Float32, reflect.Float64:
					if val, err := strconv.ParseFloat(defaultValue, 64); err == nil {
						fieldValue.SetFloat(val)
					}
				}
			}

			validateTags := field.Tag.Get("validate")
			var validators []Validator
			if validateTags != "" {
				validators = parseValidateTags(validateTags, name)
			}

			if defaultValue != "" {
				setFieldDefaultValue(fieldValue, defaultValue)
			}

			handleFieldType(c, name, description, fieldValue, validators)
		}
	}

	value := reflect.ValueOf(flags).Elem()
	processStruct(value)
	return c
}

// parseFlags parses the given flags
func (c *Command) parseFlags(args []string) error {
	// Parse flags
	tmp := os.Stderr
	os.Stderr = nil
	defer func() {
		os.Stderr = tmp
	}()

	var positionalArgs []string
	for {
		if err := c.flags.Parse(args); err != nil {
			return err
		}
		// Consume all the flags that were parsed as flags.
		args = args[len(args)-c.flags.NArg():]
		if len(args) == 0 {
			break
		}
		// There's at least one flag remaining and it must be a positional arg since
		// we consumed all args that were parsed as flags. Consume just the first
		// one, and retry parsing, since subsequent args may be flags.
		positionalArgs = append(positionalArgs, args[0])
		args = args[1:]
	}

	// Parse just the positional args so that flagset.Args()/flagset.NArgs()
	// return the expected value.
	// Note: This should never return an error.
	err := c.flags.Parse(positionalArgs)
	if err != nil {
		return err
	}

	// Validate flags with validators
	var validationErrs []error
	validationTypes := map[string]bool{}

	// Check all flags that have validations, regardless of whether they were set
	for flagName, validators := range c.flagValidations {
		for _, valid := range validators {
			// Get the actual value from the stored variable
			if valueRef, ok := c.flagVariables[flagName]; ok {
				err := valid.Validate(valueRef.Interface())
				if err != nil {
					switch err := err.(type) {
					case *validator.ValidatorError:
						err.Field = flagName
					default:
					}
					if _, ok := validationTypes[flagName]; !ok {
						validationTypes[flagName] = true
						validationErrs = append(validationErrs, err)
					}
				}
			} else {
				validationErrs = append(validationErrs, fmt.Errorf("flag '%s' not found in flagVariables", flagName))
			}
		}
	}

	if len(validationErrs) > 0 {
		return fmt.Errorf("%v", validationErrs)
	}

	if len(positionalArgs) > 0 {
		return c.parsePositionalArgs(positionalArgs)
	}
	return nil
}

// setFieldDefaultValue sets the default value for a struct field based on its type
func setFieldDefaultValue(fieldValue reflect.Value, defaultValue string) {
	switch fieldValue.Kind() {
	case reflect.Bool:
		if val, err := strconv.ParseBool(defaultValue); err == nil {
			fieldValue.SetBool(val)
		}
	case reflect.String:
		fieldValue.SetString(defaultValue)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if val, err := strconv.ParseInt(defaultValue, 10, 64); err == nil {
			fieldValue.SetInt(val)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if val, err := strconv.ParseUint(defaultValue, 10, 64); err == nil {
			fieldValue.SetUint(val)
		}
	case reflect.Float32, reflect.Float64:
		if val, err := strconv.ParseFloat(defaultValue, 64); err == nil {
			fieldValue.SetFloat(val)
		}
	}
}

// handleFieldType handles adding flags based on the struct field type
func handleFieldType(c *Command, name, description string, fieldValue reflect.Value, validators []Validator) {
	switch fieldValue.Kind() {
	case reflect.Bool:
		c.Bool(name, description, fieldValue.Addr().Interface().(*bool), validators...)
	case reflect.String:
		c.String(name, description, fieldValue.Addr().Interface().(*string), validators...)
	case reflect.Int:
		c.Int(name, description, fieldValue.Addr().Interface().(*int), validators...)
	case reflect.Int8:
		c.Int8(name, description, fieldValue.Addr().Interface().(*int8), validators...)
	case reflect.Int16:
		c.Int16(name, description, fieldValue.Addr().Interface().(*int16), validators...)
	case reflect.Int32:
		c.Int32(name, description, fieldValue.Addr().Interface().(*int32), validators...)
	case reflect.Int64:
		c.Int64(name, description, fieldValue.Addr().Interface().(*int64), validators...)
	case reflect.Uint:
		c.Uint(name, description, fieldValue.Addr().Interface().(*uint), validators...)
	case reflect.Uint8:
		c.Uint8(name, description, fieldValue.Addr().Interface().(*uint8), validators...)
	case reflect.Uint16:
		c.Uint16(name, description, fieldValue.Addr().Interface().(*uint16), validators...)
	case reflect.Uint32:
		c.Uint32(name, description, fieldValue.Addr().Interface().(*uint32), validators...)
	case reflect.Uint64:
		c.Uint64(name, description, fieldValue.Addr().Interface().(*uint64), validators...)
	case reflect.Float32:
		c.Float32(name, description, fieldValue.Addr().Interface().(*float32), validators...)
	case reflect.Float64:
		c.Float64(name, description, fieldValue.Addr().Interface().(*float64), validators...)
	case reflect.Slice:
		handleSliceType(c, name, description, fieldValue, validators)
	}
}

// handleSliceType handles adding flags for slice types
func handleSliceType(c *Command, name, description string, fieldValue reflect.Value, validators []Validator) {
	switch fieldValue.Type().Elem().Kind() {
	case reflect.Bool:
		variable := fieldValue.Addr().Interface().(*[]bool)
		c.BoolSlice(name, description, variable, validators...)
	case reflect.String:
		variable := fieldValue.Addr().Interface().(*[]string)
		c.StringSlice(name, description, variable, validators...)
	case reflect.Int:
		variable := fieldValue.Addr().Interface().(*[]int)
		c.IntSlice(name, description, variable, validators...)
	case reflect.Int8:
		variable := fieldValue.Addr().Interface().(*[]int8)
		c.Int8Slice(name, description, variable, validators...)
	case reflect.Int16:
		variable := fieldValue.Addr().Interface().(*[]int16)
		c.Int16Slice(name, description, variable, validators...)
	case reflect.Int32:
		variable := fieldValue.Addr().Interface().(*[]int32)
		c.Int32Slice(name, description, variable, validators...)
	case reflect.Int64:
		variable := fieldValue.Addr().Interface().(*[]int64)
		c.Int64Slice(name, description, variable, validators...)
	case reflect.Uint:
		variable := fieldValue.Addr().Interface().(*[]uint)
		c.UintSlice(name, description, variable, validators...)
	case reflect.Uint8:
		variable := fieldValue.Addr().Interface().(*[]uint8)
		c.Uint8Slice(name, description, variable, validators...)
	case reflect.Uint16:
		variable := fieldValue.Addr().Interface().(*[]uint16)
		c.Uint16Slice(name, description, variable, validators...)
	case reflect.Uint32:
		variable := fieldValue.Addr().Interface().(*[]uint32)
		c.Uint32Slice(name, description, variable, validators...)
	case reflect.Uint64:
		variable := fieldValue.Addr().Interface().(*[]uint64)
		c.Uint64Slice(name, description, variable, validators...)
	case reflect.Float32:
		variable := fieldValue.Addr().Interface().(*[]float32)
		c.Float32Slice(name, description, variable, validators...)
	case reflect.Float64:
		variable := fieldValue.Addr().Interface().(*[]float64)
		c.Float64Slice(name, description, variable, validators...)
	}
}
