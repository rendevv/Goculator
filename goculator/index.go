package main

import (
	"fmt"
	"goculator/calculator"
)

func main() {
	fmt.Println("========== Go Kalkulator ===========")
	kali := "kali"
	kurang := "kurang"
	bagi := "bagi"
	tambah := "tambah"
	sisaBagi := "sisa bagi"
	result, err := calculator.Calculate(2, 4, tambah)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Hasil %s: %d\n", tambah, result)
	}

	result, err = calculator.Calculate(5, 3, sisaBagi)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Hasil %s: %d\n", sisaBagi, result)
	}

	result, err = calculator.Calculate(5, 3, kali)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Hasil %s: %d\n", kali, result)
	}

	result, err = calculator.Calculate(5, 3, kurang)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Hasil %s: %d\n", kurang, result)
	}

	result, err = calculator.Calculate(5, 3, bagi)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Hasil %s: %d\n", bagi, result)
	}
	fmt.Println("=====================")
}
