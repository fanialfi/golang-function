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

import (
  "fmt"
  "strings"
)

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

Selain digunakan untuk mengembalikan value, keyword `return` juga bisa digunakan untuk menghentikan proses dalam sebuah blok kode dimana ia dipakai.

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


## Variadic function

merupakan konsep atau cara pembuatan function dengan parameter sejenis yang tak terbatas / banyak, sifatnya mirip seperti _slice_, nilai yang disisipkan bertipe data sama, dan untuk mengaksesnya sama, dengan menggunakan index.

untuk mendeklarasikan parameter Variadic seperti mendeklarasikan parameter pada umumnya, bedanya parameter variadic saat mendeklarasikan parameter tepat sebelum jenis tipe datanya ditambahkan tanda titik 3 (`...`) yang nantinya semua parameter yang dikirim akan ditampung oleh variabel tersebut.

`func contoh(names ...string){}`

contoh penerapan :

```go
package main

import "fmt"

func main(){
  var rataRata = hitungRataRata(1,2,3,4,5,6,7,8,9,10)
  var msg = fmt.Sprintf("rata ratanya adalah : %.2f",rataRata)
  fmt.Println(msg)
}

func hitungRataRata(numbers ...int)float64{
  var total int = 0

  for _, number := range numbers{
    total += number
  }

  var rataRata = float64(total) / float64(len(numbers))

  return rataRata
}
```

pada function `hitungRataRata`, parameter `numbers` disisipkan dengan titik 3 (`...`) menandakan bahwa parameter tersebut adalah sebuah parameter variadic bertipe `int`.

pemanggilan function tersebut sama seperti biasa, hanya saja jumplah parameter-nya bisa banyak.

beberapa penjelasan dari kode diatas :

- penggunaan function `fmt.Sprintf()`

  function `fmt.Sprintf()` ini pada dasarnya sama seperti `fmt.Printf()` hanya saja tidak menampilkan nilai, melainkan mengebalikan nilainya dalam bentuk string.
  Selain `fmt.Sprintf()` ada juga `fmt.Sprint()` dan `fmt.Sprintln()`.
- penggunaan function `float64()`

  sebelumnya kita sudah tau bahwa `float64` merupakan tipe data, tipe yang dilulis sebagai function (ada kurungnya) berguna untuk **casting**, casting sendiri adalah teknik untuk mengkonversi tipe data ke tipe data yang lain.

#
 operasi bilangan (perkalian, pembagian, penjumplahan, pengurangan, dll) di GO hanya bisa dilakukan jika tipe data-nya sama. maka dari itu diperlukanlah adanya casting tipe data.

slice juga bisa digunakan sebagai parameter variadic, caranya dengan menambahkan titik 3 (`...`) setelah menuliskan variabel bertipe slice saat penggunaan function.

```go
package main

import (
  "fmt"
  "strings"
)

func main(){
  var hobies = []string{"membaca","menulis","memantun"}
  var hobiku = cetakHobies("fani",hobies...)
  fmt.Println(hobiku)
}

func cetakHobies(nama string, hobi ...string)string{
  var hobies = strings.Join(hobi,",")
  var msg = fmt.Sprintf("hallo nama saya adalah %s, hobi saya adalah \n\t%s\n",nama,hobies)

  return msg
}
```

dari contoh di atas parameter variadic dikombinasikan dengan parameter biasa, jika mengkombinasikan parameter variadic dengan parameter biasa, maka syaratnya adalah parameter variadic tersebut harus diletakkan di paling akhir.