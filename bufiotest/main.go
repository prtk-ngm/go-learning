package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"
)

func main() {

	stdout, _ := os.OpenFile("/Users/pratenig/Documents/go-learning/assignment2/dummy.txt", os.O_RDONLY|syscall.O_NONBLOCK, 0600)
	syscall.SetNonblock(int(stdout.Fd()), false)

	bufSize := 1000000 * 128

	scanner := bufio.NewScanner(bufio.NewReaderSize(stdout, bufSize))

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
