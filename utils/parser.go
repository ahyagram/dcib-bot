package utils

import (
    "strings"
)


type Parser struct {
    Prefix      []string
    Text        string
    Username    string
}


func NewParser(text string) *Parser {
    return &Parser {
        Prefix  : []string {"/", "!"},
        Text    : text,
        Username: strings.ToLower("@DetectiveDCIB_Bot"),
    }
}

func (p *Parser) Command() string {
    var (
        isPrefix     bool
        rawText      string
        splitText    []string
    )
    
    splitText = strings.Split(p.Text, " ")
    if len(splitText) == 0 {
        return ""
    }
    
    isPrefix = Include(string(splitText[0][0]), p.Prefix)
    if isPrefix == false {
        return ""
    }
    
    rawText = strings.ToLower(splitText[0])
    rawText = strings.Replace(rawText, p.Username, "", 1)
    return rawText[1:]
}


func (p *Parser) Argument() string {
    var (
        rawText      string
        splitText    []string
    )
    splitText = strings.Split(p.Text, " ")
    if len(splitText) == 0 {
        return ""
    }
    
    rawText = strings.Replace(p.Text, splitText[0], "", 1)
    return strings.TrimSpace(rawText)
}