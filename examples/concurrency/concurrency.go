package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
)

// For Windows
// func getRandomIterCount() int {
// 	randomNumber := rand.Intn(100)

// 	iterCount := 300 + randomNumber

// 	return iterCount
// }

// For Unix
func getRandomIterCount() int {
	randomFile, _ := os.Open("/dev/random")
	randomReader := bufio.NewReader(randomFile)
	randomNumber, _ := rand.Int(randomReader, big.NewInt(100))
	iterCount := int(300 + randomNumber.Int64())

	return iterCount
}

func busyWork(iterCount int) {
	for i := 0; i < iterCount; i++ {
		for j := 0; j < iterCount; j++ {
			mul := []byte(strconv.Itoa(i * j))
			ioutil.WriteFile("/dev/null", mul, 0644)
		}
	}
}

func say(s string, index int) {
	iterCount := getRandomIterCount()

	// CPU & IO intensive work for random iterations
	busyWork(iterCount)

	fmt.Printf("Hello, %s! from %d say fn; Iter Count: %d\n", s, index+1, iterCount)
}

func main() {
	for i := 0; i < 4; i++ {
		say("world", i)
	}

	// fmt.Println("Started all say goroutines!")
	// time.Sleep(10 * time.Second)
}
