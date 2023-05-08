package utils

import (
	"math/rand"
	"strconv"
)



func GenerateID() string {
	ranNum := rand.NewSource(int64(rand.Intn(100)))
	return strconv.FormatInt(ranNum.Int63(), 10)
}
