package main

import (
	"fmt"
	"os"
)

func main() {
	variables := os.Environ()
	for _, variable := range variables {
		fmt.Println(variable)
	}
}
