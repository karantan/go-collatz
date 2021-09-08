package main

import (
	"collatz/logger"
	"flag"
	"time"
)

var log = logger.New("main")

func main() {
	defer timeTrack(time.Now(), "main")

	var n int64
	flag.Int64Var(&n, "n", 1, "Generate Collatz conjecure numbers for 'n' numbers.")
	flag.Parse()

	var i int64 = 1
	for i <= n {
		collatz(i)
		i++
	}
}

func collatz(startingNumber int64) {
	log.Infow("Running Collatz conjecure", "number", startingNumber)

	noLoops := 0
	highestNumber := startingNumber
	number := startingNumber
	for {
		log.Info(number)
		number = collatzMath(number)
		if number == 1 {
			break
		}
		if number > highestNumber {
			highestNumber = number
		}
		noLoops++
	}
	log.Infof(
		"For number %d it took %d loops and the highest number was %d",
		startingNumber,
		noLoops,
		highestNumber,
	)
}

func collatzMath(n int64) int64 {
	if n%2 == 0 {
		return n / 2
	} else {
		return n*3 + 1
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Infof("%s took %s", name, elapsed)
}
