package libraries

import "regexp"

func RegexCheckNumeric(id string) bool {
	regex := regexp.MustCompile(`^[0-9]+$`)
	if !regex.MatchString(id) {
		return false
	}

	return true
}
