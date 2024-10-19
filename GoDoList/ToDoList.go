package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

var Path string

func main() {
	GetCurrentPath()
	EnsureFileExists("File.txt")
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
	if userInput == "1" {
		DisplayToDoList()
	} else if userInput == "2" {
		AddToList()
	} else if userInput == "q" {
		os.Exit(0)
	}
}

func PrintMainMenu() {
	MainMenu := table.NewWriter()
	MainMenu.SetTitle("Welcome To Go-Do List")
	MainMenu.SetOutputMirror((os.Stdout))
	MainMenu.AppendRow(table.Row{"1", "Display ToDo List"})
	MainMenu.AppendRow(table.Row{"2", "Add"})
	MainMenu.AppendRow(table.Row{"3", "Mark Finish"})
	MainMenu.AppendRow(table.Row{"4", "Remove"})
	MainMenu.AppendRow(table.Row{"q", "Quit"})
	MainMenu.Render()
	println()
	fmt.Print("User Input: ")
}

func EnsureFileExists(filename string) {
	if _, err := os.Stat(Path); err == nil {
		fmt.Print("File Already Exist \n")
	} else {
		file, err := os.Create(filename)
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

func DisplayToDoList() {
	DisplayTaskTable()
	Continue("Press anything to continue...")
}

func AddToList() {
	ClearScreen()
	fmt.Print("Enter your to do list: ")

	// Flush Input
	FlushInput()

	// User Input
	reader := bufio.NewReader(os.Stdin)
	sentence, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// Remove newline
	sentence = strings.TrimSpace(sentence)
	CreatedDate := GetCurrentTime()

	// Format Task
	taskLine := fmt.Sprint(sentence, " | NotDone | ", CreatedDate, "\n")

	// Open File
	file, err := os.OpenFile("File.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write the sentence to file
	if _, err = file.WriteString(taskLine); err != nil {
		log.Fatal(err)
	}
}

func RemoveTask() {
	DisplayTaskTable()
	// Open file
	f := ReadFile()
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		text := scanner.Text()
		lines = append(lines, text)
	}

	FlushInput()
	fmt.Print("Enter the number of the task to remove: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') 

	var userInput int
	FlushInput()
    _, err3 := fmt.Sscanf(input, "%d", &userInput)
    if err3 != nil{
        log.Fatal(err3)
    }

	if 0 < userInput && userInput < len(lines) {
		lines = append(lines[:userInput-1], lines[userInput:]...)
	}

	// Write the updated list
	err2 := os.WriteFile("File.txt", []byte(strings.Join(lines, "\n")), 0644)
	if err2 != nil {
		log.Fatal(err2)
	}
	Continue("Press anything to continue...")
}

func Continue(output string) {
	FlushInput()
	println()
	println(output)
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
	file, err := os.Open("File.txt")
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func DisplayTaskTable() {
	ClearScreen()
	f := ReadFile()
	defer f.Close()

	MainTable := table.NewWriter()
	MainTable.AppendHeader(table.Row{"Tasks", "Status"})
	MainTable.SetTitle("To Do List")
	MainTable.SetOutputMirror(os.Stdout)
	MainTable.SetAutoIndex(true)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		Result := strings.Split(line, "|")
		MainTable.AppendRow(table.Row{Result[0], Result[1]})
		// Use this to separate "Not Done" and "Done"
		// MainTable.AppendSeparator()
	}
	MainTable.Render()
}

func FlushInput() {
	var discard string
	fmt.Scanln(&discard)
}

func GetCurrentTime() string {
	currentTime := time.Now()
	formattedDate := currentTime.Format("02 Jan 2006")
	return formattedDate
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
