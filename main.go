package main

import (
	"bufio"
	"fmt"
	"github.com/icrowley/fake"
	"github.com/timrourke/fakeql/model"
	"log"
	"os"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	log.Println("Please wait. Writing to file \"tmp.sql\".")

	file, err := os.Create("tmp.sql")
	checkErr(err)

	defer file.Close()

	writer := bufio.NewWriter(file)

	columns := map[string]func() string{
		"first_name": func() string { return fake.FirstName() },
		"last_name":  func() string { return fake.LastName() },
	}

	modelFactory := model.NewModelFactory(columns)

	for i := 0; i < 1500000; i++ {
		_, err = writer.WriteString(
			fmt.Sprintf("INSERT INTO table (%s) VALUES (%s);\n",
				modelFactory.GetColumnsString(), modelFactory.GetRandomValuesString()))
		checkErr(err)
	}

	writer.Flush()
}
