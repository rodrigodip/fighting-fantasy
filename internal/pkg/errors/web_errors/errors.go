package weberrors

import (
	"fmt"
	"strings"
)

func ParseErrorForWeb(err string) string {
	parsedErr := strings.Split(err, ":")
	return fmt.Sprintf("%s: %s", parsedErr[len(parsedErr)-2], parsedErr[len(parsedErr)-1])
}
