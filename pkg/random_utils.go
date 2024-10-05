package random_utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func RandomString(lenght int) string {

	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < lenght; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner's name
func RandomOwner() string {
	return RandomString(7)
}

// RandomMoney generates a random amount of money
func RandomMoney() int {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
