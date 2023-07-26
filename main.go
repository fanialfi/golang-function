package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	printMessage("Hallo", []string{"fani", "alfi", "fanialfi"})
	fmt.Println()

	// rand.Seed(time.Now().Unix())
	var randomValue int

	randomValue = randomWithRange(1, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(1, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(1, 3)
	fmt.Println("random number:", randomValue)

	// mencetak waktu sekarang dalam format unix => 10 digit
	var sekarang = time.Now().Unix()

	fmt.Println(sekarang)
	fmt.Printf("tipe data %T\n", sekarang)
}

// function return value
func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, ",")
	fmt.Println(message, nameString)
}
