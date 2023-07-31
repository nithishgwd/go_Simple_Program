package main

import (
	"fmt"
	"strings"
)

func struct_loop() {
	for _, item := range menuData() {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, price)
		}
	}
}
