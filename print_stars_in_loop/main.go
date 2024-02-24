package main

import (
	"fmt"
)

func print_stars(height int32) {
	//等腰三角?
	/*
			   *
			  ***
			 *****
			 height 3
			 space  2  1  0
			 stars  1  3  5


			    *
			   ***
		      *****
		  	 *******
		 	*********
			 height 5
			 space  4  3  2  1  0
			 stars  1  3  5  7  9
	*/
	stars := 1
	for y := 0; y < int(height); y++ {
		space := height - int32(y) - 1 //print spacce
		for i := 0; i < int(space); i++ {
			fmt.Print(" ")
		}
		//print start stars

		for z := 0; z < stars; z++ {
			fmt.Print("*")
		}
		fmt.Println("")
		stars += 2
	}

}

func main() {
	test_case := 10
	print_stars(int32(test_case))

}
