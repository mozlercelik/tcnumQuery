package main

import (
	"fmt"
	"gotcquery/tcnumquery"
)

func main() {
	var tcnum string
	fmt.Print("TC kimlik numaranızı girin: ")
	fmt.Scan(&tcnum)

	if tcnumquery.Querytc(tcnum) {
		fmt.Println("TC Numarası geçerlidir.")
	} else {
		fmt.Println("TC Numarası geçersizdir.")
	}

}
