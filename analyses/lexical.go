package analyses

import (
	"fmt"
	"log"
	"strings"
)

type Table struct {
	Index []int
	Token []string
	Type  []string
}

const KNOWEDALPHABET = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789%&+-*/;= `

var RESERVEDWORDS = [...]string{"asm", "auto", "break", "case", "catch", "char", "class", "const",
	"continue", "default", "delete", "do", "double", "else", "enum", "extern", "float", "for",
	"friend", "goto", "if", "inline", "int", "long", "new", "operator", "private", "protected",
	"public", "register", "return", "short", "signed", "sizeof", "static", "struct", "switch",
	"template", "this", "throw", "try", "typedef", "union", "unsigned", "virtual", "void",
	"volatile", "while"}

const SYMBOLS = `%&.-*/;=`

func LexicalAnalysis(file string) {
	checkAlphabet(file)
	checkReservedWord(file)
}

func checkAlphabet(file string) {
	for _, char := range file {
		exist := strings.Contains(KNOWEDALPHABET, string(char))
		if !exist {
			log.Fatal("Lexical Error: Invalid Character.")
		}
	}
}

func checkReservedWord(file string) {
	var word string
	for _, char := range file {
		word = word + string(char)
		fmt.Println(word)
		if word == "int" {
			log.Fatal("It's integer")
		}
	}
}

func isSymbol() {

}
