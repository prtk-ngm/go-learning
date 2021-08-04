package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	alex.contact.email = "alex@gmail.com"
	alex.contact.zipCode = "560100"

	alex.updateFirstName("Jimmy")
	alex.print()

	//create map
	colors := make(map[string]string)

	colors["red"] = "#ff0000"
	colors["green"] = "#ff1111"
	colors["white"] = "#ffffff"

	printMap(colors)
	delete(colors, "red")
	printMap(colors)

}

func (pointerToPerson *person) updateFirstName(newName string) {

	pointerToPerson.firstName = newName

}

func (p person) print() {

	fmt.Printf("%+v", p)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}

}
