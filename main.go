package main

import (
	"bufio"
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
		scanner := bufio.NewScanner(os.Stdin)

		lines := make([]string, 0)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		selector := selecto.NewSelecto(lines)

		selected, err := selector.Select()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(os.Stdout, selected)
	default:
		return
	}
}
