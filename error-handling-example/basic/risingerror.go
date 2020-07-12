package main

import (
	"errors"
	"fmt"
)

func main() {
	err := eat("pill")
	if err != nil {
		fmt.Println(err)
	}
}

func eat(thing string) error {
	if thing == "pill" {
		return errors.New("you can not eat pill")
	} else {
		fmt.Println("you're eating " + thing)

	}
	return nil
}
