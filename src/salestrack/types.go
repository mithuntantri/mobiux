package main

type SalesDataset struct {
    Date string `json:"date"`
    Year int `json:"year"`
    Month int `json:"month"`
    SKU string `json:"sku"`
    UnitPrice int `json:"unit_price"`
    Quantity int `json:"quantity"`
    TotalPrice int `json:"total_price"`
}

type GroupedSKU struct{
	SKU string `json:"sku`
	TotalRevenue int `json:"total_revenue"`
	MinOrder int `json:"min_order"`
	MaxOrder int `json:"max_order"`
}