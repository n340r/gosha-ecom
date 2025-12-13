package httpapi

func MockGetProductsResponse() GetProductsResponse {
	p := Product{
		Type:        "product",
		MinPrice:    5000,
		MaxPrice:    5000,
		CatalogID:   2,
		ID:          74,
		Name:        "порно студия[black]",
		ImageURL:    "https://s3-s1.retailcrm.tech/ru-central1/retailcrm/goshamartynovich-3d1331c972115a9776c29ace80fb18eb/product/68cc60f1c9f5e-porno-black_1-copy.jpg",
		Description: "оверсайз фит/\nметод нанесения-шелкография/\nсрок изготовления 2-5 дней /\nбланк- всяческий :от гилдана до  майки с wb (heavycotton20000ultra)",
		Manufacturer:"SERVISEX",
		Options: []Option{
			{
				Code: "size",
				Values: []OptionValue{
					{Value: "L", Default: true},
					{Value: "M", Default: false},
					{Value: "XL", Default: false},
					{Value: "S", Default: false},
				},
			},
		},
		Offers: []Offer{
			{
				Name: "порно студия[black], S",
				Price: 5000,
				Images: []string{
					"https://s3-s1.retailcrm.tech/ru-central1/retailcrm/goshamartynovich-3d1331c972115a9776c29ace80fb18eb/product/68cc60f1c9f5e-porno-black_1-copy.jpg",
				},
				ID: 199,
				PurchasePrice: 1000,
				VatRate: "none",
				Properties: OfferProps{Size: "S"},
				Quantity: 3,
				Active: true,
			},
			{
				Name: "порно студия[black], M",
				Price: 5000,
				Images: []string{
					"https://s3-s1.retailcrm.tech/ru-central1/retailcrm/goshamartynovich-3d1331c972115a9776c29ace80fb18eb/product/68cc60f1c9f5e-porno-black_1-copy.jpg",
				},
				ID: 200,
				PurchasePrice: 1000,
				VatRate: "none",
				Properties: OfferProps{Size: "M"},
				Quantity: 4,
				Active: true,
			},
		},
		UpdatedAt: "2025-09-19 01:36:25",
		Active: true,
		Quantity: 15,
		Markable: false,
	}

	return GetProductsResponse{
		Success: true,
		Pagination: Pagination{
			Limit:          20,
			TotalCount:     1,
			CurrentPage:    1,
			TotalPageCount: 1,
		},
		Products: []Product{p},
	}
}
