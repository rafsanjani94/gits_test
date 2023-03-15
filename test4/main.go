package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("input: ")
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	input := strings.TrimSpace(readLine(reader))

	input = strings.ReplaceAll(input, "-", "")
	input = strings.ReplaceAll(input, " ", "")

	if len(input) < 4 {
		fmt.Println(input)
		return
	}

	var result []string
	for {
		// fmt.Println(len(input))
		if len(input) == 0 {
			break
		}

		nSub := 3
		if len(input) == 2 || len(input) == 4 {
			nSub = 2
		}
		block := string([]rune(input)[:nSub])
		result = append(result, block)

		input = strings.TrimLeft(input, block)
		// fmt.Println(block)
		// fmt.Println(input)
		// break
	}

	// fmt.Println(result)
	fmt.Println(strings.Join(result, "-"))
}
