/* Go. Тестовое задание. */

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No path to the file!")
		fmt.Println("Please, restart program with it!")
		return
	} else {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer file.Close()
		fileScanner := bufio.NewScanner(file)
		linesCounter := uint64(0)
		for fileScanner.Scan() {
			linesCounter++
		}
		fmt.Println(linesCounter)
	}
}
