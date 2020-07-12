package main

import "fmt"

func main() {
	mayCauseError()
}

func mayCauseError() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	panic("panic an error")
}
