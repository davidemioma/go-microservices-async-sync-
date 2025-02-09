package common

type ItemType struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Quantity int32  `json:"quantity"`
	PriceID  string `json:"priceId"`
}

type OrderType struct {
	ID         string     `json:"id"`
	CustomerID string     `json:"customerID"`
	Status     string     `json:"status"`
	Items      []ItemType `json:"items"`
}