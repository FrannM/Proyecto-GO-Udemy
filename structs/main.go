package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	frann := person{
		firstName: "Frann",
		lastName:  "Maurizzio",
		contactInfo: contactInfo{
			email:   "fmaurizzio@gmail.com",
			zipCode: 1980,
		},
	}
	//frannPointer := &frann
	frann.updateName("Franco")
	frann.print()

	// DECLARE person with ZERO VALUE
	//var person person

	//UPDATE FIELDS OF THE STRUCT
	//person.firstName = "Frann"
	//person.lastName = "Maurizzio"
	//person.contact.email = "fmaurizzio@gmail.com"
	//person.contact.zipCode = 1980
	//fmt.Println(frann)
	//fmt.Printf("%+v", frann)
}

func (pointer *person) updateName(newFirstName string) {
	(*pointer).firstName = newFirstName
}

func (pointer person) print() {
	fmt.Printf("%+v", pointer)
}
