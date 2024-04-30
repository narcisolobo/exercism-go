package twofer

import (
	"fmt"
	"strings"
)

func ShareWith(name string) string {
	n := "you"
	if len(strings.TrimSpace(name)) > 0 {
		n = name
	}
	return fmt.Sprintf("One for %s, one for me.", n)
}
