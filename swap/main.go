package main

import "log"

type swap_case struct {
	a int32
	b int32
}

func (test_case *swap_case) swap() {
	c := test_case.a
	test_case.a = test_case.b
	test_case.b = c

}
func main() {
	test_case := swap_case{a: 1, b: 2}
	log.Println(test_case.a, test_case.b)
	test_case.swap()
	log.Println(test_case.a, test_case.b)
}
