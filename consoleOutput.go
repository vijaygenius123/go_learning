package main

import (
	"fmt"
)

func main(){

	str1 := "The quick brown fox"
	str2 := "jumped over the"
	str3 := "the lazy brow dog"

	aNumber := 42
	isTrue := true

	stringLen, err := fmt.Println(str1+str2+str3)

	if err == nil{
		fmt.Println("String Length: ", stringLen)
	}

	fmt.Printf("Value Of aNumber: %v\n", aNumber)
	fmt.Printf("Value Of isTrue: %v\n", isTrue)
	fmt.Printf("Value Of aNumber: %.2f\n", float64(aNumber))

	fmt.Printf("Data types: %T %T %T %T %T\n", str1,str2,str3,aNumber,isTrue)

	myString := fmt.Sprintf("Data types as var: %T %T %T %T %T\n", str1,str2,str3,aNumber,isTrue)
	fmt.Print(myString)

}