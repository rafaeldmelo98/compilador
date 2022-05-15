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

const SYMBOLS = `%&.+-*/=`

func LexicalAnalysis(file string, table *Table) {
	checkAlphabet(file)
	labelTokens(file, table)
	printTable(table)
}

func checkAlphabet(file string) {
	file = strings.TrimSuffix(file, "\n")
	fmt.Println(file)
	for _, char := range file {

		exist := strings.Contains(KNOWEDALPHABET, string(char))
		if !exist {
			log.Fatal("Lexical Error: Invalid Character.")
		}
	}
}

func labelTokens(file string, table *Table) {
	var word string
	for _, char := range file {
		if isSymbol(string(char)) && word == "" {
			table.Index = append(table.Index, len(table.Index)+1)
			table.Token = append(table.Token, string(char))
			table.Type = append(table.Type, "symbol")
		}
		if isSymbol(string(char)) && word != "" {

		}
		if isSemicolon(string(char)) {
			table.Index = append(table.Index, len(table.Index)+1)
			table.Token = append(table.Token, string(char))
			table.Type = append(table.Type, "symbol-semicolon")
		}
		if string(char) == ";" && word != "" {
			defineVariable(word, table)
			continue
		}
		word = word + string(char)
		fmt.Println(word)
		if isReservedWord(word, table) {
			table.Index = append(table.Index, len(table.Index)+1)
			table.Token = append(table.Token, word)
			table.Type = append(table.Type, "reserved word")
			continue
		}
		if isEmptySpace(word) {
			word = ""
		}
	}
}

func isSymbol(token string) bool {
	isSymbol := strings.Contains(SYMBOLS, token)
	if isSymbol {
		return true
	}
	return false
}

func isSemicolon(token string) bool {
	if token == ";" {
		return true
	}
	return false
}

func isEmptySpace(token string) bool {
	return token == " "
}

func isReservedWord(token string, table *Table) bool {
	for _, word := range RESERVEDWORDS {
		if token == word {
			return true
		}
	}
	return false
}

func printTable(table *Table) {
	fmt.Println("TABLE\nIndex - Token - Type")
	for i := 0; i < len(table.Index); i++ {
		fmt.Printf("%d - %s - %s\n", table.Index[i], table.Token[i], table.Type[i])
	}
}

func setTokenInTable(table *Table, token string, typeToken string) {
	table.Index = append(table.Index, len(table.Index)+1)
	table.Token = append(table.Token, token)
	table.Type = append(table.Type, typeToken)
}
