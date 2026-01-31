package cliz

import (
	"reflect"
	"strconv"
)

// parsePositionalArgs parses the positional arguments and stores them in the provided structure.
// The structure fields are mapped to positional arguments using the 'position' tag.
// The position tag specifies the index of the positional argument to map to the field.
func (c *Command) parsePositionalArgs(positionalArgs []string) error {
	if c.positionalArgsMap == nil {
		return nil
	}

	for i, arg := range positionalArgs {
		pos := strconv.Itoa(i)
		if valueRef, ok := c.positionalArgsMap[pos]; ok {
			switch valueRef.Kind() {
			case reflect.String:
				valueRef.SetString(arg)
			case reflect.Int:
				intValue, err := strconv.Atoi(arg)
				if err == nil {
					valueRef.SetInt(int64(intValue))
				}
			case reflect.Int8:
				intValue, err := strconv.Atoi(arg)
				if err == nil {
					valueRef.SetInt(int64(intValue))
				}
			case reflect.Int16:
				intValue, err := strconv.Atoi(arg)
				if err == nil {
					valueRef.SetInt(int64(intValue))
				}
			case reflect.Int32:
				intValue, err := strconv.Atoi(arg)
				if err == nil {
					valueRef.SetInt(int64(intValue))
				}
			case reflect.Int64:
				intValue, err := strconv.Atoi(arg)
				if err == nil {
					valueRef.SetInt(int64(intValue))
				}
			case reflect.Uint:
				// Handle negative numbers by setting to 0 (default for invalid uint conversion)
				if uintValue, err := strconv.ParseUint(arg, 10, 64); err == nil {
					valueRef.SetUint(uintValue)
				}
			case reflect.Uint8:
				// Handle negative numbers by setting to 0 (default for invalid uint conversion)
				if uintValue, err := strconv.ParseUint(arg, 10, 64); err == nil {
					valueRef.SetUint(uintValue)
				}
			case reflect.Uint16:
				// Handle negative numbers by setting to 0 (default for invalid uint conversion)
				if uintValue, err := strconv.ParseUint(arg, 10, 64); err == nil {
					valueRef.SetUint(uintValue)
				}
			case reflect.Uint32:
				// Handle negative numbers by setting to 0 (default for invalid uint conversion)
				if uintValue, err := strconv.ParseUint(arg, 10, 64); err == nil {
					valueRef.SetUint(uintValue)
				}
			case reflect.Uint64:
				// Handle negative numbers by setting to 0 (default for invalid uint conversion)
				if uintValue, err := strconv.ParseUint(arg, 10, 64); err == nil {
					valueRef.SetUint(uintValue)
				}
			case reflect.Float32:
				floatValue, err := strconv.ParseFloat(arg, 64)
				if err == nil {
					valueRef.SetFloat(floatValue)
				}
			case reflect.Float64:
				floatValue, err := strconv.ParseFloat(arg, 64)
				if err == nil {
					valueRef.SetFloat(floatValue)
				}
			case reflect.Bool:
				boolValue, err := strconv.ParseBool(arg)
				if err == nil {
					valueRef.SetBool(boolValue)
				}
			}
		}
	}

	return nil
}

// AddPositionalArgs adds positional arguments to the command based on the provided struct.
// The struct fields are mapped to positional arguments using the 'position' tag.
// The position tag specifies the index of the positional argument to map to the field.
func (c *Command) AddPositionalArgs(argsStruct any) *Command {
	value := reflect.ValueOf(argsStruct).Elem()
	typ := value.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := value.Field(i)

		position := field.Tag.Get("position")
		if position == "" {
			continue
		}

		c.positionalArgsMap[position] = fieldValue
	}

	return c
}
