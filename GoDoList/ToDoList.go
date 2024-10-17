package main

import (
	// "bufio"
	"fmt"
	"log"
	"os"
)

var Path string

func main() {
  GetCurrentPath()
	CheckFileExistance()
	MainMenu()
}

var bIsValid bool

func MainMenu() {
	var bNotQuit bool = true
	for bNotQuit {

		// Clear the screen
		fmt.Print("\033[H\033[2J")
		PrintMainMenu()
		fmt.Println("Bool: ", bIsValid)

		// Get user input
		var userInput string
		fmt.Scan(&userInput)
		// reader := bufio.NewReader(os.Stdin)
		// sentence, _ := reader.ReadString('\n')
		IsValidInput(userInput, &bIsValid)
	}
}

func PrintMainMenu() {
	fmt.Print("1. To Do List \n")
	fmt.Print("2. Add  \n")
	fmt.Print("3. Finish \n")
	fmt.Print("4. Remove \n")
	fmt.Print("q. Quit \n")
}

func CheckFileExistance() {
	if _, err := os.Stat(Path); err == nil {
		fmt.Print("File Already Exist \n")
	} else {
		file, err := os.Create("File.txt")
		fmt.Print("File created")
		if err != nil {
			log.Fatal(err)
			file.Close()
		}
	}
}

// is valid input
var ValidInput = [6]string{"1", "2", "3", "4", "5", "q"}

func IsValidInput(Input string, bIsValid *bool) {
	*bIsValid = true
	for i := 0; i < len(ValidInput); i++ {
		if Input == ValidInput[i] {
			*bIsValid = true
			return
		}
	}
	*bIsValid = false
}

func GetCurrentPath(){
  fileName := "\\File.txt"
  myDir, err := os.Getwd()
  if (err != nil){
    fmt.Println(err)
  }
  Path = myDir + fileName
  fmt.Print("The path: ",Path)
}

// Icon
// cmd = '⌘',
// config = '🛠',
// event = '📅',
// ft = '📂',
// init = '⚙',
// keys = '🗝',
// plugin = '🔌',
// runtime = '💻',
// require = '🌙',
// source = '📄',
// start = '🚀',
// task = '📌',
// lazy = '💤 ',
