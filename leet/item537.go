package leet

import (
	"fmt"
	"strconv"
	"strings"
)

func complexNumberMultiply(num1 string, num2 string) string {
	arr1 := strings.Split(num1, "+")
	arr2 := strings.Split(num2, "+")
	real1, _ := strconv.Atoi(arr1[0])
	real2, _ := strconv.Atoi(arr2[0])
	imaginary1, _ := strconv.Atoi(arr1[1][0 : len(arr1[1])-1])
	imaginary2, _ := strconv.Atoi(arr2[1][0 : len(arr2[1])-1])

	real := real1*real2 + -imaginary1*imaginary2
	imaginary := real1*imaginary2 + real2*imaginary1
	return fmt.Sprintf("%d+%di", real, imaginary)
}
