package utils

import(
	"encoding/json"
	"strings"
)

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "  ")
	return string(s)
}

func SanitizeFileName(name string) string {
    invalidChars := []string{"/", "\\", ":", "*", "?", "\"", "<", ">", "|"}
	for _, char := range invalidChars {
		name = strings.ReplaceAll(name, char, "-")
	}
	return name
}

func IsNilPointer[T any](p *T) bool {
    return p == nil
}
