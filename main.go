package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	mark := flag.String("mark", "2019-09-01T00:00:00+05:30", "The remaining time for the countdown timer in RFC3339 format (e.g. 2019-09-01T00:00:00+05:30)")
	flag.Parse()

	if *mark == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	v, err := time.Parse(time.RFC3339, *mark)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for range time.Tick(1 * time.Second) {
		timeRemaining := getTimeRemaining(v)

		if timeRemaining.t <= 0 {
			fmt.Println("Countdown reached!")
			break
		}

		fmt.Printf("%02d Days, %02d Hours, %02d Minutes, %02d Seconds. \n", timeRemaining.d, timeRemaining.h, timeRemaining.m, timeRemaining.s)

	}
}

type countdown struct {
	t int
	d int
	h int
	m int
	s int
}

func getTimeRemaining(t time.Time) countdown {
	currentTime := time.Now()
	difference := t.Sub(currentTime)

	total := int(difference.Seconds())
	days := int(total / (60 * 60 * 24))
	hours := int(total / (60 * 60) % 24)
	minutes := int(total/60) % 60
	seconds := int(total % 60)

	return countdown{
		t: total,
		d: days,
		h: hours,
		m: minutes,
		s: seconds,
	}
}
