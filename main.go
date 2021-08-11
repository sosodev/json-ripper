package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sosodev/json-ripper/ripper"
)

func main() {
	stdinScanner := bufio.NewScanner(os.Stdin)

	fullStdin := ""
	for stdinScanner.Scan() {
		text :=	stdinScanner.Text()
		fullStdin += text
		fmt.Println(text)
	}

	if stdinScanner.Err() != nil {
		panic(stdinScanner.Err())
	}

	ripped, err := ripper.RipJSON(fullStdin)
	if err != nil {
		fmt.Printf("failed to rip JSON: %s", err)
		os.Exit(1)
	}

	fmt.Printf("\n%s\n", ripped)
}
