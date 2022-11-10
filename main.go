package main

import (
	"fmt"
	"strings"
)

func main()  {
	var conferenceName = "Go Conference"	
	const conferenceTickets uint = 50
	var name string
	exampleList := []string{} // := string[]{}
	exampleList = append(exampleList, "Cris Camilo")
	exampleList = append(exampleList, "Cris 2")
	exampleList = append(exampleList, "Cris 3")
	// Get name from user by console
	fmt.Scan(&name)
	
	fmt.Println("Wellcome", conferenceName, "Booking application")
	fmt.Printf("It %v Works with Us. \n", name)

	for i, item := range exampleList {
		var names []string = strings.Fields(item)
		var firstName string = names[0]
		fmt.Println("%v name: %v \n", i, firstName)
	}
}