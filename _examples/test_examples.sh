#!/bin/bash

echo "Testing cliz examples with parameter validations..."
echo "============================================="

SCRIPT_DIR=$(dirname "$(realpath "$0")")
TMP_DIR=$(mktemp -d)

test_example() {
    local dir_name=$SCRIPT_DIR/$1
    local dir_base=$1
    local test_name=$2
    local cmd_path="$TMP_DIR/$dir_name"

    echo ""
    echo "Testing $test_name..."
    echo "Example: $dir_name"
    
    if go build -o "$cmd_path" "$dir_name/main.go"; then
        echo "✓ Build successful"
        
        if "$cmd_path" --help > /dev/null 2>&1; then
            echo "✓ Help command works"
        else
            echo "✗ Help command failed"
        fi
        
        case "$dir_base" in
            "basic")
                if "$cmd_path" --name "John" --age 25 > /dev/null 2>&1; then
                    echo "✓ Basic parameters validation passed"
                else
                    echo "✗ Basic parameters validation failed"
                fi
                ;;
            "subcommands")
                if "$cmd_path" --name "John" --age 25 > /dev/null 2>&1; then
                    echo "✓ Root command validation passed"
                else
                    echo "✗ Root command validation failed"
                fi
                if "$cmd_path" greet --name "Alice" > /dev/null 2>&1; then
                    echo "✓ Greet subcommand validation passed"
                else
                    echo "✗ Greet subcommand validation failed"
                fi
                if "$cmd_path" farewell --name "Bob" > /dev/null 2>&1; then
                    echo "✓ Farewell subcommand validation passed"
                else
                    echo "✗ Farewell subcommand validation failed"
                fi
                ;;
            "struct")
                if "$cmd_path" --name "John" --age 30 --email "john@example.com" --password "Password1234" --score 85 --status "active" --website "https://example.com" > /dev/null 2>&1; then
                    echo "✓ Struct flags validation passed"
                else
                    echo "✗ Struct flags validation failed"
                fi
                ;;
            "inherit")
                if "$cmd_path" --name "John" --age 25 > /dev/null 2>&1; then
                    echo "✓ Inherit flags root command validation passed"
                else
                    echo "✗ Inherit flags root command validation failed"
                fi
                if "$cmd_path" greet --name "Alice" --age 30 > /dev/null 2>&1; then
                    echo "✓ Inherit flags greet command validation passed"
                else
                    echo "✗ Inherit flags greet command validation failed"
                fi
                if "$cmd_path" farewell --name "Bob" --age 35 > /dev/null 2>&1; then
                    echo "✓ Inherit flags farewell command validation passed"
                else
                    echo "✗ Inherit flags farewell command validation failed"
                fi
                ;;
            "ints")
                if "$cmd_path" --int8 10 --int16 100 --int32 1000 --uint 10 --uint8 20 --uint16 30 --uint32 40 > /dev/null 2>&1; then
                    echo "✓ Integer flags validation passed"
                else
                    echo "✗ Integer flags validation failed"
                fi
                ;;
            "floats")
                if "$cmd_path" --float64 3.14 > /dev/null 2>&1; then
                    echo "✓ Floating point flag validation passed"
                else
                    echo "✗ Floating point flag validation failed"
                fi
                if "$cmd_path" --float64s 1.5 --float64s 2.5 --float64s 3.5 > /dev/null 2>&1; then
                    echo "✓ Floating point slice validation passed"
                else
                    echo "✗ Floating point slice validation failed"
                fi
                ;;
            "slices")
                if "$cmd_path" --bools true --bools false --strings "apple" --strings "banana" --ints 10 --ints 20 --int8s 1 --int8s 2 --int16s 2 --int16s 3 --int32s 3 --int32s 4 --int64s 4 --int64s 5 --uints 5 --uints 6 --uint8s 6 --uint8s 7 --uint16s 7 --uint16s 8 --uint32s 8 --uint32s 9 --uint64s 9 --uint64s 10 > /dev/null 2>&1; then
                    echo "✓ Slice flags validation passed"
                else
                    echo "✗ Slice flags validation failed"
                fi
                ;;
            "banner")
                if "$cmd_path" --name "John" --age 25 > /dev/null 2>&1; then
                    echo "✓ Custom banner flags validation passed"
                else
                    echo "✗ Custom banner flags validation failed"
                fi
                ;;
            "hidden")
                if "$cmd_path" --name "John" > /dev/null 2>&1; then
                    echo "✓ Default command validation passed"
                else
                    echo "✗ Default command validation failed"
                fi
                if "$cmd_path" greet --name "Alice" --verbose > /dev/null 2>&1; then
                    echo "✓ Greet command validation passed"
                else
                    echo "✗ Greet command validation failed"
                fi
                if "$cmd_path" farewell --name "Bob" > /dev/null 2>&1; then
                    echo "✓ Farewell command validation passed"
                else
                    echo "✗ Farewell command validation failed"
                fi
                ;;
            "prerun")
                if "$cmd_path" --name "John Doe" --age 25 --email "john@example.com" --password "StrongP@ss123" --port 8080 --username "userjohn" > /dev/null 2>&1; then
                    echo "✓ PreRun validation chain passed"
                else
                    echo "✗ PreRun validation chain failed"
                fi
                ;;
            "advanced")
                if "$cmd_path" --phone "+1-555-123-4567" --zip "12345" --gender "male" --color "blue" --country "USA" --message "hello world" --code "APP123" --email "john@example.com" --url "https://example.com" --password "Str0ngP@ssw0rd" --username "johndoe" --username-pattern "user01" > /dev/null 2>&1; then
                    echo "✓ Advanced validators validation passed"
                else
                    echo "✗ Advanced validators validation failed"
                fi
                ;;
            "positional")
                if "$cmd_path" "source.txt" "dest.txt" "copy" --force > /dev/null 2>&1; then
                    echo "✓ Positional args validation passed"
                else
                    echo "✗ Positional args validation failed"
                fi
                ;;
            "mixed")
                if "$cmd_path" --string "hello" --int 42 --float 3.14 --strings "apple" --strings "banana" --ints 1 --ints 2 --ints 3 > /dev/null 2>&1; then
                    echo "✓ Mixed flags validation passed"
                else
                    echo "✗ Mixed flags validation failed"
                fi
                ;;
            "inheritance")
                if "$cmd_path" --name "John" --age 25 > /dev/null 2>&1; then
                    echo "✓ Command inheritance root validation passed"
                else
                    echo "✗ Command inheritance root validation failed"
                fi
                if "$cmd_path" server start --name "John" --age 25 --server-name "webserver" --port 8080 > /dev/null 2>&1; then
                    echo "✓ Command inheritance server start validation passed"
                else
                    echo "✗ Command inheritance server start validation failed"
                fi
                if "$cmd_path" server stop --name "John" --age 25 --server-name "webserver" > /dev/null 2>&1; then
                    echo "✓ Command inheritance server stop validation passed"
                else
                    echo "✗ Command inheritance server stop validation failed"
                fi
                if "$cmd_path" user --name "Alice" create > /dev/null 2>&1; then
                    echo "✓ Command inheritance user create validation passed"
                else
                    echo "✗ Command inheritance user create validation failed"
                fi
                ;;
            "struct_tags")
                if "$cmd_path" --name "John" --age 30 --debug --custom "test" > /dev/null 2>&1; then
                    echo "✓ Flag struct tags validation passed"
                else
                    echo "✗ Flag struct tags validation failed"
                fi
                ;;
            *)
                echo "⚠ No custom validation defined for $dir_name"
                ;;
        esac
        
        rm "$cmd_path"
        echo "✓ Cleanup successful"
    else
        echo "✗ Build failed"
    fi
    
    echo ""
}

rm -rf "$TMP_DIR"

test_example "basic" "Basic flags"
test_example "subcommands" "Subcommands"
test_example "struct" "Struct flags"
test_example "inherit" "Inherit flags"
test_example "ints" "Integer flag types"
test_example "floats" "Floating point flag types"
test_example "slices" "Slice flag types"
test_example "banner" "Custom banner and error handler"
test_example "hidden" "Hidden and default commands"
test_example "prerun" "PreRun and validation chain"
test_example "advanced" "Advanced validators"
test_example "positional" "Positional arguments with struct"
test_example "mixed" "Mixed flag types"
test_example "inheritance" "Command inheritance"
test_example "struct_tags" "Flag struct tags"

echo "============================================="
echo "Testing completed!"
echo ""