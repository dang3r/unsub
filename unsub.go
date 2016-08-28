package main

import (
	"flag"
	"fmt"
	"github.com/dang3r/unsub/dominos"
	"os"
)

type unsubscribe func(email string) error

var serviceMap = map[string]unsubscribe{
	"dominos": dominos.Unsub,
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage : ./unsub <service> <email>")
		os.Exit(1)
	}

	service := flag.String("s", "dominos", "Name of the requested service")
	email := flag.String("e", "", "Email to be unsubscribed")
	flag.Parse()

	if handler, exists := serviceMap[*service]; exists {
		if err := handler(*email); err != nil {
			fmt.Printf("Error executing handler for %s because %v", *service, err)
			os.Exit(1)
		}
		fmt.Printf("Successfully removed %v from %v\n", *email, *service)
	} else {
		fmt.Printf("Service %s not supported", *service)
	}
}
