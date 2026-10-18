package main

import (
	"fmt"
	"ft_linear_regression/internal/utils"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./predict <theta0> <theta1>")
		return
	}

	theta0, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		log.Fatal("Error parsing value: "+os.Args[1], err)
	}
	theta1, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		log.Fatal("Error parsing value: "+os.Args[2], err)
	}

	var km float64
	for {
		fmt.Print("Enter km: ")
		fmt.Scan(&km)
		fmt.Println("Estimated price:", utils.EstimatePrice(km, theta0, theta1))
	}
}
