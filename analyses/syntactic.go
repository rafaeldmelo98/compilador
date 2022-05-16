package analyses

import "fmt"

func SyntacticAnalysis(table *Table) error {
	fmt.Println("This is the table", table)
	return nil
}

func checkSyntax() {}
