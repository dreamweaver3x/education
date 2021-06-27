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

	unique := make(map[rune]struct{})
	line = strings.TrimSuffix(line, "\n")

	runes := []rune(line)

	for _, v := range runes {
		_, ok := unique[v]
		if ok {
			fmt.Println("symbols are not unique")
			return
		}
		unique[v] = struct{}{}
	}

	fmt.Println("symbols are unique")

}
