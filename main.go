package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {

	idStr := "1"
	id, err := strconv.Atoi(idStr)
	fmt.Println(id,err)
	newError()
}
func newError() error {
	return errors.New("error")
}
