package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {

	idStr := "1"
	id, _ := strconv.Atoi(idStr)
	fmt.Println(id)
	err:=first()
fmt.Println(err)
}
func first() error {
	return 			errors.New("error")
}
