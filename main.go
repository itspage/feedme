package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/itspage/feedme/outlets"
	"github.com/itspage/feedme/user"
)

func main() {
	loggingOnOff()

	// User Preferences
	prefs, err := user.LoadPreferences()
	if err != nil {
		log.Fatal(err)
	}

	outlet, err := outlets.Select(prefs)

	// Login
	profile, err := outlet.Login(prefs)
	if err != nil {
		log.Fatal(err)
	}

	// Build order
	order, err := outlet.NewOrder(prefs, profile)
	if err != nil {
		log.Fatal(err)
	}

	// Confirm
	fmt.Printf("Hello %v! I guess you're hungry?", profile.Name)
	fmt.Println()
	fmt.Println()
	//fmt.Println("--------------------------")
	fmt.Println("Your order is:")
	fmt.Println(order)
	fmt.Println("--------------------------")
	fmt.Println()
	fmt.Println("Are you ready to order? [Y/n] ")
	if !confirm() {
		fmt.Println("Abort.")
		os.Exit(1)
	}

	// Go ahead and place the order
	err = order.Submit()
	if err != nil {
		log.Fatal(err)
	}
}

// loggingOnOff enables or disables standard logger according to environment
func loggingOnOff() {
	if os.Getenv("FEEDME_LOGGING") == "" {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}
}

func confirm() bool {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if input := scanner.Text(); input != "Y" && input != "" {
		return false
	}
	return true
}
