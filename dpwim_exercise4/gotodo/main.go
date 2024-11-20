package main

import (
	"fmt"
	"gotodo/internal/gotodo"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err)
		os.Exit(1)
	}
	filename := os.Getenv("TODO_FILENAME") // we can set this env variable from terminal or from the .env file
	fmt.Println(filename)
	var storage gotodo.Storage
	switch {
	case filepath.Ext(filename) == ".json": // if the extension of the file is json
		storage = gotodo.NewJsonStorage("todos.json")
	case filepath.Ext(filename) == ".gob": // if the extension of the file is gob
		storage = gotodo.NewGobStorage("todos.gob")
	default:
		fmt.Fprintf(os.Stderr, "Unsupported storage format: %s\n")
		os.Exit(1)
	}

	command := gotodo.NewCmdFlags()

	var render gotodo.Render
	if command.List != "" {
		switch command.List {
		case gotodo.TableList:
			render = gotodo.NewTableRender(os.Stdout)
		case gotodo.CSVList:
			render = gotodo.NewCsvRender(os.Stdout)
		default:
			fmt.Fprintf(os.Stderr, "Unsupported list format: %s\n", command.List)
			os.Exit(1)
		}
	}

	todos := gotodo.NewTodos(render, storage)

	if err := todos.Load(); err != nil { // if the file is empty, this function returns an EOF error
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	command.Run(todos)

	if err := todos.Save(); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}
