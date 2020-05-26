package main

import (
	"bufio"
	"os"
	"fmt"
	"sync"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var wg sync.WaitGroup
	input := make(chan string, 20)
	wg.Add(1)
	go func() {
		for t := range input {
			fmt.Fprintf(os.Stdout, "Cow say: %s\n", t)
		}
		wg.Done()
	}()

	for scanner.Scan() {
		txt := scanner.Text()
		input <- txt
		// fmt.Fprintf(os.Stdout, "Scanned: %s\n", txt)
	}

	close(input)
	wg.Wait()
}
