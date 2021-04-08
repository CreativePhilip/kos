package parsers

import (
	"errors"
	"fmt"
	"strings"
)

type MessageData struct {
	name           string
	topLevelFields map[string]string
}

func ParseError(lineIdx int) error {
	return errors.New(fmt.Sprintf("Invalid line nr %d", lineIdx))
}

func InvalidTypeError(item string) error {
	return errors.New(fmt.Sprintf("Not recognised type value %s", item))
}

func ParseMessage(data string) (map[string]string, error) {
	lines := strings.Split(data, "\n")
	parsedData := map[string]string{}

	for idx, line := range lines {
		tokens := strings.Split(strings.Trim(line, " "), " ")

		if len(tokens) != 2 {
			return nil, ParseError(idx + 1)
		}

		// TODO: Allow for nested messages in the future
		if !IsValidKey(tokens[1]) {
			return nil, InvalidTypeError(tokens[1])
		}

		parsedData[tokens[0]] = tokens[1]

	}
	return parsedData, nil
}

func GoCodeGenFromMessage(message MessageData) string {
	code := ""

	// FIXME: Get a valid package name, this requires a plan for the project structure
	code += "package %s \n"
	code += "\n\n"

	code += fmt.Sprintf("type %s struct {\n", message.name)

	for key, value := range message.topLevelFields {
		typeString, _ := ConvertTypeToGo(value)
		code += fmt.Sprintf("\t%s %s\n", key, typeString)
	}

	code += "}"

	return code
}
