package models

type Product struct {
	ID       uint64
	Label    string
	Quantity int
	Price    int
}

type ProductResponse struct {
	ID       uint64
	Label    string
	Quantity int
	Price    int
}

type ProductRequest struct {
	Label    string
	Quantity int
	Price    int
}