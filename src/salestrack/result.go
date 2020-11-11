package main

import (
	"fmt"
	"strings"
)

func getResult() string{
	result := "---------------------------------------------------------------------------------------------------------------------------------\n"
	result += "| YEAR\t| MONTH\t\t| SALES\t\t| Most Popular Item\t\t\t\t| Most Revenue Generating Item\t\t|\n"
	result += "---------------------------------------------------------------------------------------------------------------------------------\n"
	for key, values := range collections{
		year := strings.Split(key, "-")[0]
		month := strings.Split(key, "-")[1]
		var total_sales int
		
		var groupedSKUs = make(map[string]int)
		var groupedSKUsMinOrders = make(map[string]int)
		var groupedSKUsMaxOrders = make(map[string]int)
		
		most_popular_item := values[0]
		var most_revenue_generating_item string
		var most_revenue_generating_item_value int

		for _, v := range values{
			groupedSKUs[v.SKU] = groupedSKUs[v.SKU] + v.TotalPrice

			if v.Quantity < groupedSKUsMinOrders[v.SKU] || groupedSKUsMinOrders[v.SKU] == 0{
				groupedSKUsMinOrders[v.SKU] = v.Quantity
			}

			if v.Quantity > groupedSKUsMaxOrders[v.SKU] || groupedSKUsMaxOrders[v.SKU] == 0{
				groupedSKUsMaxOrders[v.SKU] = v.Quantity
			}

			if groupedSKUs[v.SKU] > most_revenue_generating_item_value{
				most_revenue_generating_item = v.SKU
				most_revenue_generating_item_value = groupedSKUs[v.SKU]
			}

			if v.Quantity > most_popular_item.Quantity{
				most_popular_item = v
			}
			total_sales = total_sales + v.TotalPrice
		}

		min_order_quantity := groupedSKUsMinOrders[most_popular_item.SKU]
		max_order_quantity := groupedSKUsMaxOrders[most_popular_item.SKU]
		avg_order_quantity := (min_order_quantity+max_order_quantity)/2

		result += fmt.Sprintf("| %s\t| %20 s\t| %d\t| %s (Min:%d, Max:%d, Avg:%d)\t| %s (Revenue:%d)\t|\n", year, month, total_sales, most_popular_item.SKU, min_order_quantity, max_order_quantity, avg_order_quantity, most_revenue_generating_item, most_revenue_generating_item_value)
	}
	result += "---------------------------------------------------------------------------------------------------------------------------------"
	return result
}
