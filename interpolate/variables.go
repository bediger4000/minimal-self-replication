package interpolate

import (
	"strings"
	"unicode"
)

func Variables(originalString string, dict map[string][]byte) string {
	var substituted strings.Builder

	findingVariable := false
	runes := []rune(originalString)
	max := len(runes)

	var variableName strings.Builder

	for i := 0; i < max; i++ {
		r := runes[i]

		if !findingVariable && r == '$' {
			findingVariable = true
			continue
		}

		if !findingVariable {
			substituted.WriteRune(r)
			continue
		}

		if !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_') {
			// r contains the next rune *after* the variable name

			// Look up variable name's value, all name's runes now written to variableName
			if value, ok := dict[variableName.String()]; ok {
				// Write variable's value into output string
				substituted.Write(value)
			}
			// No interpolated bytes for a variable not found
			variableName.Reset()
			if r != '$' {
				findingVariable = false
				substituted.WriteRune(r)
			}
			continue
		}

		// Put together variable's value by reading "identifier" runes one by one
		variableName.WriteRune(r)
	}

	if variableName.Len() > 0 {
		if value, ok := dict[variableName.String()]; ok {
			// Write variable's value into output string
			substituted.Write(value)
		}
		// No interpolated bytes for a variable not found
	}

	return substituted.String()
}
