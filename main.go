package main

import (
	"fmt"
	"log"
)

func main() {

	addresses, err := readAddressDataFile("my_address.json")

	if err != nil {
		log.Fatalln("Failed to read:", err)
	}

	fmt.Println(addresses)
}
