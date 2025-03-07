package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// Create a new channel with make(chan val-type). Channels are typed by the values they convey.
	beersCmd := flag.NewFlagSet("beers", flag.ExitOnError)
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatal("expected a subcommand")
		os.Exit(2)
	}

	switch flag.Arg(0) {
	case "beers":
		ID := beersCmd.String("id", "", "id of the beer")
		beersCmd.Parse((os.Args[2:]))

		if *ID != "" {
			fmt.Println("ID:", *ID)
		}
		fmt.Println("beers subcommand")

	}

}
