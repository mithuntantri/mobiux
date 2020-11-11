package main

import (
    "fmt"
    "log"
)

var (
	allSalesDataset []SalesDataset
    collections = make(map[string][]SalesDataset)
	total_sales int
)


func main() {
	dataset := "sales-data.csv"
	
	// Step1: Load Sales Dataset from CSV to Interface
	err := loadDataset(dataset)
	if err != nil{
		log.Fatal(err)
		panic(err)
	}
	
	fmt.Println(getResult())
	fmt.Println("TOTAL SALES ::::::::::::::: ", total_sales)	
}
