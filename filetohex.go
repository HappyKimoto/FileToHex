package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getUserInput(promptMessage string) string {
	// get user input
	fmt.Print(promptMessage)
	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n')
	check(err)
	// remove CR, LF, and double quotes
	userInput = strings.ReplaceAll(userInput, "\r\n", "")
	userInput = strings.ReplaceAll(userInput, "\"", "")
	// return clean user input
	return userInput
}

func main() {
	// title
	fmt.Println("=== File to Hex ===")

	// get file path
	fpIn := getUserInput("File Path: ")

	// get output size
	outSize, err := strconv.ParseInt(getUserInput("Bytes in Hex: "), 16, 64)
	check(err)

	// read file to bytes
	dat, err := os.ReadFile(fpIn)
	check(err)

	// get file size
	fInfo, err := os.Stat(fpIn)
	check(err)

	// set output size
	if outSize < 0 {
		outSize = fInfo.Size()
	} else if fInfo.Size() < outSize {
		outSize = fInfo.Size()
	}

	// loop through each byte
	const colSize = 0x10
	row := 0 - colSize
	for i := 0; i < int(outSize); i++ {

		if i%colSize == 0 {
			// add new line if column end is reached
			row += colSize
			fmt.Printf("\n%08x:0x%02x,", row, dat[i])
		} else {
			// otherwise print the current byte
			fmt.Printf("0x%02x,", dat[i])
		}
	}

}
