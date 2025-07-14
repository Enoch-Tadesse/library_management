package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
	"library_management/controllers"
	"library_management/services"
)

var l services.LibraryManager = services.NewLibrary()

func main() {

	reader := bufio.NewReader(os.Stdin)
loop:
	controllers.DisplayMenu()
	option, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("failed to read option")
		goto loop
	}
	option = strings.TrimSpace(option)
	if option == "q" {
		fmt.Println("Library Closed")
		return
	}

	item, exist := controllers.Menu[option]
	if !exist {
		fmt.Println("Invalid Value")
		goto loop
	}

	item.Action(l)
	goto loop
}
