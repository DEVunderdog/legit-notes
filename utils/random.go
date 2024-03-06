package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const specialCharacters = "@#$&*"

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max - min + 1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}



func PasswordGenerator() string {
	var sb strings.Builder
	k := len(alphabet)
	m := len(specialCharacters)
	for i := 0; i < 8; i++ {
		c := alphabet[rand.Intn(k)] + specialCharacters[rand.Intn(m)]
		sb.WriteByte(c)
	}

	return sb.String()
	
}

func RandomUser() (string, string){
	username := RandomString(8)
	email := RandomString(8)
	email += "@gmail.com"

	return username, email
}

func RandomPassword() string {
	return PasswordGenerator()
}

func RandomNote() (string, string) {
	title := RandomString(10)
	description := RandomString(15)
	return title, description
}

