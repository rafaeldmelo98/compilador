package main

import (
	"compilador-trabalho1/analyses"
	"compilador-trabalho1/utils"
	"fmt"
)

func main() {
	file := utils.ReadFile("file.txt")
	err := analyses.LexicalAnalysis(file)
	utils.CheckIfError(err)
	err = analyses.SyntacticAnalysis(file)
	utils.CheckIfError(err)
	err = analyses.SemanticAnalysis(file)
	utils.CheckIfError(err)
	fmt.Println("PASSED!")
}
