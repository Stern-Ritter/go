package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Stern-Ritter/go/hw02_fix_app/printer"
	"github.com/Stern-Ritter/go/hw02_fix_app/reader"
)

func main() {
	consoleReader := bufio.NewReader(os.Stdin)
	replacer := strings.NewReplacer("\r", "", "\n", "")
	var path string

	fmt.Print("Enter data file path: ")
	path, err := consoleReader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	path = replacer.Replace(path)
	if len(path) == 0 {
		path = "data.json"
	}

	staff, err := reader.ReadJSON(path)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	printer.PrintStaff(staff)
}
