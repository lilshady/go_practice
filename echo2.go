package main

import "os"
import "fmt"

func main() {
	result := ""
	var sep string
	for index,s := range os.Args[1:]{
		result += sep + s
		sep = "*"
		fmt.Println(index,s)
	}
	fmt.Println(result)
}
