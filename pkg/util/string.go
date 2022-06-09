package util

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"net/mail"
	"regexp"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func StringWithCharset(length int, charset string) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// placeholder
func RandString(length int) string {
	return StringWithCharset(length, charset)
}

// placeholder
func RandIntString(length int) string {
	return StringWithCharset(length, "0123456789")
}

// placeholder
func Sha256String(str string) string {
	b := []byte(str)
	h := sha256.New()
	h.Write(b)
	sha256_str := hex.EncodeToString(h.Sum(nil))


	return sha256_str
}

// placeholder
func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// placeholder
func IsNumeric(str string) bool {
	m, err := regexp.MatchString(`^\d+$`, str)
	if err != nil {
		return false
	} else {
		return m
	}
}



