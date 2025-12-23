package db

type StatementType int

const (
	StatementUnknown StatementType = iota
	StatementSelect
	StatementCreate
)

type Statement struct {
	Type StatementType
}

func ParseStatemnt(input string) Statement {
	if input == "SELECT" {
		return  Statement {
 			Type: StatementSelect,
		}
	}
	
	if input == "CREATE" {
		return Statement {
			Type: StatementCreate,
		}
	}

	return Statement {
		Type: StatementUnknown,
	} 
 }
