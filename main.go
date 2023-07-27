package main

import (
	"fmt"
	"math"
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

	fmt.Println()

	// function multiple return
	var hasil, hasilLagi = calculate(10.1)
	fmt.Println(hasil, hasilLagi)
	fmt.Println(math.Pow(2, 3))

	// Variadic function
	fmt.Println()
	var rataRata = hitungRataRata(1, 2, 3, 3, 3, 3, 3, 3, 3, 4, 5, 6, 7, 7, 8, 2, 3, 9, 12, 12, 33, 4, 5, 66, 6, 6)
	fmt.Println(rataRata)

	// variadic function slice parameter
	fmt.Println()
	var hobies = []string{"membaca", "nonton film", "coding", "belajar"}
	var nameAndHobie = displayHobies("fani", hobies...)
	fmt.Println(nameAndHobie)
}

// function multiple return
func calculate(d float64) (float64, float64) {
	// hitungg luas
	var area = math.Pi * math.Pow(d/2.0, 1.0)

	// hitung keliling
	var circumference = math.Pi * d

	// kembalikan 2 nilai
	return area, circumference
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

// Variadic function
func hitungRataRata(numbers ...int) string {
	var total int = 0

	for _, number := range numbers {
		total = total + number
	}

	var rataRata float64 = float64(total) / float64(len(numbers))
	var msg = fmt.Sprintf("rata rata nilai adalah %.2f", rataRata)

	return msg
}

// variadic function slice parameter
func displayHobies(name string, hobies ...string) string {
	var hobiku = strings.Join(hobies, ",")
	var msg = fmt.Sprintf("hallo nama saya %s hobi saya adalah : \n\t%s\n", name, hobiku)

	return msg
}
