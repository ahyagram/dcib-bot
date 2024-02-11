package database

import (
    "fmt"
    "log"
    "os"
)


func Create(fileName string) *os.File {
    var (
        errors      error
        file        *os.File
        filePath    string
    )
    
    filePath = fmt.Sprintf("./database/db/%s", fileName)
    file, errors = os.Create(filePath)
    if errors != nil {
        log.Fatal(errors)
    }
    
    file.WriteString("{}")
    return file
}