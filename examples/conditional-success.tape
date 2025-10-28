# Conditional Test - Success Case
# This demonstrates a conditional that succeeds (pattern appears)

Output conditional-success.gif

Set FontSize 32
Set Width 1200
Set Height 600

Type "echo 'Hello World'"
Enter
Sleep 500ms

# Test conditional: wait for prompt (should succeed immediately)
If Wait@5s />$/
    Type "# Pattern found!"
    Enter
    Sleep 1s
Else
    Type "# Pattern not found (timeout)"
    Enter
    Sleep 1s
EndIf

Type "echo 'Test complete'"
Enter
Sleep 2s
