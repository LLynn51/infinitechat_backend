package utils

import (
	"fmt"
	"math"
	"math/rand"
)

func RandomNumber(len int) string {
	//go Go 1.20 之前版本
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//return fmt.Sprintf("%06d", r.Intn(1000000))

	max := int(math.Pow10(len)) // 10^digits
	return fmt.Sprintf("%0*d", len, rand.Intn(max))
}
