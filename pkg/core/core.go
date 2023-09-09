package core

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func UseFile(name string) {
	bytes, err := os.ReadFile(name)

	if err != nil {
		log.Fatalf("encountered an error when reading file: %s\n", err)
	}

	execute(string(bytes))
}

func UseRepl() {
	reader := bufio.NewReader(os.Stdin)
	ch := make(chan os.Signal, 1)

	go func() {
		fmt.Print("⚡ Flux REPL | Press ^C to Exit\n\n")

		for {
			fmt.Print("⚡ >> ")
			input, err := reader.ReadString('\n')

			if err != nil {
				log.Fatalf("encountered an error when reading input: %s\n", err)
				break
			}

			input = strings.TrimSpace(input)

			if input == "" {
				continue
			}

			execute(input)
		}
	}()

	signal.Notify(
		ch, os.Interrupt, syscall.SIGTERM,
	)

	<-ch

	fmt.Println()
}

func execute(source string) {
	fmt.Println(source)
}
