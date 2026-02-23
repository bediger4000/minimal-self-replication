package lexing

import (
	"strings"
)

func NewCommandLexer(runes []rune) func() string {
	outCh := make(chan string)
	go doCommandLexing(outCh, runes)
	return func() string {
		command := <-outCh
		return command
	}
}

// doCommandLexing runs in its own goroutine.
// It collects runes into "commands", which are either
// distinct lines in the input or ';' terminated substrings
// of lines in the input. Any quoted string literals of
// the interpreted language retain their quote characters.
func doCommandLexing(outCh chan string, runes []rune) {
	var inQuotedString bool
	var quoteChar rune
	cmd := &strings.Builder{}

	for _, r := range runes {

		if r == '\n' {
			sendString(outCh, cmd)
			continue
		}

		if inQuotedString {
			cmd.WriteRune(r)
			if r == quoteChar {
				inQuotedString = false
			}
			continue
		}

		// Not reading through a string literal (single- or double-quoted)

		if r == '\'' || r == '"' {
			cmd.WriteRune(r)
			quoteChar = r
			inQuotedString = true
			continue
		}

		// ; terminates a command when not reading a quoted string.
		if r == ';' {
			sendString(outCh, cmd)
			continue
		}

		cmd.WriteRune(r)
	}

	// Might hit end-of-file without finding end-of-command.
	sendString(outCh, cmd)

	close(outCh)
}

// sendString creates a string from the strings.Builder
// argument, and writes it to the output channel.
// Never writes a zero-length string, because that
// signifies end-of-input.
func sendString(outCh chan string, bldr *strings.Builder) {
	if bldr.Len() == 0 {
		return
	}

	str := strings.TrimSpace(bldr.String())

	bldr.Reset()

	if str == "" {
		return
	}

	outCh <- str
}
