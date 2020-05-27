package main

import (
	"time"
	"fmt"
	"os"
	"flag"
	"log"
	"strconv"
)

// gena is just for generating testing strings
func main() {
	endString := flag.String("end", "END", "End string")
	dieDurationString := flag.String("d", "3s", "Die after")
	generateEachString := flag.String("e", "1s", "Generate string each N")
	seqBool := flag.Bool("seq", false, "Generate int sequence")
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
	seq := 0

	for {
		select {
		case _ = <- time.After(ge):
			text := time.Now().String()
			if *seqBool {
				text = strconv.Itoa(seq)
			}
			fmt.Fprintln(os.Stdout, text)
			seq++
		case _ = <- brk:
			fmt.Fprintln(os.Stdout, *endString)
			return
		}
	}
}
