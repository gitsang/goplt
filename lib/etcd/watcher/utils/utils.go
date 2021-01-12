package utils

import "math/rand"

const (
	chars = "abcdefghijklmnopqrstuvwxyz0123456789"
)

func RandomString(l uint) string {
	s := make([]byte, l)
	for i := 0; i < int(l); i++ {
		s[i] = chars[rand.Intn(len(chars))]
	}
	return string(s)
}
