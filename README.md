# golang function

function merupakan sekumpulan blok kode yang dibungkus dengan nama tertentu, penerapan fungsi yang tepat dapat menjadikan kode lebih modular dan juga **dry**, tak perlu menuliskan banyak kode yang kedunaannya sama, cukup tulis sekali lalu gunakan berkali kali.

Cara pembuatan `function` cuku mudah, gunakan keyword `func` lalu diikuti dengan nama function-nya, tanda kurung & kurung kurawal.

```go
func main(){
  // do nothing here
}
```

seperti contoh diatas, jika ingin menggunakan parameter di function-nya cukup tulis nama parameter beserta tipe data parameter-nya didalam blok tanda kurung (`()`), jika ingin mengirimkan banyak parameter maka pisahkan parameter satu dengan yang lain dengan tanda koma (`,`)

contoh penggunaan function dengan parameter :

```go
package main

import "fmt"

func main(){
  const data = [3]string{"fani","alfi","fanialfi"}

  cetakHallo("hallo",data)
}

func cetakHallo(message string,names []string){
  const nm = strings.Join(names, ",")
  fmt.Println(message, nm)
}
```

didalam function `cetakHallo` parameter _names_ yang berupa slice digabungkan menjadi string dengan pembatas menggunakan tanda koma (`,`), penggabungannya dengan menggunakan function bawaan `strings.Join()` yang berada dalam package `strings`.

sebuah function bisa mengembalikan suatu nilai, function yang mengembalikan nilai harus ditentukan tipe data nilai kembalian pada saat pendeklarasiannya.

```go
package main

import (
  "fmt"
  "math/rand"
)
func main(){
  var randomValue = cetakRandom(1,10)
  fmt.Println("angka random :",randomValue)
  var randomValue = cetakRandom(1,10)
  fmt.Println("angka random :",randomValue)
  var randomValue = cetakRandom(1,10)
  fmt.Println("angka random :",randomValue)
}

func cetakRandom(min,max int)int{
  var number = rand.Int()% (max - min + 1) + min

  return number
}
```

function `cetakRandom` tersebut berfungsi untuk mengenerate angka random sesuai dengan range yang ditentukan yang kemudian angka tersebut dijadikan sebagai return value dari function tersebut.

Untuk menentukan tipe data return value dari function yaitu dengan cara menulisakan tipe data yang diinginkan setelah kurung parameter `func cetakRandom(min, max int) int {}` sedangkan untuk mengembalikan nilainya itu sendiri dengan cara menggunakan keyword `return` diikuti dengan datanya.

Jika sebuah function dijalankan dan menemukan keyword `return` dimana dibawah keyword tersebut masih ada block kode, maka block kode tersebut tidak akan dijalankan, atau sama saja ketika sudah bertemu dengan keyword `return` maka program akan berhenti.

Untuk function dengan parameter dimana tipe data parameter-nya sama, maka bisa cukup tuliskan tipe data-nya sekali saja di akhir.

Selain digunakan untuk mengembalikan value keyword `return` juga bisa digunakan untuk menghentikan proses dalam sebuah blok kode dimana ia dipakai.

Jika ada kebutuhan dimana data yang dikembalikan harus banyak, biasanya digunakan tipe data `map` atau `slice`,dan juga GO menyediakan kapabilitas bagi programmer untuk membuat function memiliki banyak return.

contoh function mereturn banyak nilai :

```go
package main

import (
  "fmt"
  "math"
)

func main(){
  var dm float64 = 20
  var luas, keliling = calculate(dm)

  fmt.Printf("luas lingkaran : %f\n",luas)
  fmt.Printf("keliling lingkaran : %f\n",keliling)
}

func calculate(diameter float64)(float64, float64){
  // hitung luas
  var luas = math.Pi * math.Pow(diameter / 2.0, 2.0)

  // hitung keliling
  var keliling = math.Pi * diameter

  // kembalikan 2 nilai

  return luas, keliling
}
```

cara pendefinisian multiple return di atas yaitu dengan menuliskan tipe data apa saja yang akan di return dan disimpan didalam tanda kurung (`()`) setelah pendeklarasian paameter pada function, lalu untuk setiap tipe data return value dipisahkan dengan tanda koma `(float64, float64){}`.

didalam keyword `return` harus dituliskan semua return value nya `return luas, keliling`

lalu untuk penggunaan function tersebut jika ingin disimpan dalam sebuah variabel maka harus disimpan didalam 2 variabel `var luas, keliling = calculate(10)`.

selain itu value yang dijadikan sebagai _return_ bisa didefinisikan diawal 

```go
func calculate(diameter float64)(luas float64, keliling float64){
  var luas = math.Pi * math.Pow(diameter / 2.0, 2.0)
  var keliling = math.Pi * diameter
  return 
}
```

beberapa hal baru diatas yang perlu dibahas

- penggunaan function `math.Pow(a,b)`
  
  function `math.Pow()` digunakan untuk memangkatkan nilai, `math.Pow(2,3)` bararti 2 dipangkat 3 hasilnya 8 (`a ** b`) dimana function ini berada didalam package `math`.

- penggunaan konstanta `math.Pi`

  `math.Pi` adalah konstanta yang merepresentasikan **pi** atau **22/7**