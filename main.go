package main

import (
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
		s, err := selecto.NewSelecto(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}

		result := s.WaitForSelect()
		if result.Error != nil {
			log.Fatal(err)
		}

		selectedLine := *result.Line
		if selectedLine != "" {
			fmt.Fprintf(os.Stdout, "%s\n", selectedLine)
		}
	default:
		return
	}
}
