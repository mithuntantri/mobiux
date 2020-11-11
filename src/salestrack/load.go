package main

import (
	"fmt"
	"time"
	"strings"
	"strconv"
)

func loadDataset(dataset string) error{
	fmt.Println("===== LOADING DATASET FROM CSV =====")
	records, err := ReadDataset(dataset)
    if err != nil {
        panic(err)
    }

    // Loop through lines & turn into object
    for i, record := range records {

    	// Skip the Header for Mapping to Interface
    	if i != 0{

	        var data SalesDataset
	        data.Date = record[0]

	        data.Year, err = strconv.Atoi(strings.Split(record[0], "-")[0])
	        if err != nil{ return err }
	        
	        data.Month, err = strconv.Atoi(strings.Split(record[0], "-")[1])
	        if err != nil{ return err }

	        data.SKU = record[1]

	        // Convert ASCII representation of Unit Price to Integer
	        data.UnitPrice, err = strconv.Atoi(record[2])
	        if err != nil{ return err }
	        
	        // Convert ASCII representation of Quantity to Integer
	        data.Quantity, err =  strconv.Atoi(record[3])
	        if err != nil{ return err }
	        
	        // Convert ASCII representation of TotalPrice to Integer
	        data.TotalPrice, err = strconv.Atoi(record[4])
	        if err != nil{ return err }

	        layout := "2006-01-02"
			t, err := time.Parse(layout, data.Date)
	        if err != nil{ return err }

			key := data.Date[:4]+"-"+t.Month().String()
			collections[key] = append(collections[key], data)

	        // Append to the All Sales Dataset
	        allSalesDataset = append(allSalesDataset, data)

	        // Step2: Calculate Total Sales of the store
	        total_sales = total_sales + data.TotalPrice

	    }
    }
    return nil
}