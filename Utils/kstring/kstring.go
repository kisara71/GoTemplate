package kstring

import (
	"errors"
	"fmt"
)

func ToInt(num string) (int, error) {
	if num == "" {
		return 0, errors.New("empty string")
	}
	sign := 1
	res := 0
	start := 0
	if num[0] == '-' {
		sign = -1
		start = 1
	} else if num[0] == '+' {
		start = 1
	}

	for ; start < len(num); start++ {
		if num[start] < '0' || num[start] > '9' {
			return 0, fmt.Errorf("invalid character: %c", num[start])
		}
		res = res*10 + int(num[start]-'0')
	}
	return res * sign, nil
}
