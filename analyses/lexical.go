package analyses

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Table struct {
	Index []int
	Token []string
	Type  []string
}

const KNOWEDALPHABET = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789%&+-*/;:.,=(){}><!"' `

var RESERVEDWORDS = [...]string{"asm", "auto", "break", "case", "catch", "char", "class", "const",
	"continue", "default", "delete", "do", "double", "else", "enum", "extern", "float", "for",
	"friend", "goto", "if", "inline", "int", "long", "new", "operator", "private", "protected",
	"public", "register", "return", "short", "signed", "sizeof", "static", "struct", "switch",
	"template", "this", "throw", "try", "typedef", "union", "unsigned", "virtual", "void",
	"volatile", "while", "print", "scan"}

const SYMBOLS = `%&+-*/:=(){}><!`

func LexicalAnalysis(file string, table *Table) {
	file = cleanFile(file)
	checkAlphabet(file)
	labelTokens(file, table)
	printTable(table)
}

func cleanFile(file string) string {
	file = strings.ReplaceAll(file, "\n", "")
	file = strings.ReplaceAll(file, "\t", "")
	file = strings.ReplaceAll(file, " ", "")
	return file
}

func checkAlphabet(file string) {
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
		if isString(word) {
			setTokenInTable(table, word, "string")
			word = ""
			continue
		}
		if !isNumeric(string(char)) && isNumeric(word) && word != "" && string(char) != "." {
			setTokenInTable(table, word, "numeric")
			checkNextCharacterIfSymbolOrSemicolon(string(char), table)
			word = ""
			continue
		}
		if isSymbol(string(char)) && word == "" {
			setTokenInTable(table, string(char), "symbol")
			word = ""
			continue
		}
		if isSymbol(string(char)) && word != "" {
			setTokenInTable(table, word, "variable")
			setTokenInTable(table, string(char), "symbol")
			word = ""
			continue
		}
		if isSemicolon(string(char)) && word == "" {
			setTokenInTable(table, string(char), "symbol-semicolon")
			word = ""
			continue
		}
		word = word + string(char)
		if isReservedWord(word, table) {
			setTokenInTable(table, word, "reserved word")
			word = ""
			continue
		}
	}
}

func checkNextCharacterIfSymbolOrSemicolon(token string, table *Table) {
	if isSymbol(string(token)) {
		setTokenInTable(table, string(token), "symbol")
	} else if isSemicolon(string(token)) {
		setTokenInTable(table, string(token), "symbol-semicolon")
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

func isReservedWord(token string, table *Table) bool {
	for _, word := range RESERVEDWORDS {
		if token == word {
			return true
		}
	}
	return false
}

func isNumeric(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	if err == nil {
		return true
	}
	_, err = strconv.Atoi(token)
	if err == nil {
		return true
	}
	return false
}

func isString(token string) bool {
	if len(token) < 2 {
		return false
	}
	return string(token[0]) == "\"" && string(token[len(token)-1]) == "\""
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
