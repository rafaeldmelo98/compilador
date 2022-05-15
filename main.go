package main

import (
	"compilador-trabalho1/analyses"
	"compilador-trabalho1/utils"
	"fmt"
)

func main() {
	file := utils.ReadFile("file.txt")
	analyses.LexicalAnalysis(file)
	analyses.SyntacticAnalysis(file)
	analyses.SemanticAnalysis(file)
	fmt.Println("PASSED!")
}
