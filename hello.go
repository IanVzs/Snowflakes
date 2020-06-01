package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Printf("hello, world\n")
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	fmt.Println(cmp.Diff("Hello World", "Hello World"))
}
