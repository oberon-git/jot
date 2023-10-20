package main

import (
    "fmt"
    "os"
    "flag"
)

func main() {
    args := parse_args()
    switch args.action {
        if 
}

func new_note(args Args) {
        
}

type Args struct {
    action string
    category string
    content string
}

func parse_args() Args {
    action := flag.String("a", "add", "action to take") 
    category := flag.String("c", "misc", "category of the note")
    content := flag.String("s", "", "content of note")

    flag.Parse()
      
    if *action != "add" && *content != "" {
        fmt.Println("Error: only the add action can have a non empty note content")
        os.Exit(1)
    } else if *action == "add" && *content == "" {
        fmt.Print("Note: ")
        fmt.Scan(content)
        if *content == "" {
            fmt.Println("Error: Must provide non empty note")
            os.Exit(1)
        }
    }

    return Args{action: *action, category: *category, content: *content}
}

