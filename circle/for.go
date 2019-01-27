package main

import "fmt"

func for_circle() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

//跟 C 或者 Java 中一样，可以让前置、后置语句为空。
func for_circle_2() {
	sum := 1
	for ;sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

//基于此可以省略分号：C 的 while 在 Go 中叫做 `for`。
func for_is_gos_while() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}