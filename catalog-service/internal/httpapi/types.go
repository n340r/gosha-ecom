package httpapi

type GetProductsResponse struct {
	Success    bool       `json:"success"`
	Pagination Pagination `json:"pagination"`
	Products   []Product  `json:"products"`
}

type Pagination struct {
	Limit          int `json:"limit"`
	TotalCount     int `json:"totalCount"`
	CurrentPage    int `json:"currentPage"`
	TotalPageCount int `json:"totalPageCount"`
}

type Product struct {
	Type         string   `json:"type"`
	MinPrice     int      `json:"minPrice"`
	MaxPrice     int      `json:"maxPrice"`
	CatalogID    int      `json:"catalogId"`
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	ImageURL     string   `json:"imageUrl"`
	Description  string   `json:"description"`
	Manufacturer string   `json:"manufacturer"`
	Options      []Option `json:"options,omitempty"`
	Offers       []Offer  `json:"offers"`
	UpdatedAt    string   `json:"updatedAt"`
	Active       bool     `json:"active"`
	Quantity     int      `json:"quantity"`
	Markable     bool     `json:"markable"`
}

type Option struct {
	Code   string        `json:"code"`
	Values []OptionValue `json:"values"`
}

type OptionValue struct {
	Value   string `json:"value"`
	Default bool   `json:"default"`
}

type Offer struct {
	Name          string      `json:"name"`
	Price         int         `json:"price"`
	Images        []string    `json:"images"`
	ID            int         `json:"id"`
	PurchasePrice int         `json:"purchasePrice"`
	VatRate       string      `json:"vatRate"`
	Properties    OfferProps  `json:"properties"`
	Quantity      int         `json:"quantity"`
	Active        bool        `json:"active"`
	Prices        []OfferPrice `json:"prices,omitempty"`
}

type OfferPrice struct {
	PriceType string `json:"priceType"`
	Price     int    `json:"price"`
	Ordering  int    `json:"ordering"`
	Currency  string `json:"currency"`
}

type OfferProps struct {
	Size  string `json:"size,omitempty"`
	Color string `json:"color,omitempty"`
}
