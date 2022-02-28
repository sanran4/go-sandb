package main

import (
	"encoding/json"
	"fmt"

	"github.com/sanran4/go-sandb/db"
)

type Processor struct {
	Name      string
	Core      json.Number
	Frequency string
}

type Memory struct {
	Name     string
	SizeinGB json.Number
}

type Computer struct {
	Name      string
	Processor Processor
	Memory    Memory
	Price     json.Number
}

func main() {
	dbdir := "./"

	//
	db, err := db.New(dbdir, nil)
	if err != nil {
		fmt.Printf("error creating database at '%s'\n", dbdir)
	}

	comps := []Computer{
		{"Inspiron", Processor{"Intel I5", "4", "1.5GHz"}, Memory{"DDR4", "8"}, "65000"},
		{"Vastro", Processor{"Intel I3", "2", "2.0GHz"}, Memory{"DDR4", "4"}, "40000"},
		{"XPS", Processor{"Intel I7", "8", "2.5GHz"}, Memory{"DDR4", "8"}, "55000"},
	}

	// Collection to be created in database
	collection := "computers"

	// Creating collection with data in the database
	for _, v := range comps {
		db.Write(collection, v.Name, Computer{
			Name:      v.Name,
			Processor: v.Processor,
			Memory:    v.Memory,
			Price:     v.Price,
		})
	}

	//Reading all data from collection
	data, err := db.ReadAll(collection)
	if err != nil {
		fmt.Printf("error reading data from '%s'\n", collection)
	}
	fmt.Println(data)

	//// Deleate spesific record from collection
	// if err := db.Delete(collection, "Vastro"); err != nil {
	// 	fmt.Println("Error:", err)
	// }

	//// Delete the whole collection
	// if err := db.Delete(collection, ""); err != nil {
	// 	fmt.Println("Error:", err)
	// }
}
