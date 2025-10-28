# Conditional Test - Timeout Case
# This demonstrates a conditional that times out (pattern doesn't appear)

Output conditional-timeout.gif

Set FontSize 32
Set Width 1200
Set Height 600

Type "echo 'Starting test...'"
Enter
Sleep 500ms

# Test conditional: wait for a pattern that won't appear
If Wait@2s /NONEXISTENT_PATTERN/
    Type "# This should not appear"
    Enter
    Sleep 1s
Else
    Type "# Timeout occurred - pattern not found!"
    Enter
    Sleep 1s
EndIf

Type "echo 'Test complete'"
Enter
Sleep 2s
