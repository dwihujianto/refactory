package main

import(
	"bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Nama Warung ? ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("Tanggal ? ")
	scanner.Scan()
	date := scanner.Text()

	fmt.Print("Nama Kasir ? ")
	scanner.Scan()
	casher := scanner.Text()

	items := make(map[int]map[string]string)


    var i = 0
    for {
        fmt.Print("Nama item : ")
        scanner.Scan()
        item := scanner.Text()

        fmt.Print("Harga item : ")
        scanner.Scan()
        price := scanner.Text()

        if len(item) != 0 && len(price) != 0 {
            items[i] = map[string]string{
		    	"item" : item,
		    	"price": price,
		    }
        } else {
            break
        }

        i = i + 1
    }

    nameLeft := leftPad("Nama Warung", ".", 16)
    nameRight := rightPad(name, ".", 16)
    fmt.Println(nameLeft + nameRight)

    dateLeft := leftPad("Tanggal", ".", 16)
    dateRight := rightPad(date, ".", 16)
    fmt.Println(dateLeft + dateRight)

    casherLeft := leftPad("Kasir", ".", 16)
    casherRight := rightPad(casher, ".", 16)
    fmt.Println(casherLeft + casherRight)

    fmt.Println(strings.Repeat("=", 32))

    var total int

    for _ , value := range items {
    	left := leftPad(value["item"], ".", 16)
    	right := rightPad(value["price"], ".", 16)
		fmt.Println(left + right)
		i, _ := strconv.Atoi(value["price"])
		total = total + i
	}

	fmt.Println(leftPad("Total", ".", 16) + rightPad(toCurrency(total), ".", 16))
}

func toCurrency(i int) string {
	c := strconv.Itoa(i)
	return c
}

func leftPad(s string, padStr string, overallLen int) string {
	var padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}

func rightPad(s string, padStr string, overallLen int) string {
	var padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}