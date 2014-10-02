package factfuncts

import (
	"encoding/json"
	"fmt"
	"strings"
)

// chomp record separators (\n and \r) off the end of strings
func Chomp(s string) string {
	s1 := strings.Trim(s, "\n")
	s2 := strings.Trim(s1, "\r")
	return s2
}

func Strip(s string) string {
	return strings.Replace(s, `"`, "", -1)
}

func FactToJsonString(f map[string]interface{}) string {
	b, err := json.Marshal(f)
	if err != nil {
		return fmt.Sprintf("%v", err)
	}
	return fmt.Sprintf("%s", b)
}
