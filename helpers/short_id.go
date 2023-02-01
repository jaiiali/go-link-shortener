package helpers

import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

const charset = "0123456789" +
	"abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func NewShortID() string {
	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

	length, err := strconv.Atoi(os.Getenv("APP_SHORT_ID_LENGTH"))
	if err != nil {
		panic(err)
	}

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}
