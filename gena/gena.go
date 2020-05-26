package main

import (
	"time"
	"fmt"
	"os"
	"flag"
	"log"
)

func main() {
	endString := flag.String("end", "END", "End string")
	dieDurationString := flag.String("d", "3s", "Die after")
	generateEachString := flag.String("e", "1s", "Generate string each N")
	flag.Parse()

	d, err := time.ParseDuration(*dieDurationString)
	if err != nil {
		log.Fatalf("Cannot parse duration %v", err)
	}
	ge, err := time.ParseDuration(*generateEachString)
	if err != nil {
		log.Fatalf("Cannot parse duration %v", err)
	}
	brk := time.After(d)

	for {
		select {
		case _ = <- time.After(ge):
			fmt.Fprintln(os.Stdout, time.Now())
		case _ = <- brk:
			fmt.Fprintln(os.Stdout, *endString)
			return
		}
	}
}
