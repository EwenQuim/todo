package validator

import (
	"strings"

	"github.com/EwenQuim/todo-app/database"
	"github.com/google/uuid"
)

func UUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func CleanItem(str string, s database.Service) string {
	r := s.Regex.FindStringSubmatch(str)

	if len(r) >= 2 {
		return strings.TrimSpace(r[1]) + ": " + strings.TrimSpace(r[2])
	}
	return str

}
