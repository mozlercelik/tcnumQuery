# tcnumQuery
Bu fonksiyon ile TC kimlik adresinin doğruluğunu kontrol edebilirsiniz.

## Usage

```
go run main.go
```
### Or
```go
package main

import (
	"fmt"
	"gotcquery/tcnumquery"
)

func main() {
	var tcno string = "1234567890"

	if tcnumquery.Querytc(tcno) {
		fmt.Println("TC Numarası geçerlidir.")
	} else {
		fmt.Println("TC Numarası geçersizdir.")
	}

}
```
