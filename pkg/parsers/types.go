package parsers

func ConvertTypeToGo(typeString string) (string, error) {
	switch typeString {
	case "Int":
		return "int", nil
	case "String":
		return "string", nil
	default:
		return "", nil
	}
}

func IsValidKey(key string) bool {
	switch key {
	case "Int", "String":
		return true
	}
	return false
}
