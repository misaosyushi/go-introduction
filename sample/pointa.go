package sample

import "fmt"

func Pointa() {
	a, b := 1,1

	double(a, &b)

	fmt.Println(a)
	fmt.Println(b)
}

func double(x int, y *int)  {
	x = x * 2
	*y = *y * 2
}
