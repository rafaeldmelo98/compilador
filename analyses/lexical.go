package analyses

import (
	"compilador-trabalho1/utils"
	"fmt"
	"strings"
)

const KNOWEDALPHABET = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789%&+-*/;= `

var RESERVEDWORDS = [...]string{"asm", "auto", "break", "case", "catch", "char", "class", "const",
	"continue", "default", "delete", "do", "double", "else", "enum", "extern", "float", "for",
	"friend", "goto", "if", "inline", "int", "long", "new", "operator", "private", "protected",
	"public", "register", "return", "short", "signed", "sizeof", "static", "struct", "switch",
	"template", "this", "throw", "try", "typedef", "union", "unsigned", "virtual", "void",
	"volatile", "while"}

func LexicalAnalysis(file string) error {
	checkAlphabet(file)
	return nil
}

func checkAlphabet(file string) {
	for _, char := range file {
		exist := strings.Contains(KNOWEDALPHABET, string(char))
		if !exist {
			utils.CheckIfError(fmt.Errorf("Lexical Error: Invalid Character."))
		}
	}
}
