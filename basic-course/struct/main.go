package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode string
}

type person struct {
	firstName   string
	lastName    string
	contactInfo contactInfo
}

func (p person) toString() string {
	return p.firstName + " " + p.lastName
}

func (p person) updateName_wrong(name string) {
	// p is just pointing at a copy of the actual person
	p.lastName = name
	// so this is not for an update for person, just do some logic and return result
}

func (p *person) updateName(name string) {
	(*p).firstName = name
}

func main() {
	// 2 ways to define and assign value
	alex := person{
		"Alex",
		"Anderson",
		contactInfo{
			"alexa@gmail.com",
			"100000",
		},
	}

	alex = person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			"alexa@gmail.com",
			"100000",
		},
	}

	barkPointer := &alex
	fmt.Println("\n----- Alex -----\n")
	fmt.Println(*&alex)
	barkPointer.updateName("Bark")
	fmt.Printf("%+v", *barkPointer)
	alex.updateName("Go_implicitly_converted_alex_to_a_pointer_to_alex")
	fmt.Println(alex)
	alex.updateName_wrong("Bark")
	fmt.Println(alex)
	fmt.Println("\n----- Alex -----\n")

	fmt.Println("\n----- Update Slice -----\n")
	slice := make([]string, 10)
	print(slice)
	mySlice := []string{"Hi", "There", "How", "Are", "You?"}

	fmt.Println("When declaring and assigning values to a slice, " +
		"it implicitly create an array and a structure " +
		"containing length, capacity, " +
		"and references to the underlying actual values in the array")

	updateSlice(mySlice)
	fmt.Println(mySlice)
	fmt.Println("\n----- Update Slice -----\n")
}

func updateSlice(slice []string) {
	slice[0] = "Hello"
}
