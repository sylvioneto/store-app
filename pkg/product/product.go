package product

// Product is the structure that defines store item
type Product struct {
	ID int
	Name string
	Price float64
	Quantity int
}

// GetTestData returns mock data
func GetTestData() []Product {
	products := []Product{
		{1, "Rice", 11.99, 5},
		{2, "Water", 0.99, 20},
		{3, "Meat", 29.11, 1},
	}
	return products
}
