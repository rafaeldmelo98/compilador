package main

import (
	"compilador-trabalho1/analyses"
	"compilador-trabalho1/utils"
	"fmt"
)

func main() {
	var table analyses.Table
	file := utils.ReadFile("file.txt")
	analyses.LexicalAnalysis(file, &table)
	analyses.SyntacticAnalysis(&table)
	analyses.SemanticAnalysis(file)
	fmt.Println("PASSED!")
}
