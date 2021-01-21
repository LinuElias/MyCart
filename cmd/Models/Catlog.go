package Models

type Catlog struct {
	CategoryName string        `json:"categoryName"`
	SubCategory  []SubCategory `json:"subCategory"`
}

type SubCategory struct {
	SubCategoryName string     `json:"subCategoryName"`
	Products        []Products `json:"products"`
}

type Products struct {
	ProductName   string  `json:"productName"`
	Price         float64 `json:"price"`
	Specification string  `json:"specification,omitempty"`
}
