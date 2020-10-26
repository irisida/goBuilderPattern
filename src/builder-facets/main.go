package main

import "fmt"

/*
	person struct that we will build from multiple builder
	structs below. It is separated by the different facets
	of the person being built. The address and job facets
	are treated as separate structs and separate builds
	that join fluidly.
*/
type Person struct {
	// address
	StreetAddress, Postcode, City string

	// job
	CompanyName, Position string
	Salary                int
}

type PersonBuilder struct {
	person *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

// facet builder
type PersonAddressBuilder struct {
	PersonBuilder
}

// address related entry point
func (pb *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*pb}
}

// address related field completion function
func (pab *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	pab.person.StreetAddress = streetAddress
	return pab
}

// address related field completion function
func (pab *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	pab.person.City = city
	return pab
}

// address related field completion function
func (pab *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	pab.person.Postcode = postcode
	return pab
}

// facet builder
type PersonJobBuilder struct {
	PersonBuilder
}

// Job related entry point
func (pb *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*pb}
}

// job related field completion
func (pjb *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	pjb.person.CompanyName = companyName
	return pjb
}

// job related field completion
func (pjb *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	pjb.person.Position = position
	return pjb
}

// job related field completion
func (pjb *PersonJobBuilder) Earns(salary int) *PersonJobBuilder {
	pjb.person.Salary = salary
	return pjb
}

// main build method
func (b *PersonBuilder) Build() *Person {
	return b.person
}

func main() {
	// declare a new personBuilder
	pb := NewPersonBuilder()

	// construct the person
	pb.
		Lives().
		At("23 Arcacia Road").
		In("London").
		WithPostcode("LDN100").
		Works().
		AsA("Programmer").
		At("Facebook").
		Earns(100000)

	person := pb.Build()
	fmt.Printf("value is: %v \n", person)
	fmt.Printf("Type is: %T \n", person)

}
