package items

// Item is the structure that defines store item
type Item struct {
	ID int
	Name string
	Price float64
	Quantity int
}

func GetTestData() []Item {
	items := []Item{
		{1, "Rice", 11.99, 5},
		{2, "Water", 0.99, 20},
		{3, "Meat", 29.11, 1},
	}
	return items
}
