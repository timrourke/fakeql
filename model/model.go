package model

import (
	"bytes"
)

type ModelFactory struct {
	Columns       map[string]func() string
	ColumnsString string
	NumColumns    int
}

func NewModelFactory(columns map[string]func() string) *ModelFactory {
	m := new(ModelFactory)
	m.Columns = columns
	m.NumColumns = len(columns)
	m.ColumnsString = m.GetColumnsString()
	return m
}

func (modelFactory *ModelFactory) GetColumnsString() string {
	var columnsBuffer bytes.Buffer
	i := 0

	for columnName := range modelFactory.Columns {
		columnsBuffer.WriteString("`")
		columnsBuffer.WriteString(columnName)
		columnsBuffer.WriteString("`")
		i++

		if i < modelFactory.NumColumns {
			columnsBuffer.WriteString(", ")
		}
	}

	return columnsBuffer.String()
}

func (modelFactory *ModelFactory) GetRandomValuesString() string {
	var valuesBuffer bytes.Buffer
	i := 0

	for _, fakerFunc := range modelFactory.Columns {
		valuesBuffer.WriteString("\"")
		valuesBuffer.WriteString(fakerFunc())
		valuesBuffer.WriteString("\"")
		i++

		if i < modelFactory.NumColumns {
			valuesBuffer.WriteString(", ")
		}
	}

	return valuesBuffer.String()
}
