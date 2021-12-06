package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	//panic("Please implement the Welcome function")
	//fmt.Sprintf("%s is %d years old.\n", name, age)
	//fmt.Println("Welcome to my party, ", name)
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	//panic("Please implement the HappyBirthday function")
	//fmt.Println("Happy birthday " + name + "! You are now " + age + " years old!")
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	//panic("Please implement the AssignTable function")
	// Welcome to my party, Christiane!
	// You have been assigned to table 027. Your table is on the left, exactly 23.8 meters from here.
	// You will be sitting next to Frank.
	mystr := Welcome(name)
	mystr += fmt.Sprintf("\nYou have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.", table, direction, distance)
	//mystr += fmt.Sprintf(" Your table is %s, exactly %.1f meters from here.", direction, distance)
	mystr += fmt.Sprintf("\nYou will be sitting next to %s.", neighbor)
	return mystr
}
