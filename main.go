package main

import (
	// "bufio"
	// "io"
	"fmt"
	"log"
	"os"
	"selecto/selecto"
)

func printHelp() {
	fmt.Println(`Selecto use STDIN to select`)
	fmt.Println(`Usage: selecto [--stdin]
For help use --help`)
}

func main() {
	x := os.Args
	if len(x) == 1 {
		printHelp()
		return
	}

	switch x[1] {
	case "--help":
		printHelp()
		return
	case "--stdin":
		resultChan, err := selecto.NewSelecto(os.Stdin)
		if err == nil {
			log.Fatal(err)
		}

		fmt.Fprintln(os.Stdout, "Wait for chan")
		result := <- resultChan
		fmt.Fprintln(os.Stdout, "Chan ok")
		if result.Error != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(os.Stdout, "%s\n", *result.Line)
	default:
		return
	}
}
