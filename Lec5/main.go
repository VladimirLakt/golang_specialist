package main

import (
	"fmt"
	"strings"
)

func main() {
	for {
		var newLine string

		fmt.Scan(&newLine)
		if strings.Compare(strings.TrimSpace(newLine), "") == 0 {
			return
		}
		fmt.Println(newLine)

	}

}
