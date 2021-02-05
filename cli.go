package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// TODO: add more cli things
func cli() {
	reader := bufio.NewReader(os.Stdin)
	// Allow for boot up text to land before we start reading text in
	time.Sleep(3 * time.Second)
	for {
		fmt.Print("StrongAuras > ")
		text, _ := reader.ReadString('\n')
		fmt.Println(strings.Index(text, "exit"))
		if strings.Index(text, "exit") == 0 {
			fmt.Println("Exiting")
			os.Exit(0)
		}
	}
}
