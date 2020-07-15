package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	contact   contactInfo
}
type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	//alex := person{"Alex", "Anderson"}

	alex := person{
		firstname: "Alex",
		lastname:  "Anderson",
		contact: contactInfo{
			email:   "test@test",
			zipCode: 999,
		},
	}
	alex.print()
	alexPointer := &alex
	alexPointer.updateName("jim")
	alex.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)

}
func (p *person) updateName(firstname string) {
	(*p).firstname = firstname
}
