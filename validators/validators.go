package validators

import (
	"strings"
)

func NameValidator(name string) string {
	if len(name) > 3 {
		return name
	}
	return "failed"

}
func EmailValidator(email string) string {

	if strings.Contains(email, "@") {
		return strings.ToLower(email)
	} else {
		return "failed"
	}
}
