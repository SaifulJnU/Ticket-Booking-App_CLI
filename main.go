package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var conferenceName = "Go Conference"
var totalTickets uint = 50
var totalRemainingTickets uint = 50

//array
//var bookings []string //slice
//bookings := []string{} //shortcut slice declaration
//shortcut declaration is not allowed in global scope

// map
// var bookings = make([]map[string]string, 0)

var bookings = make([]User, 0) //struct types array

type User struct {
	userFirstName   string
	userLastName    string
	userMailAddress string
	userTickets     uint
}

// for synchronizing goroutine we need to use waitgroup
var wg = sync.WaitGroup{}

func main() {

	//Greetings
	greetUsers()

	//infinite for loop
	//for {
	//calling  inputUser() and destructure all the returned variables
	userFirstName, userLastName, userMailAddress, userTickets := inputUser()

	//calling  userInputValidation() and destructure all the returned variables
	isValidName, isValidMail, isValidUserTickets := userInputValidation(userFirstName, userLastName, userMailAddress, userTickets)

	if isValidName && isValidMail && isValidUserTickets {

		//store name into the array
		//here append is working for slice not array
		bookTickets(userFirstName, userLastName, userMailAddress, userTickets)

		//calling sendTickets function
		//go key for asynchronous go concurrency
		wg.Add(1) //for one thread
		go sendTickets(userTickets, userFirstName, userLastName, userMailAddress)

		//store firstName into firstNames slice
		printFirstNames()

		//corner case
		if totalRemainingTickets == 0 {
			fmt.Println("Our", conferenceName, "all tickets is booked out, come back next year!!!")
			//break
		}

	} else {
		if !isValidName {
			fmt.Println("Your entered name is too short!")
		}
		if !isValidMail {
			fmt.Println("Your entered mail is invalid!")
		}
		if !isValidUserTickets {
			fmt.Println("Your entered number of tickets is invalid!")

		}
		//continue
	}

	//switch statement
	// city := "London"

	// switch city {
	// case "UK":

	// case "Berlin", "Munich":

	// case "Amstaderm", "Costria":

	// case "Dhaka", "Lakshmipur":

	// default:
	// 	fmt.Println("Your selected city is not valid!")

	// }

	wg.Wait() //for waiting the newly created thread

}

//{

func greetUsers() {
	fmt.Println("Welcome to", conferenceName, "booking system")
	fmt.Println("Total number of ticket is", totalTickets, "and still abvailable is", totalRemainingTickets)
	fmt.Println("Get your tickets from here to attened")
}

func inputUser() (string, string, string, uint) {
	//Take user input and store
	var userFirstName string
	var userLastName string
	var userMailAddress string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scanln(&userFirstName)

	fmt.Println("Enter your last name:")
	fmt.Scanln(&userLastName)

	fmt.Println("Enter your mail address:")
	fmt.Scanln(&userMailAddress)

	fmt.Println("How many tickets do you want to book?")
	fmt.Scanln(&userTickets)

	return userFirstName, userLastName, userMailAddress, userTickets

}

func userInputValidation(userFirstName string, userLastName string, userMailAddress string, userTickets uint) (bool, bool, bool) {

	isValidName := len(userFirstName) >= 2 && len(userLastName) >= 2
	isValidMail := strings.Contains(userMailAddress, "@")
	isValidUserTickets := userTickets > 0 && userTickets <= totalRemainingTickets

	return isValidName, isValidMail, isValidUserTickets
}

func bookTickets(userFirstName string, userLastName string, userMailAddress string, userTickets uint) {

	// var user = make(map[string]string)
	// user["firstName"] = userFirstName
	// user["lastName"] = userLastName
	// user["email"] = userMailAddress
	// user["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	var user = User{
		userFirstName:   userFirstName,
		userLastName:    userLastName,
		userMailAddress: userMailAddress,
		userTickets:     userTickets,
	}

	bookings = append(bookings, user)

	//bookings = append(bookings, userFirstName+" "+userLastName)

	totalRemainingTickets = totalRemainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v, you will receive a confirmation mail to %v\n", userFirstName, userLastName, userTickets, userMailAddress)

	fmt.Printf("%v tickets are still remaining for %v\n", totalRemainingTickets, conferenceName)

}

func printFirstNames() {
	firstNames := []string{}
	for _, booking := range bookings {

		firstNames = append(firstNames, booking.userFirstName)

	}
	fmt.Printf("Total bookings are: %v\n", firstNames)
}

// delay function
func sendTickets(userTickets uint, userFirstName string, userLastName string, userMailAddress string) {

	time.Sleep(20 * time.Second)

	var tickets = fmt.Sprintf("%v tickets for %v %v \n to mail: %v", userTickets, userFirstName, userLastName, userMailAddress)

	fmt.Println("#################################")
	fmt.Printf("Sending tickets: \n %v\n", tickets)
	fmt.Println("#################################")

	wg.Done() // when thread is over

}
