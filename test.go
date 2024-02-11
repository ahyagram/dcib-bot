package main

import (
    "fmt"
    "gobot/database"
)

func main() {
    database.Update("test.json", "name", "Queensya")
    database.Update("test.json", "age", "12")
    
    fmt.Println(database.Get("test.json", "name"))
    fmt.Println(database.Get("test.json", "class"))
}
