# Conditional Test - Complex Example
# This demonstrates nested conditionals and realistic use cases

Output conditional-complex.gif

Set FontSize 28
Set Width 1400
Set Height 700

Type "# Testing conditional commands"
Enter
Sleep 1s

# Test 1: Check if a file exists
Type "echo 'test content' > /tmp/testfile.txt"
Enter
Sleep 500ms

Type "ls /tmp/testfile.txt"
Enter
Sleep 500ms

If Wait@3s /testfile/
    Type "# File found! Creating backup..."
    Enter
    Sleep 500ms
    Type "cp /tmp/testfile.txt /tmp/testfile.bak"
    Enter
    Sleep 500ms
Else
    Type "# File not found, creating new one..."
    Enter
    Sleep 500ms
    Type "echo 'new content' > /tmp/testfile.txt"
    Enter
    Sleep 500ms
EndIf

Sleep 1s

# Test 2: Check command output
Type "date"
Enter
Sleep 500ms

If Wait@2s /202/
    Type "# Date command succeeded!"
    Enter
    Sleep 1s
Else
    Type "# Date command failed"
    Enter
    Sleep 1s
EndIf

Type "echo 'All tests complete'"
Enter
Sleep 2s
