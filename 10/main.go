package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	floats := make([]float64, len(x))

	for i, v := range x {
		floats[i], _ = strconv.ParseFloat(v, 64)
	}

	mapka := make(map[int][]float64)

	for _, v := range floats {
		mapka[(int(v)/10)*10] = append(mapka[(int(v)/10)*10], v)
	}

	for k, v := range mapka {
		fmt.Printf("key: %d, value: %v\n", k, v)
	}
}
