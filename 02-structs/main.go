package main

import "fmt"

type contactInfo struct {
	email string
	phone int
}

type person struct {
	firstname string
	lastName  string
	contact   contactInfo
}

func main() {
	// alex := person{
	// 	lastName:  "Anderson",
	// 	firstname: "Alex",
	// }

	var alex person
	alex.firstname = "Alex"
	alex.lastName = "Anderson"
	alex.contact.email = "alex.anderson@email.com"
	alex.contact.phone = 9876543210

	jim := person{
		firstname: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email: "jim.party@email.com",
			phone: 9876543211,
		},
	}

	alex.print()

	jim.updateName("Jimmy")
	jim.print()
}

func (p *person) updateName(newFirstName string) {
	p.firstname = newFirstName
}

func (p person) print() {
	fmt.Println(p.firstname)
	fmt.Printf("%+v", p)
	fmt.Println()
}
