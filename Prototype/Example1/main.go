package main

type Address struct {
	StreetAddress, City, Country, string
}
type Person struct {
	Name string
	Address *Address
}


func main() {
	john := Person{"John", 
	&Address{"123 London Rd", "London", "UK"}
}
	// jane := john
	// jane.Name = "Jane"
	// jane.Address.StreetAddress = "321 Baker St"

	jane := john
	jane.Name = "Jane"
	jane.Address = &Address{
		john.Address.StreetAddress,
		john.Address.City,
		john.Address.Country,
	}

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
