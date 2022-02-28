package main

import (
	"encoding/json"
	"fmt"
	"github.com/sanran4/go-sandb/db"
)

type Address struct {
	City    string
	State   string
	Country string
	Pincode json.Number
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}

func main() {
	dir := "./"

	db, err := db.New(dir, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}

	employees := []User{
		{"Sanjeev", "30", "9111111111", "Dell", Address{"Bangalore", "Karnataka", "India", "560001"}},
		{"Ranjan", "25", "9222222222", "Deloite", Address{"Bangalore", "Karnataka", "India", "560002"}},
		{"Aman", "22", "9333333333", "RVCE", Address{"Bangalore", "Karnataka", "India", "560003"}},
		{"Kshitiz", "6", "9444444444", "DPS", Address{"Gaya", "Bihar", "India", "823001"}},
		{"Hritvik", "4", "9555555555", "Kidzee", Address{"Gaya", "Bihar", "India", "823001"}},
		{"Pooja", "26", "9666666666", "Homemaker", Address{"Gaya", "Bihar", "India", "823001"}},
	}

	for _, value := range employees {
		db.Write("users", value.Name, User{
			Name:    value.Name,
			Age:     value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,
		})
	}

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(records)

	allusers := []User{}

	for _, f := range records {
		empFound := User{}
		if err := json.Unmarshal([]byte(f), &empFound); err != nil {
			fmt.Println("Error:", err)
		}
		allusers = append(allusers, empFound)
	}
	fmt.Println(allusers)

	// if err := db.Delete("users", "Ranjan"); err != nil {
	// 	fmt.Println("Error:", err)
	// }

	// if err := db.Delete("users", ""); err != nil {
	// 	fmt.Println("Error:", err)
	// }
}
