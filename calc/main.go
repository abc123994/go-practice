package main

import (
	"log"
	"strconv"
	"strings"
)

func add(a float64, b float64) (res float64) {
	res = a + b
	return
}
func minus(a float64, b float64) (res float64) {
	res = a - b
	return
}
func div(a float64, b float64) (res float64) {
	res = a / b
	return
}
func prod(a float64, b float64) (res float64) {
	res = a * b
	return
}
func calc(cmd string) (res float64) {
	//1 parse the cmd
	//2 pick the operation
	//3 convert left and right number
	numbers := strings.Split(cmd, "+")
	if len(numbers) > 1 {
		left, _ := strconv.ParseFloat(numbers[0], 64)
		right, _ := strconv.ParseFloat(numbers[1], 64)
		res = add(left, right)
		return
	}
	numbers = strings.Split(cmd, "-")
	if len(numbers) > 1 {
		left, _ := strconv.ParseFloat(numbers[0], 64)
		right, _ := strconv.ParseFloat(numbers[1], 64)
		res = minus(left, right)
		return
	}
	numbers = strings.Split(cmd, "*")
	if len(numbers) > 1 {
		left, _ := strconv.ParseFloat(numbers[0], 64)
		right, _ := strconv.ParseFloat(numbers[1], 64)
		res = prod(left, right)
		return
	}
	numbers = strings.Split(cmd, "/")
	if len(numbers) > 1 {
		left, _ := strconv.ParseFloat(numbers[0], 64)
		right, _ := strconv.ParseFloat(numbers[1], 64)
		res = div(left, right)
		return
	}
	return

}
func main() {
	test_case := "10+3"
	output := calc(test_case)
	log.Println(test_case, "=", output)
	test_case = "2-3"
	output = calc(test_case)
	log.Println(test_case, "=", output)
	test_case = "2*3"
	output = calc(test_case)
	log.Println(test_case, "=", output)
	test_case = "2/3"
	output = calc(test_case)
	log.Println(test_case, "=", output)
}
