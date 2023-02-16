package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Input struct {
	player  int
	skors   []int
	game    int
	results []int
}

var myInput Input

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

func setInput(key string) {
	fmt.Println("input " + key)
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	switch key {
	case "player", "game":
		inputPlayer, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		if key == "player" {
			myInput.player = int(inputPlayer)
		} else {
			myInput.game = int(inputPlayer)
		}
	case "skors", "results":
		inputSkors := strings.TrimSpace(readLine(reader))
		skors := strings.Split(inputSkors, " ")
		for _, skor := range skors {
			skorInt, err := strconv.ParseInt(skor, 10, 64)
			checkError(err)
			if key == "skors" {
				myInput.skors = append(myInput.skors, int(skorInt))
			} else {
				myInput.results = append(myInput.results, int(skorInt))
			}
		}
	}
}

func sortSkor(skors []int) []int {
	sort.Slice(skors, func(i, j int) bool {
		return skors[j] < skors[i]
	})

	return skors
}

func uniqueSkor(skors []int) []int {
	existSkor := make(map[int]bool)
	listSkor := []int{}
	for _, skor := range skors {
		if _, value := existSkor[skor]; !value {
			existSkor[skor] = true
			listSkor = append(listSkor, skor)
		}
	}

	return listSkor
}

func getRank(skors []int, results []int) {
	for _, result := range results {
		rank := 1
		for _, skor := range skors {
			if result >= skor {
				fmt.Print(rank)
				fmt.Print(" ")
				break
			}

			rank++
		}

		if result < skors[len(skors)-1] {
			fmt.Print(rank)
			fmt.Print(" ")
		}
	}
}

func main() {
	inputs := []string{"player", "skors", "game", "results"}
	for _, input := range inputs {
		setInput(input)
	}

	if myInput.player != len(myInput.skors) {
		fmt.Println("daftar skor pemain tidak sama")
		return
	}

	if myInput.game != len(myInput.results) {
		fmt.Println("hasil skor permainan tidak sama")
		return
	}

	myInput.skors = sortSkor(myInput.skors)
	myInput.skors = uniqueSkor(myInput.skors)
	getRank(myInput.skors, myInput.results)
}
