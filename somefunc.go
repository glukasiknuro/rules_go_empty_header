package main

import "C"

import "fmt"
import _ "golang.org/x/sys/unix"

//export SomeFunc
func SomeFunc() {
    fmt.Println("Hello from SomeFunc")
}

func main() {
    fmt.Println("Hello from somefunc main function")
}
