package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// Assign kv
	// emails["Bob"] = "bob@mail.com"
	// emails["Sharon"] = "sharon@mail.com"
	// emails["Mike"] = "mike@mail.com"

	// Declare map and add kv
	emails := map[string]string{"Bob": "bob@mail.com", "Sharon": "sharon@mail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)
}
