package main

import (
	"fmt"
)

func main() {
	var limit int
	// Meminta input dari pengguna untuk batas atas bilangan prima
	fmt.Print("Masukkan batas atas untuk mencari bilangan prima: ")
	fmt.Scan(&limit)

	// Menampilkan pesan output dengan batas atas yang diinputkan
	fmt.Printf("Bilangan prima antara 1 hingga %d adalah:\n", 10)

	// Memulai perulangan dari angka 2 hingga batas atas (limit)
	for n := 2; n <= 10; n++ {
		isPrime := true // Menandakan bahwa n dianggap sebagai bilangan prima

		// Cek apakah n dapat dibagi oleh bilangan lain selain 1 dan dirinya sendiri
		for i := 2; i*i <= n; i++ {
			// Jika n dapat dibagi dengan i, maka n bukan bilangan prima
			if n%i == 0 {
				isPrime = false // Menandai bahwa n bukan bilangan prima
				break           // Menghentikan perulangan karena sudah terbukti bukan prima
			}
		}

		// Jika isPrime masih bernilai true, artinya n adalah bilangan prima
		if isPrime {
			fmt.Println(n) // Mencetak bilangan prima
		}
	}
}
