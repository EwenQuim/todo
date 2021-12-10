package validator

import (
	"strings"

	"github.com/google/uuid"
)

func UUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func CleanItem(str string) string {
	splitted := strings.SplitN(str, ":", 2)

	if len(splitted) == 2 {
		return strings.ToLower(strings.TrimSpace(splitted[0])) + ": " + strings.TrimSpace(splitted[1])
	}

	return strings.TrimSpace(str)
}

func GetGroupAndContent(str string) (string, string) {
	splitted := strings.SplitN(str, ":", 2)

	if len(splitted) == 2 {
		return strings.ToLower(strings.TrimSpace(splitted[0])), strings.TrimSpace(splitted[1])
	}

	return "", strings.TrimSpace(str)
}
