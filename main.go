package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Hello World!")
	rand.Seed(time.Now().UTC().UnixNano()) // need to init the random seed
	faces, numRoll := parseInput()
	doRoll(faces, numRoll)
}

func parseInput() (faces int, numRoll int) {
	dice := flag.String("d", "d6", "Type of dice under dX format (e.g. d6 for a dice having 6 faces)")
	rollFlag := flag.Int("n", 1, "The number of die to roll (default: 1)")
	flag.Parse()
	fmt.Printf("dice type: %v\n", *dice)
	fmt.Printf("numRoll: %v\n", *rollFlag)

	matched, _ := regexp.Match("^d\\d+$", []byte(*dice))
	fmt.Printf("Matched: %v\n", matched)

	if !matched {
		log.Fatalf("Improper format")
	}

	diceSides := (*dice)[1:]
	faces, err := strconv.Atoi(diceSides)
	fmt.Printf("Number of sites: (%T)%v\n", faces, faces)
	if err != nil {
		log.Fatal(err)
	}
	return faces, *rollFlag
}

func doRoll(faces int, numRoll int) {
	for i := 0; i < numRoll; i++ {
		roll := rand.Intn(faces) + 1
		fmt.Printf("You rolled a %v\n", roll)
	}
}
