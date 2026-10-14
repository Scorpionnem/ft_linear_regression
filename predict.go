package main

import ("fmt"
		"strconv"
		"log"
		"os")

func estimatePrice(km float64, theta0 float64, theta1 float64) float64 {
	return (theta0 + (theta1 * km));
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./predict <theta0> <theta1>")
		return
	}

	theta0, err := strconv.ParseFloat(os.Args[1], 64)
	if (err != nil) {
		log.Fatal("Error parsing value: " + os.Args[1], err)
	}
	theta1, err := strconv.ParseFloat(os.Args[2], 64)
	if (err != nil) {
		log.Fatal("Error parsing value: " + os.Args[1], err)
	}

	var km float64
	for true {
		fmt.Print("Enter km: ")
	  	fmt.Scan(&km)
		fmt.Println("Estimated price:", estimatePrice(km, theta0, theta1))
	}
}
