package main

import (
    "os"
    "fmt"
    "time"
    "strings"
    "encoding/json"
)

func getBaseDir() string {
    dirname, err := os.UserHomeDir()
    if err != nil {
        invalid("cannot access ~/.jot directory")
    }
    
    basedir := dirname + "/.jot"
    if !pathExists(basedir) {
        os.Mkdir(basedir, os.ModePerm)
    }

    return basedir
}

func getCategoryPath(category string) string {
    basedir := getBaseDir()
    return basedir + "/" + category + ".json"
}

func createNewCategory(args Args) {
    cpath := getCategoryPath(args.category)
    if pathExists(cpath) {
        invalid("category already exists")
    }
     
    noteCategory := NoteCategory{time.Now(), make([]Note, 0)}
    data, err := json.Marshal(noteCategory)
    if err != nil {
        invalid("failed to create category")
    }
    
    os.WriteFile(cpath, data, 0755)
}

func addToCategory(args Args) {
    cpath := getCategoryPath(args.category)
    if !pathExists(cpath) {
        invalid("category does not exist")
    }

    data, err := os.ReadFile(cpath)
    if err != nil {
        invalid("failed to add note")
    }

    noteCategory := new(NoteCategory)
    err = json.Unmarshal(data, noteCategory)
    if err != nil {
        invalid("failed to add note")
    }
    
    noteCategory.Notes = append(noteCategory.Notes, Note{time.Now(), args.content})
    data, err = json.Marshal(noteCategory)
    if err != nil {
        invalid("failed to add note")
    }

    os.WriteFile(cpath, data, 0755)
}

func listNotesInCategory(args Args) { 
    cpath := getCategoryPath(args.category)
    if !pathExists(cpath) {
        invalid("category does not exist")
    }

    data, err := os.ReadFile(cpath)
    if err != nil {
        invalid("failed to add note")
    }

    noteCategory := new(NoteCategory)
    err = json.Unmarshal(data, noteCategory)
    if err != nil {
        invalid("failed to add note")
    }

    for _, note := range noteCategory.Notes {
        fmt.Println(note.CreatedDate.Format(time.Stamp), " - ", note.Content)
    }
}   

func listCategories() {
    basedir := getBaseDir()
    entries, err := os.ReadDir(basedir)
    if err != nil {
        invalid("failed to list categories")
    }
 
    fmt.Println("Categories")
    for _, e := range entries {
        category := strings.Replace(e.Name(), ".json", "", 1)
        fmt.Println("    ", category)
    }
}

type NoteCategory struct {
    CreatedDate time.Time
    Notes []Note
}

type Note struct {
    CreatedDate time.Time
    Content string
}

