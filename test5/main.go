package main

import (
	"bufio"
	"fmt"
	"io"
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
	// fmt.Println("input: ")
	// reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	// input := strings.TrimSpace(readLine(reader))

	var rooms map[string][]string
	rooms = map[string][]string{
		"engineer": []string{
			"devi", "eko", "irianto", "eka", "bagus", "putranto",
		},
		"product": []string{
			"bayu", "eko", "putranto", "adi", "persada", "bagus",
		},
	}

	membersInput := []string{"putranto", "adi"}
	var output []string
	for room, members := range rooms {
		exist := 0
		for _, member := range members {
			for _, input := range membersInput {
				if input == member {
					exist++
				}
			}
		}

		if exist == len(membersInput) {
			output = append(output, room)
		}
	}

	fmt.Println(output)
}
