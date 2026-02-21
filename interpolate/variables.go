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
			// r is next rune *after* the variable name

			// 2. look up variable name's value, all name's runes in variableName
			if value, ok := dict[variableName.String()]; ok {
				// 3. write variable's value into output string
				substituted.Write(value)
			} else {
				substituted.WriteRune('$')
				substituted.WriteString(variableName.String())
			}
			variableName.Reset()
			if r != '$' {
				findingVariable = false
				substituted.WriteRune(r)
			}
			continue
		}

		// put together variable's value by reading "identifier" runes one by one
		variableName.WriteRune(r)
	}

	if variableName.Len() > 0 {
		if value, ok := dict[variableName.String()]; ok {
			// 3. write variable's value into output string
			substituted.Write(value)
		} else {
			substituted.WriteRune('$')
			substituted.WriteString(variableName.String())
		}
	} else if findingVariable {
		substituted.WriteRune('$')
	}

	return substituted.String()
}
