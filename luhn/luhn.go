package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	isEven := len(id)%2 == 0
	if len(id) <= 1 {
		return false
	}

	sum := 0
	for i := len(id) - 1; i >= 0; i-- {
		num, err := strconv.Atoi(string(id[i]))
		if err != nil {
			return false
		}
		if isEven {
			if i%2 == 0 {
				num *= 2
			}
		} else {
			if i%2 != 0 {
				num *= 2
			}
		}
		if num > 9 {
			num -= 9
		}
		sum += num
	}

	return sum%10 == 0
}
