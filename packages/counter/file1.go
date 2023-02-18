package counter

import (
	"strconv"
)

func CountStrings(A, sim, B string) string{
	var a, b int
	var err error
	a, err = strconv.Atoi(A)
	if err != nil{
		return ""
	}
	b, err = strconv.Atoi(B)
	if err != nil{
		return ""
	}

	if sim == "+"{
		return strconv.Itoa(a + b)
	}
	if sim == "-"{
		return strconv.Itoa(a - b)
	}
	return ""
}
