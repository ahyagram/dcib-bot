package database

import (
    "fmt"
    "os"
)


func Read(fileName string) []byte {
    var (
        errors      error
        fileByte    []byte
        filePath    string
    )
    
    filePath = fmt.Sprintf("./database/db/%s", fileName)
    fileByte, errors = os.ReadFile(filePath)
    if errors != nil {
        fileByte = []byte("{}")
    }
    
    return fileByte
}