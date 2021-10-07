package tools

import "regexp"

func Grep(re string, str string) bool {
	return regexp.MustCompile(re).Match([]byte(str))
}
