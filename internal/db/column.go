package db

type ValueType int

const (
	ValueUnknown ValueType = iota
	ValueInt
	ValueText
)

type ColumnDef struct {
	Name string
	Type ValueType
}


