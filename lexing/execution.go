package lexing

import (
	"fmt"
	"minimal-self-replication/interpolate"
	"strings"
)

func Execute(command string, variableValues map[string][]byte) {
	if strings.HasPrefix(command, "echo") {
		fields := strings.SplitN(command, " ", 2)
		if len(fields) != 2 {
			return
		}
		fmt.Printf("%s", stringEval(fields[1], variableValues))
		fmt.Println()
		return
	}

	fields := strings.SplitN(command, "=", 2)
	if len(fields) != 2 {
		return
	}

	variableValues[fields[0]] = []byte(stringEval(fields[1], variableValues))
}

func stringEval(original string, variableValues map[string][]byte) string {
	if original[0] == '\'' {
		return strings.Trim(original, "'")
	}

	if original[0] == '"' {
		return interpolate.Variables(strings.Trim(original, "\""), variableValues)
	}

	return interpolate.Variables(original, variableValues)
}
