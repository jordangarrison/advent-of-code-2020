package main

import (
	"fmt"
	"os"

	"github.com/jordangarrison/advent-of-code/2021/go/day1"
	"github.com/jordangarrison/advent-of-code/2021/go/day2"
	"github.com/jordangarrison/advent-of-code/2021/go/day3"
	"github.com/jordangarrison/advent-of-code/2021/go/day4"
	"github.com/jordangarrison/advent-of-code/2021/go/day5"
	"github.com/jordangarrison/advent-of-code/2021/go/util"
)

func main() {
	// Get arguments from command line
	args := os.Args[1:]

	for _, arg := range args {
		// Run day corresponding to argument
		switch arg {
		case "1":
			part1, err := util.GetData(1, 1)
			if err != nil {
				panic(err)
			}
			util.Stats(day1.Run, part1)
		case "2":
			part1, err := util.GetData(2, 1)
			if err != nil {
				panic(err)
			}
			util.Stats(day2.Run, part1)
		case "3":
			part1, err := util.GetData(3, 1)
			if err != nil {
				panic(err)
			}
			util.Stats(day3.Run, part1)
		case "4":
			part1, err := util.GetData(4, 1)
			if err != nil {
				panic(err)
			}
			util.Stats(day4.Run, part1)
		case "5":
			part1, err := util.GetData(5, 1)
			if err != nil {
				panic(err)
			}
			util.Stats(day5.Run, part1)
		default:
			fmt.Printf("Day %s not implemented\n", arg)
		}
	}
}
