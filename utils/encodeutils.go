package utils

import (
	"math"
	"strings"
)

const (
	base = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b    = 62
)

//Function encodes given database ID to base62 string

func ToBase62(num int) string {
	r := num % b
	res := string(base[r])
	div := num / 62
	q := int(math.Floor(float64(div)))

	for q != 0 {
		r = q % b
		temp := q / b
		q = int(math.Floor(float64(temp)))
		res = string(base[int(r)]) + res
	}

	return string(res)
}

//Function that decodes base62 string to database ID

func ToBase10(str string) int {
	res := 0
	for _, r := range str {
		res = (b * res) + strings.Index(base, string(r))
	}
	return res
}
