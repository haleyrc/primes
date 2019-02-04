package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/frazercomputing/primes/rm"
)

func main() {
	val := os.Args[1]
	intVal, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		log.Fatalf("invalid integer value: %s: %v\n", val, err)
	}
	if rm.Prime(intVal, rm.DefaultTrials) {
		fmt.Println("prime")
	} else {
		fmt.Println("not prime")
	}
}
