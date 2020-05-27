package main

import (
	"bufio"
	"os"
	"fmt"
	"sync"
)

// reedo is just a helper tool to check concept
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
		input <- scanner.Text()
	}

	close(input)
	wg.Wait()
}
