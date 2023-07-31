package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func menuPrint() {
	fmt.Println("Please select an option")
	fmt.Println("1) Print menu")

	in := bufio.NewReader(os.Stdin)

	ch, _ := in.ReadString('\n')

	ch = strings.TrimSpace(ch)
}
