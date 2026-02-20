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

		// not reading through a string literal (single- or double-quoted)

		if r == '\'' || r == '"' {
			cmd.WriteRune(r)
			quoteChar = r
			inQuotedString = true
			continue
		}

		// ; terminates a command when not reading a quoted string
		if r == ';' {
			sendString(outCh, cmd)
			continue
		}

		cmd.WriteRune(r)
	}

	// might hit end-of-file without finding end-of-command
	sendString(outCh, cmd)

	close(outCh)
}

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
