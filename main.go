package main

import (
	"fmt"
	"log"

	"github.com/dbgventures/polarismail-go/polarismail"
	"github.com/dbgventures/polarismail-go/types"
)

func main() {
	creds := types.Credentials{
		Username: "dbgventures",
		Password: "mezkis-4xurmE-vamhem",
	}

	client, err := polarismail.NewClient(&creds)
	if err != nil {
		log.Fatal(err)
	}

	stats, err := client.Admin().GetStats()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", stats)
}
