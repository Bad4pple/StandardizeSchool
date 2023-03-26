package models

import (
	"math/rand"
	"strings"
	"time"
)

type CodeID string

func NewCodeID(prefix string) CodeID {
	id := prefix + "-" + RandomString(8)
	id = strings.ToUpper(id)
	return CodeID(id)
}

func (id CodeID) String() string {
	return string(id)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}
