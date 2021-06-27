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
	var buffer strings.Builder

	for i := 0; i < len(x); i++ {
		buffer.WriteString(x[len(x)-i-1] + " ")
	}

	newLine := buffer.String()
	fmt.Println(newLine)

}
