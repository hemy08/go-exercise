package main

import "fmt"

func add(x, y int) int {
	return x + y
}

/*
函数可以返回任意数量的返回值。

swap 函数返回了两个字符串。
 */
func swap(x, y string) (string, string) {
	return y, x
}

/*
Go 的返回值可以被命名，并且像变量那样使用。
返回值的名称应当具有一定的意义，可以作为文档使用。
没有参数的 return 语句返回结果的当前值。也就是`直接`返回。
直接返回语句仅应当用在像下面这样的短函数中。在长的函数中它们会影响代码的可读性。
 */
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

/*
var 语句定义了一个变量的列表；跟函数的参数列表一样，类型在后面。
就像在这个例子中看到的一样，`var` 语句可以定义在包或函数级别。
 */
var c, python, java bool

func main() {
	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))

	var i int
	fmt.Println(i, c, python, java)
}