package main

import "fmt"

var gg string

func main() {
	gg = "this is global vairable but initalizied in main function"
	// variables that are defined and initialized explicitly
	var deo int = 20
	var bl bool = true
	var st string = "Go language"

	fmt.Printf("Global :%s \n Interger: %d and its type : %T \n Boolean : %t and its type : %T \n string: %s and its type : %T \n", gg, deo, deo, bl, bl, st, st)

	// variables that are implicitly initalized

	deoint := 20
	deobl := false
	deost := "Program"
	fmt.Printf("Integer: %d and its type : %T \n Boolean: %t and its type : %T \n String: %s and its type : %T ", deoint, deoint, deobl, deobl, deost, deost)

}
