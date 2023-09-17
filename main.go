package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Ash> ")
		ln, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("Error", err)
			return

		}
		if string(ln) == "q" {
			break
		}
		fmt.Println("You typed", string(ln))
	}
	fmt.Println("Exiting...")
}
