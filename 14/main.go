package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("insert string: ")

	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		return
	}

	line = strings.TrimSuffix(line, "\n")

	x := strings.Split(line, " ")
	set := make(map[string]struct{})

	for _, v := range x {
		set[v] = struct{}{}
	}

	fmt.Print("Set: ")
	for k, _ := range set {
		fmt.Print(k, " ")
	}
}
