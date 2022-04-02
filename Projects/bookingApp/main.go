package main

import (
	"bookingApp/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName string = "Go conference"

const conferenceTickets uint = 50

var RemainingTickets uint = 50

//var bookings []strings
//var Bookings = make([]map[string]string, 0) // list of maps

// struct good for when the requirement is to store different data type in a single data object

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// list of struct

var Bookings = make([]UserData, 0)

var waitGroup = sync.WaitGroup{}

func main() {

	greetUsers()

	//for {
	//var firstName string
	//var lastName string
	//var email string
	//var userTickets uint
	// ask user their name

	firstName, lastName, email, userTickets := getUserInputs()

	//fmt.Println("")
	//fmt.Print("Enter you first Name: ")
	//fmt.Scan(&firstName)
	//fmt.Print("Enter your last name : ")
	//fmt.Scan(&lastName)
	//fmt.Print("Enter you Email : ")
	//fmt.Scan(&email)
	//fmt.Print("Enter number of tickets: ")
	//fmt.Scan(&userTickets)

	isValidName, isValidTicketNumber, isValidEmail := helper.ValidateUserInputs(firstName, lastName, email, userTickets, RemainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(firstName, lastName, email, userTickets, RemainingTickets)
		// go keyword creates a new thread
		waitGroup.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		//fmt.Printf("Thank you %v %v for booking %v tickets. You will recevie a conmfirmation e-mail at %v \n", firstName, lastName, userTickets, email)
		//remainingTickets = remainingTickets - userTickets
		//
		//bookings = append(bookings, firstName+" "+lastName)

		firstNames := getFirstNames()

		fmt.Println("First names of bookings : ", firstNames)

		//fmt.Printf(" \n the whole array %v \n", bookings)
		//fmt.Printf("\n The first value %v \n The type of the array %T \n", bookings, bookings[0], bookings)
		fmt.Printf("Size of the array %v ", len(Bookings))

		//var noTicketRemaining bool = remainingTickets == 0

		if RemainingTickets == 0 {
			fmt.Println("\n Our Conference is booked out")

		}
	} else {
		//fmt.Println("Your input data is not valid")
		fmt.Printf("%v tickets remaining for %v", RemainingTickets)
		//continue

		if !isValidName {
			fmt.Println("First name or last name you enterd is too short")
		}
		if !isValidTicketNumber {
			fmt.Println("Number of tickets you wated to purchase is either a negative value or it is higher than the available tickets")
		}
		if !isValidEmail {
			fmt.Println("Email address you enterd does not contains a @")
		}
	}

	city := "london"

	switch city {

	case "singapore", "gfu":
		fmt.Println("Your in singapore")
	case "sri lnaka":
		fmt.Println("You are in London")
	default:
		fmt.Println("You are in london")

	}
	//}

	waitGroup.Wait()
}

func sendTicket(userTickets uint, fileName, lastName, email string) {
	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, fileName, lastName)
	time.Sleep(10 * time.Second)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("*********************************************************")
	fmt.Println("")
	fmt.Printf("Sending ticket %v to email address %v \n", ticket, email)
	fmt.Println("")
	fmt.Println("*********************************************************")
	waitGroup.Done()
}

func getUserInputs() (firstName, lastName, email string, userTickets uint) {
	fmt.Println("")
	fmt.Print("Enter First name : ")
	fmt.Scan(&firstName)
	fmt.Print("Enter Last Name : ")
	fmt.Scan(&lastName)
	fmt.Print("Enter E-Mail : ")
	fmt.Scan(&email)
	fmt.Print("Enter Number of Tickets : ")
	fmt.Scan(&userTickets)

	return
}

func getFirstNames() (firstNames []string) {
	for _, booking := range Bookings {
		//name := strings.Fields(bookObj)bookOb
		//firstNames = append(firstNames, Booking["firstName"])
		firstNames = append(firstNames, booking.firstName)
		fmt.Println(firstNames)
	}
	return
}

func greetUsers() {
	fmt.Printf("welcome to %v out booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available to porches \n", conferenceTickets, RemainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func bookTicket(firstName string, lastName string, email string, userTickets, RemainingTickets uint) {
	fmt.Printf("Thank you %v %v for booking %v tickets. You will recevie a conmfirmation e-mail at %v \n", firstName, lastName, userTickets, email)
	RemainingTickets = RemainingTickets - userTickets

	// type of a slice

	//var myslice = []string{}
	//var myMap = map[string]string

	// create a map for a user // map is simler to dictionary

	//var userData = make(map[string]string)
	//userData["Key"] = "Value"

	// creating a userdata struct

	var UserData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	//userData["firstName"] = firstName
	//userData["lastName"] = lastName
	//userData["email"] = email
	//userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	Bookings = append(Bookings, UserData)
	fmt.Printf("%v tickets remaining for %v", RemainingTickets)
	return
}
