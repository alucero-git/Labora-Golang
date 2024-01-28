package main

import "fmt"

func main() {
    X:=0
	X = X + 4
	fmt.Println(X)
	X = X + X - X * X
	fmt.Println(X)
	X = X % 4     
	fmt.Println(X)
	X = (X + 4) / 2
	fmt.Println(X)
}