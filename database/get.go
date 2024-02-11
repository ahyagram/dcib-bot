package database

import "encoding/json"


func Get(fileName string, key string) string {
    var (
        errors      error
        exist       bool
        fileByte    []byte
        fileMaps    map[string]string
        value       string
    )
    
    fileByte = Read(fileName)
    errors = json.Unmarshal(fileByte, &fileMaps)
    if errors != nil {
        return "undefined"
    }
    
    value, exist = fileMaps[key]
    if exist == false {
        value = "undefined"
    }
    
    return value
}