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

	// Closure function
	var getMinMax = func(n []int) (int, int) {
		var min, max int = 0, 0

		for index, element := range n {
			switch {
			case index == 0:
				max, min = element, element
			case element > max:
				max = element
			case element < min:
				min = element
			}
		}
		return min, max
	}

	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numbers)

	fmt.Printf("data %v\nmin : %v\nmax : %v\n", numbers, min, max)

	// Immediately-Invoked Function expression (IIFE)
	fmt.Println()
	numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	var newNumbers = func(min int) []int {
		var r []int

		for _, e := range numbers {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(4)

	fmt.Println("original numbers", numbers)
	fmt.Println("filtered numbers", newNumbers)

	// function return function
	fmt.Println()
	var find = 2
	var ab, bc = findMax(numbers, find)
	var bcNum = bc()

	fmt.Println("numbers \t:", numbers)
	fmt.Printf("find \t\t: %d\n\n", find)

	fmt.Println("found :", ab)
	fmt.Println("value :", bcNum)

	// pemanggilan function sebagai parameter
	fmt.Printf("\n############################################\n\n")
	var names = []string{"fani", "alfi", "fanialfi", "shadow", "wich", "petter", "john"}
	var nmIncFani = filter(names, func(str string) bool {
		return strings.Contains(str, "fani")
	})
	var nmLen4 = filter(names, func(str string) bool {
		return len(str) == 4
	})

	fmt.Println("original names :", names)
	fmt.Printf("\nname's include fani\ntotal\t: %d\ndata\t: %v\n\n", len(nmIncFani), nmIncFani)

	fmt.Println("----------------------------------------------------")

	fmt.Println("original names :", names)
	fmt.Printf("\nname's with length 4\ntotal\t: %d\ndata\t: %v\n", len(nmLen4), nmLen4)
}

// alias skema closure
type FilterCallback func(str string) bool // membuat skema function

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

// function return function
func findMax(numbers []int, max int) (int, func() []int) {
	var res []int

	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}

	return len(res), func() []int {
		return res
	}
}

// function sebagai parameter
func filter(data []string, callback FilterCallback) []string {
	var res []string

	for _, each := range data {
		if filtered := callback(each); filtered {
			res = append(res, each)
		}
	}

	return res
}
