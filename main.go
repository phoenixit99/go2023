package main

import (
	"fmt"
	"os"
	"strings"
)

var listCountryCode = []string{
	"US",
	"VN",
}

func isEmptyString(s string) bool {
	return len(s) == 0
}

func isValidCountryCode(arr []string, target string) bool {
	for _, s := range arr {
		if s == target {
			return true
		}
	}
	return false
}

func showFullNameByFirstRule(firstName, lastName, middleName string) {
	if isEmptyString(middleName) {
		fmt.Println("Output = ", lastName, firstName)
	} else {
		fmt.Println("Output = ", lastName, middleName, firstName)
	}
}

func showFullNameBySecRule(firstName, lastName, middleName string) {
	if isEmptyString(middleName) {
		fmt.Println("Output = ", firstName, lastName)
	} else {
		fmt.Println("Output = ", firstName, middleName, lastName)
	}
}

func reOrderFullName(firstName, lastName, middleName, countryCode string) {

	if isValidCountryCode(listCountryCode, strings.ToUpper(countryCode)) {
		switch strings.ToUpper(countryCode) {
		case "US": // "First Name Middle Name Last Name"
			showFullNameBySecRule(firstName, lastName, middleName)
		case "VN": // "Last Name First Name Middle Name"
			showFullNameByFirstRule(firstName, lastName, middleName)
		default:
			showFullNameByFirstRule(firstName, lastName, middleName)
		}

	} else {
		fmt.Println("The country code is not support or invalid")
	}
}

func parseContextName() (string, string, string, string) {
	firstName := ""
	lastName := ""
	middleName := ""
	countryCode := ""
	for i, arg := range os.Args[1:] {
		if i == 0 {
			firstName = arg
		} else if i == 1 {
			lastName = arg
		} else if i == len(os.Args[1:])-1 {
			countryCode = arg
		} else {
			middleName += arg
			middleName += " "
		}
	}
	return firstName, lastName, strings.TrimSpace(middleName), countryCode
}

func main() {

	firstName, lastName, middleName, countryCode := parseContextName()
	reOrderFullName(firstName, lastName, middleName, countryCode)

}
