package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("insert string: ")

	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		return
	}

	runes := []rune(line)

	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-1-i] = runes[len(runes)-1-i], runes[i]
	}

	line = string(runes)
	fmt.Println(line)
}
