package main

import (
	"fmt"
	"os"
)

type Books struct {
	Name string
}

func main() {
	args := os.Args
	mybooks := make([]string, 4)
	mybooks[0] = "Harry Potter"
	mybooks[1] = "LOTR"
	mybooks[2] = "Denemeler"
	mybooks[3] = "Suç ve Ceza"
	if len(mybooks) == 4 {
		mybooks := make([]string, 4)
		mybooks[0] = "Harry Potter"
		mybooks[1] = "LOTR"
		mybooks[2] = "Denemeler"
		mybooks[3] = "Suç ve Ceza"
		fmt.Printf("%s uygulamasında kullanabileceğiniz komutlar : \n search => arama işlemi için \n list => listeleme işlemi için\n", mybooks)
		fmt.Printf("Kullanıcı tarafından girilen ilk değer : %s\n", args[1])
		if args[1] == "list" {
			fmt.Println(mybooks)
		} else {
			fmt.Println("yok")
		}
		return
	}

}
