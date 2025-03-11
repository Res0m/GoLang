package main

import (
	"flag"
	"fmt"
)

func main() {
	// name := flag.String("name", "Anton", "Name of user") // (название флага, дефолтное значение, описание)
	// age := flag.Int("age", 18, "Age of user")
	// isAdmin := flag.Bool("isAdmin", false, "Is user admin")
	city := flag.String("city", "", "City of user")
	format := flag.Int("format", 1, "Формат вывода погоды")


	flag.Parse()


	fmt.Println(*city)
	fmt.Println(*format)
}	
