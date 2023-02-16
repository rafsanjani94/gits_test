package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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
	fmt.Println("input N: ")
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	input, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	var seqs []string
	for i := 0; i < int(input); i++ {
		seq := i*(i+1)/2 + 1
		seqs = append(seqs, strconv.Itoa(seq))
	}

	fmt.Print(strings.Join(seqs, "-"))
}
