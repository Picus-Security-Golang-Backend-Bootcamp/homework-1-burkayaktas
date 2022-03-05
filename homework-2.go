package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var ErrNoSufficientBalance = errors.New("Stokta yeterli kod bulunmamaktadır.")

type Booklist struct {
	Names      string
	Id         int
	Stock_code int
	ISBN_no    int
	Page_no    int
	Price      int
	Stock_no   int
	Author     string
}

// var mybooklistarray []Booklist

func main() {
	args := os.Args
	a_list := Booklist{
		Names:      "Harry Potter",
		Id:         1,
		Stock_code: rand.Int(),
		ISBN_no:    rand.Int(),
		Page_no:    rand.Int(),
		Price:      80,
		Stock_no:   50,
		Author:     "JK Rowling",
	}
	mybooklist := Booklist{
		Names:      "Lord Of The Rings",
		Id:         2,
		Stock_code: rand.Int(),
		ISBN_no:    rand.Int(),
		Page_no:    rand.Int(),
		Price:      100,
		Stock_no:   50,
		Author:     "JR Tolkien",
	}
	b_list := Booklist{
		Names:      "Game Of Thrones",
		Id:         3,
		Stock_code: rand.Int(),
		ISBN_no:    rand.Int(),
		Page_no:    rand.Int(),
		Price:      80,
		Stock_no:   34,
		Author:     "RR Martin",
	}
	mybooklistarray := []Booklist{a_list, mybooklist, b_list}
	if len(args) > 2 && args[1] == "search" {
		searchItem := strings.Join(args[2:], " ")
		searchItemInt, _ := strconv.Atoi(searchItem)
		fmt.Printf("Aranan değer:  \"%s\n", searchItem)
		for _, v := range mybooklistarray {
			if v.Stock_no == searchItemInt || v.Names == searchItem || v.Author == searchItem || v.ISBN_no == searchItemInt || v.Page_no == searchItemInt || v.Price == searchItemInt {
				fmt.Println("Buldu:", v)
			}
		}
		if len(args) > 2 && args[2] == "buy" {
			buyItem := strings.Join(args[2:], " ")
			// Sıkıntının burda olduğunu düşünüyorum.
			fmt.Printf("Aranan değer:  \"%s\n", buyItem)
			b := Booklist{"Percy Jackson", 4, rand.Int(), rand.Int(), rand.Int(), 80, 35, "Jackson Samuel"}

			err := b.Buy(1)

			if err != nil {
				fmt.Println(err.Error())
				return
			}
			// HOCAM BURDA YAPTIĞIM HATAYI BİR TÜRLÜ ANLAYAMADIM. YAPMAYA ÇALIŞTIĞIM ŞEY POINTER OLUŞTURUP BUY KOMUTU SONRASI ÇEKMEYE ÇALIŞMAKTI SEARCHE UYARLAMAYA ÇALIŞTIM FAKAT BAŞARAMADIM. SİZE BUNU ŞU AN YOLLUYORUM VE GÜN İÇERİSİNDE ÇÖZMEK İÇİN UĞRAŞACAĞIM. BANA BİR FEEDBACK VERİP YOL GÖSTERİRSENİZ DE SEVİNİRİM.
			PrintBooklist(b)

		}

	}
}
func PrintBooklist(a Booklist) {

}
func (v *Booklist) Buy(Stock_no int) error {
	if v.Stock_no < 1 {
		return ErrNoSufficientBalance
	}

	v.Stock_code -= 1
	return nil
}
