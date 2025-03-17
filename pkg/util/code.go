package util

import (
	"math/rand"
)

func GenerateCode() string {
	rand.NewSource(1)
	digits := "0123456789"
	code := ""
	for i := 0; i < 6; i++ {
		code += string(digits[rand.Intn(10)])
	}
	return code
}