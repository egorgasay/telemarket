package entity

// Item represents an item in the shop.
type Item struct {
	Name        string
	Description string
	Image       string
	Price       float32
	Quantity    int
	PathToPhoto string `json:"path_to_photo"`
}

// Information represents the information about the shop.
type Information struct {
	Avg float64
}
