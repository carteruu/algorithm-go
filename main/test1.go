package main

import "fmt"

func main() {
	var m map[string]interface{} //A
	m["a"] = 1
	if v := m["b"]; v != nil { //B
		fmt.Println(v)
	}

}
