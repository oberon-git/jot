package main

import (
    "fmt"
    "os"
)

func main() {
    args := parse_args()
    switch args.action {
    case New:
        createNewCategory(args)
    case Add:
        addToCategory(args)
    case List:
        if args.category == "" {
            listCategories()
        } else {
            listNotesInCategory(args)
        }
    }
}

type Action int
const (
    New Action = iota
    Add
    List
)

type Args struct {
    action Action
    category string
    content string
}

func parse_args() Args {
    var content string
    positional := make([]string, 0, len(os.Args))
    for i := 1; i < len(os.Args); {
        switch os.Args[i] {
        case "--help", "-h":
            fmt.Println("usage: jot [-h] [-n content] ACTION CATEGORY")
            fmt.Println("actions: new | add | list")
            fmt.Println("category: the category of the note")
            os.Exit(0)
        case "--note", "-n":
            if i + 1 >= len(os.Args) {
                fmt.Println("usage: jot [-h] [-n content] ACTION CATEGORY")
                invalid("note was not provided")
            }

            if content != "" {
                fmt.Println("usage: jot [-h] [-n content] ACTION CATEGORY")
                invalid("more than one note was provided")
            }

            content = os.Args[i + 1]
            i += 2
        default:
            positional = append(positional, os.Args[i])
            i += 1
        }
    }
    
    if len(positional) < 1 {
        fmt.Println("usage: jot [-h] [-n content] ACTION CATEGORY")
        invalid("action or category is missing")
    }
    
    var action Action
    switch positional[0] {
    case "new":
        action = New
    case "add":
        action = Add
    case "list":
        action = List
    default:
        invalid("new, add, and remove are the only accepted actions")
    }
    
    var category string
    if len(positional) == 2 {
        category = positional[1]
    } else if action != List {
        fmt.Println("usage: jot [-h] [-n content] ACTION CATEGORY")
        invalid("action or category is missing")
    }

    if action != Add && content != "" {
        invalid("only the add action can have a non empty note content")
    } else if action == Add && content == "" {
        fmt.Print("Note: ")
        fmt.Scan(content)
        if content == "" {
            invalid("must provide non empty note")
        }
    }

    return Args{action: action, category: category, content: content}
}

