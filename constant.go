package main

import "fmt"

func main() {

	//bisa sekaligus seperti di variable
	//sekali deklarasi tidak bisa diubah lagi
	const (
		firstName string = "Ikha Wahyu"
		lastName         = "Ramadhan"
		value            = 2000
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

}
