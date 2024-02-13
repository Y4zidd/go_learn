package main

import "fmt"

func main() {
  var firstname string = "Yazid Istiqlal"

  var compeny string = "Tokopedia"

  var jobs string = "devops"

  jam := "8 jam" // go deklarasi, bisa tanpa var dan memsukan tipe data seperti diatas, cukup tambahkan ":="

  semoga, tahun, _ := "semoga", "tahun", "saya" //multi variable atau bisa lebih dari 1 dan menambhakan jika ada variable yang tidak digunakan maka berikan "_"

  newString := new(string)

  fmt.Printf("Halo nama saya %s saya bekerja di %s dan masuk di defisi %s dan bekerja selama %s \n dan %s di %s depan bisa menjadi pribadi yang lebib baik lagi \n", 
  firstname, compeny, jobs, jam, semoga, tahun)

  fmt.Printf(*newString)
}
