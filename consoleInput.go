package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)


func main() {

	//var s string
	//fmt.Scanln(&s)
	//fmt.Println(s)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)
	
	fmt.Print("Enter a number: ")
	str,_ = reader.ReadString('\n')

	num, err := strconv.ParseFloat(strings.TrimSpace(str), 64)

	if err == nil {
		fmt.Printf("You entered %v", num)
	} else{
		fmt.Printf("Error  %v", err)
	}
	 
}