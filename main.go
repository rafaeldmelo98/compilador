package main

import (
	"compilador-trabalho1/analyses"
	"compilador-trabalho1/utils"
)

func main() {
	file := utils.ReadFile("file.txt")
	err := analyses.LexicalAnalysis(file)
	utils.CheckIfError(err)
	err = analyses.SyntacticAnalysis(file)
	utils.CheckIfError(err)

}
