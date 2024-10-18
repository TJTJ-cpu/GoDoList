package main

import (
	"bufio"
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

var (
	bIsValid bool
	bNotQuit bool = true
)

func MainMenu() {
	for bNotQuit {

		// Clear the screen
		ClearScreen()
		PrintMainMenu()

		// Get user input
		var userInput string
		fmt.Scan(&userInput)
		SwitchScene(userInput)
	}
}

func SwitchScene(userInput string) {
	IsValidInput(userInput, &bIsValid)
	if bIsValid == false {
		fmt.Println("Bool: ", bIsValid)
		return
	}
	if userInput == "q" {
		os.Exit(0)
	} else if userInput == "1" {
		DisplayFile()
	} else if userInput == "2" {
		AddToList()
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

// IS VALID INPUT
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

func GetCurrentPath() {
	fileName := "\\File.txt"
	myDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	Path = myDir + fileName
}

func AddToList() {
	// Wait for enter
	Continue("Enter your to do list: ")
	reader := bufio.NewReader(os.Stdin)
	sentence, _ := reader.ReadString('\n')

	// Remove the ^m
	sentence = string([]rune(sentence))[:len(sentence)-1]
	sentence += "\n"
	file, err := os.OpenFile("File.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if _, err = file.WriteString(sentence); err != nil {
		log.Fatal(err)
	}
    // Wait for user input
	read2 := bufio.NewReader(os.Stdin)
	_, _ = read2.ReadString('\n')
	reader2 := bufio.NewReader(os.Stdin)
	_, err2 := reader2.ReadString('\n')
	if err2 != nil {
		log.Fatal(err)
	}
}

func DisplayFile() {
	file := ReadFile()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	ClearScreen()
	println("This is your to do lists")
	for i := 0; i < len(lines); i++ {
		fmt.Print(i+1, ": ", lines[i], "\n")
	}
	Continue("Press anything to continue...")
}

func Continue(output string) {
	// ClearScreen()
	println(output)
	read2 := bufio.NewReader(os.Stdin)
	_, _ = read2.ReadString('\n')
	reader := bufio.NewReader(os.Stdin)
	_, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func ReadFile() *os.File {
	file, err := os.Create("File.txt")
	fmt.Print("File created")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return file
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
