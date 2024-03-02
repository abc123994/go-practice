package calc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
func cmd2math(cmd string) (math []string, priority_operation_count int32) {

	parser := []string{}
	priority_operation_count = 0
	for _, e := range cmd {
		parser = append(parser, string(e))
	}
	tmp := ""

	for i, e := range parser {
		if e != "+" && e != "*" && e != "/" && e != "-" {
			tmp = tmp + e
		} else {
			math = append(math, tmp)
			math = append(math, e)
			if e == "*" || e == "/" {
				priority_operation_count += 1
			}
			tmp = ""
		}
		if i == len(parser)-1 {
			math = append(math, tmp)
		}
	}
	return
}

type input struct {
	Cmd string
}
type out struct {
	Result float64 `json:"result"`
}

func Calc(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {

		c.JSON(http.StatusForbidden, "")
	}
	data := input{}
	out := out{}
	json.Unmarshal(jsonData, &data)
	math, priority_operation_count := cmd2math(data.Cmd)

	log.Println("convert cmd =>", math)
	for priority_operation_count > 0 {
		for i, e := range math {
			if e == "*" || e == "/" {
				_left_num, _ := strconv.ParseFloat(math[i-1], 64)
				_right_num, _ := strconv.ParseFloat(math[i+1], 64)
				_left_legacy := math[0 : i-1]
				_right_legacy := math[(i + 2):]
				var num float64
				if e == "*" {
					num = prod(_left_num, _right_num)

				} else if e == "/" {
					num = div(_left_num, _right_num)
				}
				math = _left_legacy
				math = append(math, fmt.Sprintf("%v", num))
				math = append(math, _right_legacy...)
				log.Println(math)
				priority_operation_count--
				break
			}
		}
	}
	for len(math) > 1 {

		for i, e := range math {
			if e == "+" || e == "-" {
				_left_num, _ := strconv.ParseFloat(math[i-1], 64)
				_right_num, _ := strconv.ParseFloat(math[i+1], 64)
				_left_legacy := math[0 : i-1]
				_right_legacy := math[(i + 2):]
				var num float64
				if e == "+" {
					num = add(_left_num, _right_num)
				} else if e == "-" {
					num = minus(_left_num, _right_num)
				}
				math = _left_legacy
				math = append(math, fmt.Sprintf("%v", num))
				math = append(math, _right_legacy...)
				log.Println(math)
				break

			}
		}
	}

	res, _ := strconv.ParseFloat(math[0], 64)
	out.Result = res
	c.JSON(http.StatusOK, out)
}
