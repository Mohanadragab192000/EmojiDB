package core

type FieldType int

const (
	FieldTypeInt FieldType = iota
	FieldTypeString
	FieldTypeFloat
	FieldTypeBool
)

type Field struct {
	Name   string
	Type   FieldType
	Unique bool
}

type Schema struct {
	Version int
	Fields  []Field
}

type ConflictReport struct {
	Compatiable bool
	Conflicts   []string
	Destructive bool
}
