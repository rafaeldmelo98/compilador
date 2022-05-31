package analyses

func SyntacticAnalysis(table *Table) error {
	tokens := table.Token
	// Analisar tabela

	for i := 0; i < len(table.Index); i++ {
		if tokens[i] == "reserved word" {

		}
	}
	return nil
}

func checkSyntaxIf(tokens []string) bool {
	return true
}
