package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/icrowley/fake"
	"log"
	"os"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	log.Println("hello")

	file, err := os.Create("tmp.sql")
	checkErr(err)

	defer file.Close()

	writer := bufio.NewWriter(file)

	modelFactory := map[string]func() string{
		"first_name": func() string { return fake.FirstName() },
		"last_name":  func() string { return fake.LastName() },
	}

	modelFactoryLength := len(modelFactory)

	var columnsBuffer bytes.Buffer
	i := 0
	for columnName := range modelFactory {
		columnsBuffer.WriteString("`")
		columnsBuffer.WriteString(columnName)
		columnsBuffer.WriteString("`")
		i++

		if i < modelFactoryLength {
			columnsBuffer.WriteString(", ")
		}
	}

	for i := 0; i < 1500000; i++ {
		var valuesBuffer bytes.Buffer
		j := 0
		for _, fakerFunc := range modelFactory {
			valuesBuffer.WriteString("\"")
			valuesBuffer.WriteString(fakerFunc())
			valuesBuffer.WriteString("\"")
			j++

			if j < modelFactoryLength {
				valuesBuffer.WriteString(", ")
			}
		}

		_, err = writer.WriteString(
			fmt.Sprintf("INSERT INTO table (%s) VALUES (%s);\n",
				columnsBuffer.String(), valuesBuffer.String()))
		checkErr(err)
	}

	writer.Flush()
}
