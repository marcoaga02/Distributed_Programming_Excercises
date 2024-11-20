package gotodo

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add       string
	Del       int
	Edit      string
	Completed int
	List      string //bool
	//OutputFormat string // New field for the output format
}

const (
	TableList = "table"
	CSVList   = "csv"
)

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	// here we set all types of flag and at the end, we can use flag.Usage() to show all of them
	flag.StringVar(&cf.Add, "add", "", "Add a new todo, give me the title")
	flag.IntVar(&cf.Del, "remove", -1, "Remove a todo, give me its id")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo, give me the index and a new title. id:new_title")
	flag.IntVar(&cf.Completed, "complete", -1, "Specify the index of the todo to mark as completed")
	//flag.BoolVar(&cf.List, "list", false, "List all the todos")
	flag.StringVar(&cf.List, "list", "", "List all the todos in the table or csv format specifying the type after the flag")
	//flag.StringVar(&cf.OutputFormat, "output", "table", "Specify the output format for the list: table or csv") // New flag

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Run(todos *Todos) {
	switch {
	case cf.List != "":
		err := todos.Render()

		if err != nil {
			fmt.Printf("Error: %s", err)
			os.Exit(1)
		}
	case cf.Add != "":
		todos.Add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)

		if len(parts) != 2 {
			fmt.Printf("Error: invalid format for edit. Use id:new_title")
			os.Exit(1)
		}

		id, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Printf("Error: invalid id for edit. Use id:new_title")
			os.Exit(1)
		}

		err = todos.Edit(id, parts[1])

		if err != nil {
			fmt.Printf("Error: %s. Use id:new_title", err)
			os.Exit(1)
		}
	case cf.Completed != -1:
		err := todos.Complete(cf.Completed)
		if err != nil {
			fmt.Printf("Error: %s", err)
			os.Exit(1)
		}
	default:
		flag.Usage()
		os.Exit(1)
	}
}
