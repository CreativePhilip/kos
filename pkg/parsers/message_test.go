package parsers

import (
	"fmt"
	"log"
	"testing"
)

func mockValidMessage() string {
	return "field1 Int  \nfield2 String"
}

func TestParseMessage(t *testing.T) {
	response, err := ParseMessage(mockValidMessage())

	expectedOutput := map[string]string{
		"field1": "Int",
		"field2": "String",
	}

	if err != nil {
		log.Printf(err.Error())
		t.Fatal("Error not nil in ParseMessage")
	}

	for key, value := range expectedOutput {
		if response[key] != value {
			t.Errorf(fmt.Sprintf("Invalid field parsing value[%s] shoule be equal %#v but its %#v",
				key, value, response[key]))
		}
	}
}

func TestGoCodeGenFromMessage(t *testing.T) {
	response, _ := ParseMessage(mockValidMessage())

	msg := MessageData{
		name:           "TestMsg",
		topLevelFields: response,
	}

	code := GoCodeGenFromMessage(msg)
	fmt.Println(code)
}
