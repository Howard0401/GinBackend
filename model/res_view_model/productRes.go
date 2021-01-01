package res

import "github.com/shopspring/decimal"

type Product struct {
	Id                   string          `json:"id"`
	Key                  string          `json:"key"`
	ProductId            string          `json:"productId"`
	ProductName          string          `json:"productName"`
	ProductIntro         string          `json:"productIntro"`
	CategoryId           string          `json:"categoryId"`
	CategoryName         string          `json:"categoryName"`
	ProductCoverImg      string          `json:"productCoverImg"`
	ProductBanner        string          `json:"productBanner"`
	OriginalPrice        decimal.Decimal `json:"originalPrice"`
	SellingPrice         decimal.Decimal `json:"sellingPrice"`
	StockNum             int             `json:"stockNum"`
	Tag                  string          `json:"tag"`
	SellStatus           int             `json:"sellStatus"`
	ProductDetailContent string          `json:"productDetailContent"`
	IsDeleted            bool            `json:"isDeleted"`
}
