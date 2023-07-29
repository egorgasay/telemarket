package entity

// Item represents an item in the shop.
type Item struct {
	ID          string
	Name        string
	Description string
	Image       string
	Price       string
	Quantity    int32
	PathToPhoto string `json:"path_to_photo"`
}

func (i Item) GetName() string {
	return i.Name
}

func (i Item) GetId() string {
	return i.ID
}

func (i Item) GetDescription() string {
	return i.Description
}

func (i Item) GetPrice() string {
	return i.Price
}

func (i Item) GetImage() string {
	return i.Image
}

func (i Item) GetQuantity() int32 {
	return i.Quantity
}

func (i Item) SetID(id string) {
	i.ID = id
}

// Information represents the information about the shop.
type Information struct {
	Avg float64
}

type IItem interface {
	GetName() string
	GetId() string
	GetDescription() string
	GetPrice() string
	GetImage() string
	GetQuantity() int32
}
