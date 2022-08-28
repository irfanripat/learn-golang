package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToMalaysia(address *Address) {
	address.Country = "Malaysia"
}

func main() {
	address1 := Address{"Makassar", "Sulawesi Selatan", "Indonesia"}
	address2 := &address1
	address3 := &address1

	address2.City = "Luwu"

	*address2 = Address{"Palopo", "Sulawesi Selatan", "Indonesia"}


	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	address4 := new(Address)
	address4.City = "Pare-Pare"
	fmt.Println(address4)

	address5 := Address{"Pare-Pare", "Sulawesi Selatan", "Indonesia"}

	ChangeCountryToMalaysia(&address5)
	fmt.Println(address5)
}