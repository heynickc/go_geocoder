package main

import (
	"fmt"
	"log"
)

func main() {

	addresses, err := readDataFile("my_address.json")

	if err != nil {
		log.Fatalln("Failed to read:", err)
	}

	fmt.Println(addresses)
}
