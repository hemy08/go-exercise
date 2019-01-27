package main

import "fmt"

var c, python, java bool

//运行 ：go run  main.go base_types.go short-declarations.go  varilable-with-ini.go
func main() {
	var i int
	fmt.Println(i, c, python, java)

	With_iniitalizers()
	Short_declarations()
	Base_types()
}
