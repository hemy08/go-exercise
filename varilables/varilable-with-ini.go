package main


import "fmt"

var i, j int = 1, 2

func With_iniitalizers() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
