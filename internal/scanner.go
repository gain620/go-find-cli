package internal

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

func PrintFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println("Cannot open the given file: ", file)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}