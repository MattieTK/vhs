package main

import (
	"fmt"
	"strings"
)

// Supported shells of VH.
const (
	bash       = "bash"
	cmdexe     = "cmd"
	fish       = "fish"
	nushell    = "nu"
	osh        = "osh"
	powershell = "powershell"
	pwsh       = "pwsh"
	xonsh      = "xonsh"
	zsh        = "zsh"
)

// DefaultPromptColor is the default color for the shell prompt.
const DefaultPromptColor = "#5B56E0"

// PromptMarker is an invisible OSC escape sequence embedded in shell prompts
// to detect when the shell has rendered a new prompt (i.e. is ready for input).
const PromptMarker = "\x1b]7777;\x07"

// Shell is a type that contains a prompt and the command to set up the shell.
type Shell struct {
	Name string
}

// ShellConfig returns the shell configuration with the given prompt color.
func ShellConfig(name, promptColor string) (env []string, command []string) {
	// Parse hex color to RGB components
	r, g, b := hexToRGB(promptColor)
	hexNoHash := strings.TrimPrefix(promptColor, "#")

	switch name {
	case bash:
		return []string{
				fmt.Sprintf("PS1=\\[\\e]7777;\\a\\]\\[\\e[38;2;%d;%d;%dm\\]> \\[\\e[0m\\]", r, g, b),
				"BASH_SILENCE_DEPRECATION_WARNING=1",
			},
			[]string{"bash", "--noprofile", "--norc", "--login", "+o", "history"}
	case zsh:
		return []string{fmt.Sprintf(`PROMPT=%%{`+PromptMarker+`%%}%%F{#%s}> %%F{reset_color}`, hexNoHash)},
			[]string{"zsh", "--histnostore", "--no-rcs"}
	case fish:
		return nil, []string{
			"fish",
			"--login",
			"--no-config",
			"--private",
			"-C", "function fish_greeting; end",
			"-C", fmt.Sprintf(`function fish_prompt; printf '\e]7777;\a'; set_color %s; echo -n "> "; set_color normal; end`, hexNoHash),
		}
	case powershell:
		return nil, []string{
			"powershell",
			"-NoLogo",
			"-NoExit",
			"-NoProfile",
			"-Command",
			fmt.Sprintf(`Set-PSReadLineOption -HistorySaveStyle SaveNothing; function prompt { [Console]::Write([char]27 + ']7777;' + [char]7); Write-Host '>' -NoNewLine -ForegroundColor ([System.Drawing.Color]::FromArgb(%d,%d,%d)); return ' ' }`, r, g, b),
		}
	case pwsh:
		return nil, []string{
			"pwsh",
			"-Login",
			"-NoLogo",
			"-NoExit",
			"-NoProfile",
			"-Command",
			fmt.Sprintf(`Set-PSReadLineOption -HistorySaveStyle SaveNothing; Function prompt { [Console]::Write([char]27 + ']7777;' + [char]7); Write-Host -ForegroundColor ([System.Drawing.Color]::FromArgb(%d,%d,%d)) -NoNewLine '>'; return ' ' }`, r, g, b),
		}
	case cmdexe:
		return nil, []string{"cmd.exe", "/k", fmt.Sprintf("prompt=$E]7777;$E\\$E[38;2;%d;%d;%dm^> $E[0m", r, g, b)}
	case nushell:
		return nil, []string{"nu", "--execute", fmt.Sprintf(`$env.PROMPT_COMMAND = {print -n ($"\e]7777;\u{07}"); $"\e[;38;2;%d;%d;%dm>\e[m "}; $env.PROMPT_COMMAND_RIGHT = {''}`, r, g, b)}
	case osh:
		return []string{fmt.Sprintf("PS1=\\[\\e]7777;\\a\\]\\[\\e[38;2;%d;%d;%dm\\]> \\[\\e[0m\\]", r, g, b)},
			[]string{"osh", "--norc"}
	case xonsh:
		return nil, []string{"xonsh", "--no-rc", "-D", fmt.Sprintf("PROMPT=\033]7777;\007\033[;38;2;%d;%d;%dm>\033[m ", r, g, b)}
	default:
		return nil, nil
	}
}

// hexToRGB converts a hex color string to RGB components.
func hexToRGB(hex string) (r, g, b int) {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) == 6 {
		fmt.Sscanf(hex, "%02x%02x%02x", &r, &g, &b)
	}
	return
}

// Shells contains a mapping from shell names to their Shell struct.
var Shells = map[string]Shell{
	bash:       {Name: bash},
	zsh:        {Name: zsh},
	fish:       {Name: fish},
	powershell: {Name: powershell},
	pwsh:       {Name: pwsh},
	cmdexe:     {Name: cmdexe},
	nushell:    {Name: nushell},
	osh:        {Name: osh},
	xonsh:      {Name: xonsh},
}
