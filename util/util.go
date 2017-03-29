package util

import "math/rand"

var randomStringRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

//GetRandomString returns a random string of given range
func GetRandomString(length int) string {
	r := make([]rune, length)
	for i := range r {
		r[i] = randomStringRunes[rand.Intn(len(randomStringRunes))]
	}
	return string(r)
}
