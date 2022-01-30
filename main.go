package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Hello World!")
	rand.Seed(time.Now().UTC().UnixNano()) // need to init the random seed
	faces, numRoll := parseInput()
	rolls := doRoll(faces, numRoll)
	printRoll(rolls)
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

func doRoll(faces int, numRoll int) []int {
	values := []int{}
	for i := 0; i < numRoll; i++ {
		roll := rand.Intn(faces) + 1
		values = append(values, roll)
	}
	return values
}

func printRoll(rolls []int) {
	fmt.Printf("\n\nRolls: %v\n", rolls)
	for i, roll := range rolls {
		fmt.Printf("(Roll.%v,%v) ", i+1, roll)
	}
	fmt.Printf("\n\nSum: %v\n", sumDices(rolls))
	fmt.Printf("Highest: %v\n", highestRoll(rolls))
	fmt.Printf("Lowest: %v\n", lowestRoll(rolls))
}

func sumDices(rolls []int) int {
	sum := 0
	for _, roll := range rolls {
		sum += roll
	}
	return sum
}

func highestRoll(rolls []int) int {
	sort.Ints(rolls)
	return rolls[len(rolls)-1]
}

func lowestRoll(rolls []int) int {
	sort.Ints(rolls)
	return rolls[0]
}
