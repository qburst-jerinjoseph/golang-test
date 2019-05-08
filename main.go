package main

import (
	"fmt"
	"strconv"
)

func main() {

	idStr := "1"
	id, err := strconv.Atoi(idStr)
	fmt.Println(id)
}
