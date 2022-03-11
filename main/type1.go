package main

import "fmt"

type Node struct {
}

func IsNil(v *struct{}) bool {
	return v == nil
}
func IsNilNode(v *Node) bool {
	return v == nil
}

func main() {
	var a *Node
	fmt.Println(a == interface{}(a))   //true
	fmt.Println(nil == interface{}(a)) //false

	fmt.Println(IsNil(a))     //false
	fmt.Println(IsNilNode(a)) //false
}
