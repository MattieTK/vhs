# Conditional Commands in VHS

VHS now supports conditional execution of commands based on pattern matching with timeout support. This allows you to create adaptive terminal recordings that branch based on whether certain patterns appear in the output.

## Syntax

### Basic If/Else/EndIf Structure

```tape
If Wait@<timeout> /regex/
    # Commands to execute if pattern matches
Else
    # Commands to execute if timeout occurs
EndIf
```

### If Without Else

```tape
If Wait@<timeout> /regex/
    # Commands to execute only if pattern matches
EndIf
```

### Immediate Pattern Check (No Wait)

```tape
If /regex/
    # Commands to execute if pattern is currently on screen
Else
    # Commands to execute if pattern not found
EndIf
```

## How It Works

1. **Pattern Match Success**: If the Wait command finds the pattern before the timeout, the commands in the If block execute, and the Else block is skipped.

2. **Timeout/Pattern Not Found**: If the Wait command times out or the pattern isn't found, execution jumps to the Else block (if present), or continues after EndIf.

3. **Immediate Check**: Using `If /regex/` without Wait performs an immediate check on the current screen with no timeout (essentially a 0s wait).

## Examples

### Example 1: Check if a Command Succeeded

```tape
Type "which python3"
Enter

If Wait@2s /\/bin\/python3/
    Type "# Python3 found, running script..."
    Enter
    Type "python3 script.py"
    Enter
Else
    Type "# Python3 not found, installing..."
    Enter
    Type "sudo apt install python3"
    Enter
EndIf
```

### Example 2: Conditional File Operations

```tape
Type "ls myfile.txt"
Enter
Sleep 500ms

If Wait@3s /myfile.txt/
    Type "# File exists, creating backup"
    Enter
    Type "cp myfile.txt myfile.bak"
    Enter
Else
    Type "# File not found, creating new file"
    Enter
    Type "touch myfile.txt"
    Enter
EndIf
```

### Example 3: Branching on Command Output

```tape
Type "npm --version"
Enter

If Wait@2s /[0-9]+\.[0-9]+/
    Type "# npm is installed, proceeding..."
    Enter
    Type "npm install"
    Enter
Else
    Type "# npm not found, please install Node.js first"
    Enter
EndIf
```

### Example 4: Multiple Conditions

```tape
Type "echo 'test 1'"
Enter

If Wait@1s /test 1/
    Type "# Test 1 passed"
    Enter
EndIf

Type "echo 'test 2'"
Enter

If Wait@1s /test 2/
    Type "# Test 2 passed"
    Enter
Else
    Type "# Test 2 failed"
    Enter
EndIf
```

## Wait Options in Conditionals

The Wait command within If supports all standard Wait options:

- **Timeout**: `@2s`, `@500ms`, `@1m` - How long to wait for the pattern
- **Scope**:
  - `Wait /regex/` - Default, checks current line (equivalent to `Wait+Line`)
  - `Wait+Line /regex/` - Checks only the current line
  - `Wait+Screen /regex/` - Checks the entire visible screen

### Examples with Scope

```tape
# Wait for pattern on current line only
If Wait+Line@3s />$/
    Type "# Found prompt on current line"
    Enter
EndIf

# Wait for pattern anywhere on screen
If Wait+Screen@5s /ERROR/
    Type "# Error detected on screen"
    Enter
Else
    Type "# No errors found"
    Enter
EndIf
```

## Nested Conditionals

Conditionals can be nested to create complex branching logic:

```tape
Type "uname -s"
Enter

If Wait@2s /Linux/
    Type "# Detected Linux"
    Enter
    Type "cat /etc/os-release"
    Enter
    Sleep 500ms

    If Wait@2s /Ubuntu/
        Type "# Ubuntu detected, using apt"
        Enter
        Type "sudo apt update"
        Enter
    Else
        Type "# Other Linux distro"
        Enter
    EndIf
Else
    Type "# Not Linux, checking for macOS"
    Enter
    Type "sw_vers"
    Enter

    If Wait@2s /macOS/
        Type "# macOS detected, using brew"
        Enter
        Type "brew update"
        Enter
    Else
        Type "# Unknown OS"
        Enter
    EndIf
EndIf
```

## Best Practices

1. **Set Appropriate Timeouts**: Choose timeouts based on expected command execution time. Too short may cause false negatives; too long slows down your tape.

2. **Use Specific Patterns**: Make your regex patterns specific enough to avoid false matches but general enough to handle variations.

3. **Provide Feedback**: Use Type commands to show what path was taken for better understanding when viewing the recording.

4. **Test Both Paths**: Create separate tape files to test both success and failure cases.

5. **Combine with Hide/Show**: Use Hide before conditionals for setup, and Show when ready to record:

```tape
Hide
# Setup that shouldn't be recorded
Type "export TEST_VAR='value'"
Enter
Show

If Wait@2s /success/
    Type "# Test passed!"
    Enter
Else
    Type "# Test failed!"
    Enter
EndIf
```

## Error Handling

- **Missing EndIf**: Parse error - "If without matching EndIf"
- **Multiple Else Blocks**: Parse error - "Multiple Else blocks for single If"
- **Orphaned Else/EndIf**: Parse error - "Else/EndIf without matching If"
- **Invalid Condition**: Runtime error - condition must be a valid Wait command

## Limitations

- Currently, only Wait commands are supported as conditions
- Variables and arithmetic expressions are not supported
- Loop constructs (for, while) are not available

## See Also

- [Wait Command Documentation](https://github.com/charmbracelet/vhs#wait)
- [VHS Main Documentation](https://github.com/charmbracelet/vhs)
- Example files: `examples/conditional-*.tape`
