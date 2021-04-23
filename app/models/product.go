package models

type Product struct {
	ID       uint64
	Label    string
	Quantity uint32
	Price    uint32
}

type ProductResponse struct {
	ID       uint64
	Label    string
	Quantity uint32
	Price    uint32
}

type ProductRequest struct {
	Label    string
	Quantity uint32
	Price    uint32
}