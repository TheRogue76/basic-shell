package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func execInput(in string) error {
	input := strings.TrimSuffix(in, "\n")

	cmd := exec.Command(input)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}

		if err = execInput(input); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}

	}
}
