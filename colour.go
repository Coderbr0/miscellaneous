package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main2() {
	// outputFile := strings.SplitAfter(os.Args[2], "--color=")

	// if len(os.Args) != 3 || len(outputFile) != 2 || outputFile[1] == "" {
	// 	fmt.Println("Usage: go run . [STRING] [OPTION]\n\nEX: go run . something --color=<color>")
	// 	return
	// }

	// if len(os.Args) != 2 {
	// 	fmt.Println("There is one and only one argument allowed - a", "\"string inside quotes\"")
	// 	return
	// }
	// if os.Args[1] == "" {
	// 	return
	// }
	// if os.Args[1] == "\\n" {
	// 	fmt.Println()
	// 	return
	// }

	color := os.Args[1]
	value := 31
	// picker = could be

	switch color {
	case "black":
		value = 30
	case "red":
		value = 31
	case "green":
		value = 32
	case "yellow":
		value = 33
	case "blue":
		value = 34
	case "magenta":
		value = 35
	case "cyan":
		value = 36
	case "white":
		value = 37
	}

	file, err := os.Open("standard.txt")
	if err != nil {
		log.Fatalf("Is the standard.txt banner file present? Are you sure?: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	line := []string{}
	bannerMap := make(map[int][]string)
	i := 0
	z := 32
	for scanner.Scan() {
		line = append(line, scanner.Text())
		i++
		if i == 9 {
			bannerMap[z] = line
			line = []string{}
			z++
			i = 0
		}
	}
	file.Close()

	strArt := strings.Split(os.Args[1], "\\n")

	for m := 0; m < len(strArt); m++ {
		if strArt[m] == "" {
			fmt.Printf("\n")
		} else {
			for n := 1; n <= 8; n++ {
				for t := range strArt[m] {
					fmt.Printf("\033[%vm%v\033[0m", value, bannerMap[int(strArt[m][2])][n])
					fmt.Printf("\033[%vm%v\033[0m", value, bannerMap[int(strArt[m][t])][n])
					// fmt.Printf("\x1b[38;2;255;0;0m%v", bannerMap[int(strArt[m][t])][n])
				}
				fmt.Printf("\n")
			}
		}
	}
}
