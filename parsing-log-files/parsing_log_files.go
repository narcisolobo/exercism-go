package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[-~=*]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	re := regexp.MustCompile(`"[^"]*password[^"]*"`)
	for _, line := range lines {
		lowercase := strings.ToLower(line)
		matches := re.FindAllString(lowercase, -1)
		count += len(matches)
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	var taggedLines []string
	re := regexp.MustCompile(`User(\s+\S+)`)

	for _, line := range lines {
		if matches := re.FindStringSubmatch(line); matches != nil {
			user := strings.TrimSpace(matches[1])
			taggedLine := fmt.Sprintf("[USR] %s %s", user, line)
			taggedLines = append(taggedLines, taggedLine)
		} else {
			taggedLines = append(taggedLines, line)
		}
	}

	return taggedLines
}
