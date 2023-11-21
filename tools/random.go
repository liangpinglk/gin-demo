package tools

import (
	"fmt"
	"math/rand"
)

func RandomString(strLen int) string {
	var result string
	for i := 0; i < strLen; i++ {
		num := rand.Intn(122-97) + 97
		result += fmt.Sprint(num)
	}
	return result

}
