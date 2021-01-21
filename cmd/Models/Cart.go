package Models

type Cart struct {
	Product         Products `json:"product"`
	CategoryName    string   `json:"categoryName"`
	SubCategoryName string   `json:"subCategoryName"`
}

type Bill struct {
	Product            []Products `json:"product"`
	TotalAmount        float64    `json:"totalAmount"`
	IsDiscountApplied  bool       `json:"isDiscountApplied"`
	AmountWithDiscount float64    `json:"amountWithDiscount"`
}
