package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {

	idStr := "1"
	id, err := strconv.Atoi(idStr)
	fmt.Println(id)
	first()
	fmt.Println(err)
}
func first() error {
	return errors.New("error")
}
