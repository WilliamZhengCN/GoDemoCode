package main

import "fmt"

func main() {
	ShowMainMenu()
}

func ShowMainMenu() {
	var cmd int
	fmt.Println("*********************************")
	fmt.Println("Welcome to Simple Chat:")
	fmt.Println("\t1 Login In")
	fmt.Println("\t2 Register")
	fmt.Println("\t3 Exist")
	fmt.Println("Please type 1--3 for operation:")
	fmt.Println("*********************************")

	for {
		fmt.Scanf("%d\n", &cmd)
		switch cmd {
		case 1:
			HandleLoginIn()
		case 2:
		case 3:
		default:
			fmt.Println("Input Error.")
		}
	}
}

func HandleLoginIn() {
	var username string
	var password string
	fmt.Println("User Name:")
	fmt.Scanln(&username)
	fmt.Println("Password:")
	fmt.Scanln(&password)
	userInfo := ChatUser{
		UserName: username,
		Password: password,
	}
	err := LoginIn(userInfo)
	if err != nil {
		fmt.Println("Fail to login int")
	}
}
