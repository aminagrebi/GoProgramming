package main

import "fmt"

func main() {
	
	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := 42
	isTrue := true

	stringLengnth, err := fmt.Println(str1,str2,str3)

    if err == nil {
		fmt.Println("String Length:", stringLengnth)


	}

    fmt.Printf("Value of aNumber: %v\n", aNumber)
    fmt.Printf("Value of isTrue: %v\n", isTrue)	
    fmt.Printf("Value of aNumber as float: %.2f\n", float64(aNumber))

}
