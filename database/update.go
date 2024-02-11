package database

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
)


func Update(
    fileName    string,
    key         string,
    value       string,
) {
    var (
        errors       error
        file         *os.File
        fileByte     []byte
        fileMaps     map[string]string
        filePath     string
        jsonByte     []byte
        overwrite    *os.File
    )
    
    filePath = fmt.Sprintf("./database/db/%s", fileName)
    file, errors = os.Open(filePath)
    if errors != nil {
        file = Create(fileName)
        errors = nil
    }
    
    defer file.Close()
    fileByte = Read(fileName)
    
    errors = json.Unmarshal(fileByte, &fileMaps)
    if errors != nil {
        log.Fatal(errors)
    }
    
    fileMaps[key] = value
    jsonByte, errors = json.MarshalIndent(fileMaps, "", "    ")
    
    overwrite, _ = os.Create(filePath)
    defer overwrite.Close()
    overwrite.WriteString(string(jsonByte))
}