package main


import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Hostname")
	user := flag.String("user", "root", "Username")
	password := flag.String("password", "root", "Password")
	number := flag.Int("number", 0, "Number")

	flag.Parse()
	
	fmt.Println("Host:", *host)
	fmt.Println("User:", *user)
	fmt.Println("Password:", *password)
	fmt.Println("Number:", *number)
}