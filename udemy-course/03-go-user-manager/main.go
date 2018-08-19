package main

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 78758,
		},
	}

	jim.updateName("James")
	jim.print()
}
